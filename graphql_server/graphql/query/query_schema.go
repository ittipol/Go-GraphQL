package query

import (
	"graphqlserver/graphql/query/queryfields"

	"github.com/graphql-go/graphql"
)

func Schema() *graphql.Object {
	fields := graphql.Fields{
		"user": queryfields.User(),
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}

	return graphql.NewObject(rootQuery)
}
