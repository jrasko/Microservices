package graph

import "Shopservice/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Shops    []*model.Shop
	Products []*model.Product
}
