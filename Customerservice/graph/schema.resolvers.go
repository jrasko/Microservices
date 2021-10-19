package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"Customerservice/graph/generated"
	"Customerservice/graph/model"
	"context"
)

func (r *queryResolver) AllCustomers(ctx context.Context) ([]*model.Customer, error) {
	return r.Customers, nil
}

func (r *queryResolver) Customer(ctx context.Context, id string) (*model.Customer, error) {
	for _, customer := range r.Customers {
		if customer.ID == id {
			return customer, nil
		}
	}
	return nil, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
