package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"Gatewayservice/graph/generated"
	"Gatewayservice/graph/model"
	"context"
	"fmt"
)

func (r *queryResolver) AllCustomers(ctx context.Context) ([]*model.Customer, error) {
	query := struct {
		customers []*model.Customer
	}{}
	err := r.CustomerCli.Query(ctx, &query, nil)
	return query.customers, err
}

func (r *queryResolver) Customer(ctx context.Context, id string) (*model.Customer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllOrders(ctx context.Context) ([]*model.Order, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Order(ctx context.Context, id string) (*model.Order, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllProducts(ctx context.Context) ([]*model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllShops(ctx context.Context) ([]*model.Shop, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Product(ctx context.Context, id string) (*model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Shop(ctx context.Context, id string) (*model.Shop, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
