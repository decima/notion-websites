package Cache

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"
)

type AbstractCache interface {
	Store(string, []byte) error
	Get(string) ([]byte, error)
	Clear(string) error
}

type Cache struct {
	cacheClient AbstractCache
}

func NewCache(cacheURL string, ttl int) *Cache {
	u, err := url.Parse(cacheURL)
	if err != nil {
		panic(err)
	}
	var cacheClient AbstractCache
	switch u.Scheme {
	case "redis":
		password, _ := u.User.Password()
		database, _ := strconv.Atoi(strings.TrimPrefix(u.Path, "/"))
		cacheClient = NewRedisCache(u.Host, password, database, ttl)
	}
	return &Cache{cacheClient: cacheClient}
}

func (c *Cache) Store(domain string, value interface{}) error {
	content, _ := json.Marshal(value)
	return c.cacheClient.Store(domain, content)
}

func (c *Cache) Clear(domain string) error {
	return c.cacheClient.Clear(domain)
}

func (c *Cache) Retrieve(domain string) (map[string]interface{}, error) {
	content, err := c.cacheClient.Get(domain)
	if err != nil {
		return nil, err
	}
	var resultMap map[string]interface{}
	err = json.Unmarshal(content, &resultMap)
	return resultMap, nil
}

func (c *Cache) ByteRetrieve(domain string) ([]byte, error) {
	return c.cacheClient.Get(domain)

}

func (c *Cache) LoadAndCache(domain string, f func(domain string) interface{}) interface{} {
	content, err := c.Retrieve(domain)
	if err == nil && content != nil {
		return content
	}

	result := f(domain)

	c.Store(domain, result)
	return result
}

func (c *Cache) ByteLoadAndCache(domain string, f func(domain string) []byte) []byte {
	content, err := c.ByteRetrieve(domain)
	if err == nil && content != nil {
		var b []byte
		json.Unmarshal(content, &b)
		return b
	}

	result := f(domain)

	c.Store(domain, result)
	return result
}
