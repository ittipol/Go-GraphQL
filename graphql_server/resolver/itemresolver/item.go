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

func NewItemResolver() ItemResolver {
	return &itemResolver{}
}

func (obj itemResolver) AllItems() *graphql.Field {

	itemType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Item",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"price": &graphql.Field{
				Type: graphql.Float,
			},
			"originalPrice": &graphql.Field{
				Type: graphql.Float,
			},
		},
	})

	return &graphql.Field{
		Type: graphql.NewList(itemType),
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

			req := itemRequest{
				Title: title,
				Price: price,
			}

			fmt.Printf("%v, %v \n\n", title, price)

			body, err := json.Marshal(req)

			if err != nil {
				return nil, errors.New("Unexpected error")
			}

			_, err = httpclient.HttpPost("http://localhost:5000/addItem", "application/json", body)

			if err != nil {
				return "", errors.New("Unexpected error")
			}

			return "OK", nil
		},
	}
}
