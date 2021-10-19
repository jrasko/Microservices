package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"Customerservice/graph/generated"
	"Customerservice/graph/model"
	"context"
)

func (r *entityResolver) FindCustomerByID(_ context.Context, id string) (*model.Customer, error) {
	for _, customer := range r.Customers {
		if customer.ID == id {
			return customer, nil
		}
	}
	return nil, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
