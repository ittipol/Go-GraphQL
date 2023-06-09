package schema

import (
	"graphqlserver/resolver/itemresolver"

	"github.com/graphql-go/graphql"
)

type MutationSchema interface {
	Schema() *graphql.Object
}

type mutationSchema struct {
	itemResolver itemresolver.ItemResolver
}

func NewMutationSchema(itemResolver itemresolver.ItemResolver) MutationSchema {
	return &mutationSchema{itemResolver}
}

func (obj mutationSchema) Schema() *graphql.Object {
	fields := graphql.Fields{
		"addItem": obj.itemResolver.AddItem(),
	}

	rootMutation := graphql.ObjectConfig{Name: "RootMutation", Fields: fields}

	return graphql.NewObject(rootMutation)
}
