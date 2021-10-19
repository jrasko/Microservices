package main

import (
	"Orderservice/graph"
	"Orderservice/graph/generated"
	"Orderservice/graph/model"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"log"
	"net/http"
	"os"
)

const defaultPort = "4001"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	var orders = []*model.Order{
		{
			ID:       "00001",
			Time:     "21 BBY",
			State:    model.OrderStateDone,
			Shop:     &model.Shop{ID: "00002"},
			Customer: &model.Customer{ID: "00001"},
			Products: nil,
		},
		{
			ID:       "00002",
			Time:     "0 ABY",
			State:    model.OrderStateAccepted,
			Shop:     &model.Shop{ID: "00001"},
			Customer: &model.Customer{ID: "00001"},
			Products: nil,
		},
		{
			ID:       "00003",
			Time:     "2 BBY",
			State:    model.OrderStateInProgress,
			Shop:     &model.Shop{ID: "00002"},
			Customer: &model.Customer{ID: "00002"},
			Products: nil,
		},
	}
	orders[0].Products = append(orders[0].Products, &model.OrderProducts{
		Order:    orders[0],
		Product:  &model.Product{ID: "00001"},
		Quantity: 1,
		Price:    1_000_000_000_000_000,
	})
	orders[0].Products = append(orders[0].Products, &model.OrderProducts{
		Order:    orders[0],
		Product:  &model.Product{ID: "00002"},
		Quantity: 30,
		Price:    3_000_000,
	})
	orders[1].Products = append(orders[0].Products, &model.OrderProducts{
		Order:    orders[1],
		Product:  &model.Product{ID: "00002"},
		Quantity: 5,
		Price:    500_000,
	})
	orders[2].Products = append(orders[0].Products, &model.OrderProducts{
		Order:    orders[2],
		Product:  &model.Product{ID: "00003"},
		Quantity: 2,
		Price:    10_000_000,
	})

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{
			Orders: orders,
		},
	}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
