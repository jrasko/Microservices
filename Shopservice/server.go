package main

import (
	"Shopservice/graph"
	"Shopservice/graph/generated"
	"Shopservice/graph/model"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "4002"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	shops, products := initialData()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		Shops:    shops,
		Products: products,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func initialData() ([]*model.Shop, []*model.Product) {
	var descriptions = [2]string{"Rule the Galaxy", "Pew Pew"}
	var locations = [2]string{"Geonosis", "Alderaan"}
	shops := []*model.Shop{
		{
			ID:       "00001",
			Name:     "Imperium",
			Location: &locations[0],
			Products: nil,
		},
		{
			ID:       "00002",
			Name:     "Alliance",
			Location: &locations[1],
			Products: nil,
		},
	}
	products := []*model.Product{
		{
			ID:          "00001",
			Name:        "Death Star",
			Description: &descriptions[0],
			Shops:       nil,
			Image:       "",
		},
		{
			ID:          "00002",
			Name:        "TIE-Fighter",
			Description: nil,
			Shops:       nil,
			Image:       "",
		},
		{
			ID:          "00003",
			Name:        "X-Wing",
			Description: &descriptions[1],
			Shops:       nil,
			Image:       "",
		},
	}
	shopproducts := [][]*model.ProductInShop{
		// Products from Imperium
		{
			{
				Product:   products[0],
				Shop:      shops[0],
				Price:     1_500_000_000_000,
				Currency:  "Credits",
				SalesUnit: 1,
			},
			{
				Product:   products[1],
				Shop:      shops[0],
				Price:     100_000,
				Currency:  "Credits",
				SalesUnit: 5,
			},
		},
		//Products from Alliance
		{
			{
				Product:   products[2],
				Shop:      shops[1],
				Price:     5_000_000,
				Currency:  "Credits",
				SalesUnit: 1,
			},
		},
		// Shops with Deathstar
		{
			{
				Product:   products[0],
				Shop:      shops[0],
				Price:     1_500_000_000_000,
				Currency:  "Credits",
				SalesUnit: 1,
			},
		},
		// Shops with TIE Fighter
		{
			{
				Product:   products[1],
				Shop:      shops[0],
				Price:     100_000,
				Currency:  "Credits",
				SalesUnit: 5,
			},
		},
		// Shops with X-Wing
		{
			{
				Product:   products[2],
				Shop:      shops[1],
				Price:     5_000_000,
				Currency:  "Credits",
				SalesUnit: 1,
			},
		},
	}
	shops[0].Products = shopproducts[0]
	shops[1].Products = shopproducts[1]

	products[0].Shops = shopproducts[2]
	products[1].Shops = shopproducts[3]
	products[2].Shops = shopproducts[4]

	return shops, products
}
