package mutation

import (
	"graphqlserver/graphql/mutation/mutationfields"

	"github.com/graphql-go/graphql"
)

func Schema() *graphql.Object {

	fields := graphql.Fields{
		"addUser": mutationfields.AddUser(),
	}

	rootMutation := graphql.ObjectConfig{Name: "RootMutation", Fields: fields}

	return graphql.NewObject(rootMutation)
}
