package queryfields

import (
	"encoding/json"
	"errors"
	"fmt"
	"graphqlserver/httpclient"

	"github.com/graphql-go/graphql"
)

type user struct {
	Id            int
	Title         string
	Price         float32
	OriginalPrice float32
}

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
			"originalPrice": &graphql.Field{
				Type: graphql.Float,
			},
		},
	})

	return &graphql.Field{
		Type: graphql.NewList(itemType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			users := []user{}

			resp, err := httpclient.HttpGet("http://localhost:5000/allItems")

			if err != nil {
				return nil, errors.New("Unexpected error")
			}

			err = json.Unmarshal([]byte(resp), &users)

			if err != nil {
				return nil, errors.New("Unexpected error")
			}

			fmt.Printf("%v\n\n", users)

			return users, nil
		},
	}
}
