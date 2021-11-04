package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"Shopservice/graph/generated"
	"Shopservice/graph/model"
	"context"
)

func (r *queryResolver) AllProducts(ctx context.Context) ([]*model.Product, error) {
	return r.Products, nil
}

func (r *queryResolver) AllShops(ctx context.Context) ([]*model.Shop, error) {
	return r.Shops, nil
}

func (r *queryResolver) Product(ctx context.Context, id string) (*model.Product, error) {
	for _, product := range r.Products {
		if product.ID == id {
			return product, nil
		}
	}
	return nil, nil
}

func (r *queryResolver) Shop(ctx context.Context, id string) (*model.Shop, error) {

	for _, shop := range r.Shops {
		if shop.ID == id {
			return shop, nil
		}
	}
	return nil, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
