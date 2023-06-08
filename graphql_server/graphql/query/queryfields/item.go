package queryfields

import (
	"errors"
	"graphqlserver/httpclient"

	"github.com/graphql-go/graphql"
)

func AllItems() *graphql.Field {

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
		},
	})

	return &graphql.Field{
		Type: itemType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			resp, err := httpclient.HttpGet("http://localhost:5000/allItems")

			if err != nil {
				return "", errors.New("Failed to get items")
			}

			return resp, nil
		},
	}
}
