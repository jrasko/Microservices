package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"Orderservice/graph/generated"
	"Orderservice/graph/model"
	"context"
)

func (r *queryResolver) AllOrders(ctx context.Context) ([]*model.Order, error) {
	return r.Orders, nil
}

func (r *queryResolver) Order(ctx context.Context, id string) (*model.Order, error) {
	for _, order := range r.Orders {
		if order.ID == id {
			return order, nil
		}
	}
	return nil, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
