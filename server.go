//go:generate go run github.com/99designs/gqlgen
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/mt-st1/graphql-todos/domains"
	"github.com/mt-st1/graphql-todos/graph"
	"github.com/mt-st1/graphql-todos/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					TodoRepository: domains.NewTodoRepository(),
				},
			},
		),
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
