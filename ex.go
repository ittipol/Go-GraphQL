package main

import (
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

type Resp struct {
	Data  Val `json:"data"`
	Data2 struct {
		Hello string `json:"hello"`
	} `json:"data2"`
}

type Val struct {
	Hello string `json:"hello"`
}

func ex() {

	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "World", nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		panic(fmt.Sprintf("Failed to create new GraphQL schema, [%v]", err))
	}

	query := `
		{hello}
	`

	params := graphql.Params{Schema: schema, RequestString: query}
	resp := graphql.Do(params)

	if len(resp.Errors) > 0 {
		panic(fmt.Sprintf("Failed to execute graphql operation, [%v]", resp.Errors))
	}

	fmt.Printf("%v \n\n", resp.Errors)
	fmt.Printf("%v \n\n", resp.Data)

	respJson, err := json.Marshal(resp)

	fmt.Printf("%s \n\n", respJson)

	r := Resp{}

	json.Unmarshal(respJson, &r)

	fmt.Printf("%#v \n\n", r.Data.Hello)

}
