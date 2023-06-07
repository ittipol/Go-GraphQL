package mutationfields

import (
	"errors"
	"fmt"
	"graphqlserver/httpclient"

	"github.com/graphql-go/graphql"
)

func AddUser() *graphql.Field {
	return &graphql.Field{
		Type: graphql.String,
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"age": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			name := p.Args["name"].(string)
			age := p.Args["age"].(int)

			fmt.Printf("%v, %v \n\n", name, age)

			str := fmt.Sprintf(`{
				"name": "%v",
				"age": %v
			}`, name, age)

			fmt.Println(str)

			body := []byte(str)

			resp, err := httpclient.HttpPost("http://localhost:5000/user", "application/json", body)

			if err != nil {
				return "", errors.New("Cannot Save User")
			}

			return resp, nil
		},
	}
}
