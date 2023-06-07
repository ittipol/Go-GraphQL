package queryfields

import (
	"errors"
	"graphqlserver/httpclient"

	"github.com/graphql-go/graphql"
)

func User() *graphql.Field {
	return &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			resp, err := httpclient.HttpGet("http://localhost:5000/users")

			if err != nil {
				return "", errors.New("Cannot Get Users")
			}

			return resp, nil
		},
	}
}
