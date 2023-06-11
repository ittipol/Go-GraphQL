package main

import (
	"fmt"

	"graphqlserver/resolver/categoryresolver"
	"graphqlserver/resolver/itemresolver"
	"graphqlserver/schema"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {

	// Resolver
	itemResolver := itemresolver.NewItemResolver()
	categoryResolver := categoryresolver.NewCategoryResolver()

	// Query
	querySchema := schema.NewQuerySchema(itemResolver, categoryResolver)
	// Mutation
	mutationSchema := schema.NewMutationSchema(itemResolver)

	schemaConfig := graphql.SchemaConfig{
		Query:    querySchema.Schema(),
		Mutation: mutationSchema.Schema(),
	}
	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		panic(fmt.Sprintf("Failed to create new GraphQL schema, [%v]", err))
	}

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: false,
	})

	http.Handle("/graphql", h)

	var sandboxHTML = []byte(`
		<!DOCTYPE html>
		<html lang="en">
		<body style="margin: 0; overflow-x: hidden; overflow-y: hidden">
		<div id="sandbox" style="height:100vh; width:100vw;"></div>
		<script src="https://embeddable-sandbox.cdn.apollographql.com/_latest/embeddable-sandbox.umd.production.min.js"></script>
		<script>
		new window.EmbeddedSandbox({
		target: "#sandbox",
		// Pass through your server href if you are embedding on an endpoint.
		// Otherwise, you can pass whatever endpoint you want Sandbox to start up with here.
		initialEndpoint: "http://localhost:8080/graphql",
		});
		// advanced options: https://www.apollographql.com/docs/studio/explorer/sandbox#embedding-sandbox
		</script>
		</body>
		
		</html>`)

	http.Handle("/sandbox", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(sandboxHTML)
	}))

	err = http.ListenAndServe(":8080", nil)

	if err != nil {
		panic("Failed to start server")
	}
}
