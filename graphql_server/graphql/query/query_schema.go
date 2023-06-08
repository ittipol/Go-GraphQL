package query

import (
	"graphqlserver/graphql/query/queryfields"

	"github.com/graphql-go/graphql"
)

func Schema() *graphql.Object {
	fields := graphql.Fields{
		"allItems": queryfields.AllItems(),
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}

	return graphql.NewObject(rootQuery)
}
