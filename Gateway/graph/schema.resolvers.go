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
		AllCustomers []model.Customer
	}{}
	err := r.CustomerCli.Query(ctx, &query, nil)
	var ret []*model.Customer
	for _, customer := range query.AllCustomers {
		ret = append(ret, &customer)
	}
	return ret, err
}

func (r *queryResolver) Customer(ctx context.Context, id string) (*model.Customer, error) {
	query := struct {
		Customer model.Customer `graphql:"customer(id: $id)"`
	}{}
	err := r.CustomerCli.Query(ctx, &query, map[string]interface{}{"id": id})
	return &query.Customer, err
}

func (r *queryResolver) AllOrders(ctx context.Context) ([]*model.Order, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Order(ctx context.Context, id string) (*model.Order, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllProducts(ctx context.Context) ([]*model.Product, error) {
	query := struct {
		AllProducts []model.Product
	}{}
	err := r.ShopCli.Query(ctx, &query, nil)
	var ret []*model.Product
	for _, product := range query.AllProducts {
		ret = append(ret, &product)
	}
	return ret, err
}

func (r *queryResolver) AllShops(ctx context.Context) ([]*model.Shop, error) {
	query := struct {
		AllShops []model.Shop
	}{}
	err := r.ShopCli.Query(ctx, &query, nil)
	var ret []*model.Shop
	for _, shop := range query.AllShops {
		ret = append(ret, &shop)
	}
	return ret, err
}

func (r *queryResolver) Product(ctx context.Context, id string) (*model.Product, error) {
	query := struct {
		Product model.Product `graphql:"product(id: $id)"`
	}{}
	err := r.ShopCli.Query(ctx, &query, map[string]interface{}{"id": id})
	return &query.Product, err
}

func (r *queryResolver) Shop(ctx context.Context, id string) (*model.Shop, error) {
	query := struct {
		Shop model.Shop `graphql:"shop(id: $id)"`
	}{}
	err := r.ShopCli.Query(ctx, &query, map[string]interface{}{"id": id})
	return &query.Shop, err
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
