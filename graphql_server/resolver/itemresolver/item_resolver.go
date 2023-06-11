package itemresolver

import (
	"encoding/json"
	"errors"
	"fmt"
	"graphqlserver/httpclient"

	"github.com/graphql-go/graphql"
)

type itemResolver struct {
	// repository
}

func NewItemResolver( /*repository*/ ) ItemResolver {
	return &itemResolver{ /*repository*/ }
}

func (obj itemResolver) AllItems() *graphql.Field {

	return &graphql.Field{
		Type: graphql.NewList(itemType()),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			items := []item{}

			resp, err := httpclient.HttpGet("http://localhost:5000/allItems")

			if err != nil {
				return nil, errors.New("Unexpected error")
			}

			err = json.Unmarshal([]byte(resp), &items)

			if err != nil {
				return nil, errors.New("Unexpected error")
			}

			fmt.Printf("[Items]: %v\n\n", items)

			return items, nil
		},
	}

}

func (obj itemResolver) GetItemBySlug() *graphql.Field {

	return &graphql.Field{
		Type: itemType(),
		Args: graphql.FieldConfigArgument{
			"slug": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			item := item{}

			slug := p.Args["slug"].(string)

			fmt.Printf("%v \n\n", slug)

			resp, err := httpclient.HttpGet(fmt.Sprintf("http://localhost:5000/getItemBySlug/%v", slug))

			if err != nil {
				return nil, errors.New("Unexpected error")
			}

			err = json.Unmarshal([]byte(resp), &item)

			if err != nil {
				return nil, errors.New("Unexpected error")
			}

			return item, nil
		},
	}

}

func (obj itemResolver) AddItem() *graphql.Field {

	return &graphql.Field{
		Type: graphql.String,
		Args: graphql.FieldConfigArgument{
			"title": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"price": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			title := p.Args["title"].(string)
			price := p.Args["price"].(float64)

			// Create request body
			req := addItemRequest{
				Title: title,
				Price: price,
			}

			fmt.Printf("%v, %v \n\n", title, price)

			body, err := json.Marshal(req)

			if err != nil {
				return nil, errors.New("Unexpected error")
			}

			resp, err := httpclient.HttpPost("http://localhost:5000/addItem", "application/json", body)

			if err != nil {
				return "", errors.New("Unexpected error")
			}

			return resp, nil
		},
	}

}
