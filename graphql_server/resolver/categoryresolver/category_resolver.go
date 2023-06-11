package categoryresolver

import (
	"encoding/json"
	"errors"
	"fmt"
	"graphqlserver/httpclient"

	"github.com/graphql-go/graphql"
)

type categoryResolver struct {
}

func NewCategoryResolver() CategoryResolver {
	return &categoryResolver{}
}

func (obj categoryResolver) AllCategories() *graphql.Field {
	categoryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Category",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"parentId": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"slug": &graphql.Field{
				Type: graphql.String,
			},
			"image": &graphql.Field{
				Type: graphql.String,
			},
			"active": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	})

	return &graphql.Field{
		Type: graphql.NewList(categoryType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			categories := []Category{}

			resp, err := httpclient.HttpGet("http://localhost:5000/AllCategories")

			if err != nil {
				return nil, errors.New("Unexpected error")
			}

			err = json.Unmarshal([]byte(resp), &categories)

			if err != nil {
				return nil, errors.New("Unexpected error")
			}

			fmt.Printf("[Categories]: %v\n\n", categories)

			return categories, nil
		},
	}
}
