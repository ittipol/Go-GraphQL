package ex

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

type User struct {
	ID   int
	Name string
}

type Input struct {
	graphql.Input
}

func ex2() {

	users := []User{
		{ID: 1, Name: "AAAAA"},
		{ID: 2, Name: "BBBBB"},
	}

	userType := graphql.NewObject(graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	fields := graphql.Fields{
		"allUsers": &graphql.Field{
			Type: graphql.NewList(userType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return users, nil
			},
		},
		"GetUser": &graphql.Field{
			Type: userType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.Int},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {

				id, ok := p.Args["id"].(int)

				fmt.Printf("Arg Found: [%v]\n\n", ok)

				if !ok {
					return User{}, nil
				}

				return GetUserById(id, users), nil
			},
		},
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "World", nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: fields,
	}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		panic(fmt.Sprintf("Failed to create new GraphQL schema, [%v]", err))
	}

	query := `
		{
			GetUser(id:1){id,name}
		}
	`

	params := graphql.Params{Schema: schema, RequestString: query}
	resp := graphql.Do(params)

	if len(resp.Errors) > 0 {
		panic(fmt.Sprintf("Failed to execute graphql operation, [%v]", resp.Errors))
	}

	fmt.Printf("Error %v \n\n", resp.Errors)
	fmt.Printf("%v \n\n", resp.Data)

}

func GetUserById(id int, users []User) User {

	for _, user := range users {
		if id == user.ID {
			fmt.Printf("Found ID: %v \n\n", id)
			return user
		}
	}

	return User{}
}
