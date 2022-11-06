package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/jomei/notionapi"
	"io"
	"log"
	"myAwesomeWebsite/lib"
	"myAwesomeWebsite/lib/Cache"
	"net/http"
)

var token string
var db string

var cacheType string
var ttl int

var cacheClient *Cache.Cache

func init() {
	flag.StringVar(&token, "token", "Missing Token", "secret from notion API")
	flag.StringVar(&db, "db", "Missing DB", "DB for websites")
	flag.StringVar(&cacheType, "cache", "none", "Cache type to use (redis,none)")
	flag.IntVar(&ttl, "ttl", 3600, "cache time")
	flag.Parse()

	cacheClient = Cache.NewCache(cacheType, ttl)

}

func client() *lib.NotionClient {
	return lib.NewNotionClient(token, db)
}

func main() {

	r := gin.Default()

	website := r.Group("/api/:domain")
	website.GET("", func(c *gin.Context) {
		page := loadPage(c)
		c.JSON(http.StatusOK, gin.H{
			"domain":     c.Param("domain"),
			"headers":    c.Request.Header,
			"page":       page,
			"breadCrumb": loadBreadcrumb(c),
		})
	})
	website.GET("/subpage/:pageId", func(c *gin.Context) {
		page := loadSubpage(c)
		c.JSON(http.StatusOK, gin.H{
			"domain":     c.Param("domain"),
			"subpage":    c.Param("pageId"),
			"headers":    c.Request.Header,
			"page":       page,
			"breadCrumb": loadBreadcrumb(c),
		})
	})

	website.GET("/cover", func(c *gin.Context) {
		content := loadCover(c)
		//c.Writer.Header().Set("Content-Type", "image/jpeg")
		c.Writer.Write(content)
	})
	website.POST("/db/:childDatabase", func(c *gin.Context) {
		var body map[string]interface{}
		c.ShouldBindJSON(&body)
		db := c.Param("childDatabase")
		page, err := client().StoreInDatabase(db, body)
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
		c.JSON(201, page)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func loadPage(c *gin.Context) map[string]interface{} {
	domain := c.Param("domain")
	page := cacheClient.LoadAndCache(domain, func(domain string) interface{} {
		client := client()
		page := client.SearchForDomain(domain)

		return Page{
			Page:   page,
			Blocks: client.ReadBlock(notionapi.BlockID(page.ID)),
		}
	})
	content, _ := json.Marshal(page)
	var finalPage map[string]interface{}
	json.Unmarshal(content, &finalPage)
	return finalPage
}
func loadSubpage(c *gin.Context) map[string]interface{} {
	domain := c.Param("domain")
	pageId := c.Param("pageId")
	page := cacheClient.LoadAndCache(domain+"/"+pageId, func(fullName string) interface{} {
		client := client()
		page := client.GetPage(pageId)

		return Page{
			Page:   page,
			Blocks: client.ReadBlock(notionapi.BlockID(page.ID)),
		}
	})
	content, _ := json.Marshal(page)
	var finalPage map[string]interface{}
	json.Unmarshal(content, &finalPage)
	return finalPage
}

func loadBreadcrumb(c *gin.Context) interface{} {
	domain := c.Param("domain")
	pageId := c.Param("pageId")
	if len(pageId) < 1 {
		return nil
	}

	breadCrumb := cacheClient.LoadAndCache("breadcrumb:"+domain+"/"+pageId, func(fullName string) interface{} {
		currentId := pageId
		pageDomain := client().SearchForDomain(domain)
		breadcrumb := []*notionapi.Page{}
		for {
			log.Println(currentId)
			currentPage := client().GetPage(currentId)
			if currentPage.ID == pageDomain.ID {
				break
			}
			breadcrumb = append([]*notionapi.Page{currentPage}, breadcrumb...)
			switch currentPage.Parent.Type {
			case notionapi.ParentTypeDatabaseID:
				currentId = string(currentPage.Parent.DatabaseID)
			case notionapi.ParentTypePageID:
				currentId = string(currentPage.Parent.PageID)
			}
		}
		return breadcrumb
	})
	return breadCrumb
}

func loadMenu(c *gin.Context) interface{} {
	return nil
}

func loadCover(c *gin.Context) []byte {
	domain := c.Param("domain")
	content := cacheClient.ByteLoadAndCache(domain+"/cover", func(fullName string) []byte {
		page := client().SearchForDomain(domain)
		if page == nil {
			return []byte{}
		}
		url := ""
		switch page.Cover.Type {
		case notionapi.FileTypeFile:
			url = page.Cover.File.URL
		case notionapi.FileTypeExternal:
			url = page.Cover.External.URL
		}
		req, _ := http.Get(url)
		buf := bytes.Buffer{}
		io.Copy(&buf, req.Body)
		return buf.Bytes()
	})
	return content
}

type Page struct {
	Page   *notionapi.Page `json:"page"`
	Blocks []lib.TreeBlock `json:"blocks"`
}
