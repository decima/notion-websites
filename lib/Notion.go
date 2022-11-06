package lib

import (
	"context"
	"github.com/jomei/notionapi"
	"log"
	"time"
)

type NotionClient struct {
	client     *notionapi.Client
	databaseID notionapi.DatabaseID
}

type TreeBlock struct {
	Type        string          `json:"type"`
	Block       notionapi.Block `json:"block,omitempty"`
	Children    []TreeBlock     `json:"children,omitempty"`
	Database    interface{}     `json:"database,omitempty"`
	HasChildren bool            `json:"has_children"`
}

func NewNotionClient(token string, database string) *NotionClient {
	return &NotionClient{
		client:     notionapi.NewClient(notionapi.Token(token)),
		databaseID: notionapi.DatabaseID(database),
	}
}

func (n *NotionClient) StoreInDatabase(id string, content map[string]interface{}) (*notionapi.Page, error) {
	db := n.ReadDatabase(notionapi.BlockID(id))
	props := notionapi.Properties{}
	for key, element := range content {
		settings := db.Properties[key]
		switch settings.GetType() {
		case notionapi.PropertyConfigTypeTitle:
			props[key] = notionapi.TitleProperty{
				Type: notionapi.PropertyTypeTitle,
				Title: []notionapi.RichText{{
					Type: notionapi.ObjectTypeText,
					Text: &notionapi.Text{
						Content: element.(string)},
				}},
			}
		case notionapi.PropertyConfigTypeRichText:
			props[key] = notionapi.RichTextProperty{
				Type: notionapi.PropertyTypeRichText,
				RichText: []notionapi.RichText{{
					Type: notionapi.ObjectTypeText,
					Text: &notionapi.Text{
						Content: element.(string)},
				}},
			}
		case notionapi.PropertyConfigTypeNumber:
			props[key] = notionapi.NumberProperty{
				Type:   notionapi.PropertyTypeNumber,
				Number: element.(float64),
			}
		case notionapi.PropertyConfigTypeSelect:
			props[key] = notionapi.SelectProperty{
				Type: notionapi.PropertyTypeSelect,
				Select: notionapi.Option{
					ID:   notionapi.PropertyID(element.(string)),
					Name: "undefined",
				},
			}
		case notionapi.PropertyConfigTypeMultiSelect:
			options := []notionapi.Option{}
			for _, v := range element.([]interface{}) {
				options = append(options, notionapi.Option{
					ID:   notionapi.PropertyID(v.(string)),
					Name: "undefined",
				})
			}
			props[key] = notionapi.MultiSelectProperty{
				Type:        notionapi.PropertyTypeMultiSelect,
				MultiSelect: options,
			}
		case notionapi.PropertyConfigTypeDate:
			dt, _ := time.Parse("2006-01-02T15:04", element.(string))
			startDate := notionapi.Date(dt)
			props[key] = notionapi.DateProperty{
				Type: notionapi.PropertyTypeDate,
				Date: &notionapi.DateObject{
					Start: &startDate,
				},
			}
		case notionapi.PropertyConfigTypeCheckbox:
			props[key] = notionapi.CheckboxProperty{
				Type:     notionapi.PropertyTypeCheckbox,
				Checkbox: element.(bool),
			}
		case notionapi.PropertyConfigTypeURL:
			props[key] = notionapi.URLProperty{
				Type: notionapi.PropertyTypeURL,
				URL:  element.(string),
			}
		case notionapi.PropertyConfigTypeEmail:
			props[key] = notionapi.EmailProperty{
				Type:  notionapi.PropertyTypeEmail,
				Email: element.(string),
			}
		case notionapi.PropertyConfigTypePhoneNumber:
			props[key] = notionapi.PhoneNumberProperty{
				Type:        notionapi.PropertyTypePhoneNumber,
				PhoneNumber: element.(string),
			}
		}
	}
	item := notionapi.PageCreateRequest{
		Parent: notionapi.Parent{
			Type:       notionapi.ParentTypeDatabaseID,
			DatabaseID: notionapi.DatabaseID(id),
		},
		Properties: props,
	}
	return n.client.Page.Create(context.Background(), &item)
}

func (n *NotionClient) GetPage(pageId string) *notionapi.Page {
	page, _ := n.client.Page.Get(context.Background(), notionapi.PageID(pageId))
	return page
}

func (n *NotionClient) SearchForDomain(domain string) *notionapi.Page {
	propertyFilter := notionapi.PropertyFilter{
		Property: "domain",
		RichText: &notionapi.TextFilterCondition{
			Equals: domain,
		},
	}

	var cursor notionapi.Cursor
	dbQuery := notionapi.DatabaseQueryRequest{
		Filter:      &propertyFilter,
		PageSize:    100,
		StartCursor: cursor,
	}
	pages, err := n.client.Database.Query(context.Background(), n.databaseID, &dbQuery)
	if err != nil {
		panic(err)
	}
	if len(pages.Results) != 1 {
		log.Println("expect 1 page to match, got ")
		return nil
	}
	return &pages.Results[0]

}

func (n *NotionClient) ReadBlock(id notionapi.BlockID) []TreeBlock {
	var startCursor notionapi.Cursor
	var treeBlocks []TreeBlock

	for {
		pageChildrenBlocks, _ := n.client.Block.GetChildren(context.Background(), id, &notionapi.Pagination{PageSize: 1000, StartCursor: startCursor})
		for _, b := range pageChildrenBlocks.Results {
			block := b.(notionapi.Block)
			treeB := TreeBlock{Type: block.GetType().String(), Block: block, Children: []TreeBlock{}, HasChildren: block.GetHasChildren()}
			if block.GetType() == notionapi.BlockTypeChildDatabase {
				treeB.Database = n.ReadDatabase(block.GetID())
			}
			if block.GetHasChildren() {
				treeB.Children = n.ReadBlock(block.GetID())
			}
			treeBlocks = append(treeBlocks, treeB)
		}
		if !pageChildrenBlocks.HasMore {
			break
		}
		startCursor = notionapi.Cursor(pageChildrenBlocks.NextCursor)
	}
	return treeBlocks
}

func (n *NotionClient) ReadDatabase(id notionapi.BlockID) *notionapi.Database {
	res2, err := n.client.Database.Get(context.Background(), notionapi.DatabaseID(id))
	if err != nil {
		log.Println(err)
	}
	return res2
}
