package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"Shopservice/graph/generated"
	"Shopservice/graph/model"
	"context"
)

func (r *entityResolver) FindProductByID(_ context.Context, id string) (*model.Product, error) {
	for _, product := range r.Products {
		if product.ID == id {
			return product, nil
		}
	}
	return nil, nil
}

func (r *entityResolver) FindShopByID(_ context.Context, id string) (*model.Shop, error) {
	for _, shop := range r.Shops {
		if shop.ID == id {
			return shop, nil
		}
	}
	return nil, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
