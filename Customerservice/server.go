package main

import (
	"Customerservice/graph"
	"Customerservice/graph/generated"
	"Customerservice/graph/model"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "4003"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	var customers = []*model.Customer{
		{
			ID:        "00001",
			Username:  "askywa",
			Firstname: "Annikin",
			Lastname:  "Skywalker",
		},
		{
			ID:        "00002",
			Username:  "lskywa",
			Firstname: "Luke",
			Lastname:  "Skywalker",
		},
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{
			Customers: customers,
		}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
