package main

import (
	"Gatewayservice/graph"
	"Gatewayservice/graph/generated"
	"github.com/shurcooL/graphql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		graphql.NewClient("http://localhost:4003/query", nil),
		graphql.NewClient("http://localhost:4001/query", nil),
		graphql.NewClient("http://localhost:4002/query", nil),
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
