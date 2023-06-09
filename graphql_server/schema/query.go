package schema

import (
	"graphqlserver/resolver/itemresolver"

	"github.com/graphql-go/graphql"
)

type QuerySchema interface {
	Schema() *graphql.Object
}

type querySchema struct {
	itemResolver itemresolver.ItemResolver
}

func NewQuerySchema(itemResolver itemresolver.ItemResolver) QuerySchema {
	return &querySchema{itemResolver}
}

func (obj querySchema) Schema() *graphql.Object {
	fields := graphql.Fields{
		"allItems": obj.itemResolver.AllItems(),
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}

	return graphql.NewObject(rootQuery)
}
