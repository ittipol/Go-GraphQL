package schema

import (
	"graphqlserver/resolver/categoryresolver"
	"graphqlserver/resolver/itemresolver"

	"github.com/graphql-go/graphql"
)

type QuerySchema interface {
	Schema() *graphql.Object
}

type querySchema struct {
	itemResolver     itemresolver.ItemResolver
	categoryResolver categoryresolver.CategoryResolver
}

func NewQuerySchema(
	itemResolver itemresolver.ItemResolver,
	categoryResolver categoryresolver.CategoryResolver,
) QuerySchema {
	return &querySchema{
		itemResolver:     itemResolver,
		categoryResolver: categoryResolver,
	}
}

func (obj querySchema) Schema() *graphql.Object {
	fields := graphql.Fields{
		"allItems":      obj.itemResolver.AllItems(),
		"getItemBySlug": obj.itemResolver.GetItemBySlug(),
		"allCategories": obj.categoryResolver.AllCategories(),
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}

	return graphql.NewObject(rootQuery)
}
