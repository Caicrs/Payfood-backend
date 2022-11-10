package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Caicrs/Payfood-backend/graph/generated"
	"github.com/Caicrs/Payfood-backend/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteUser - deleteUser"))
}

// CreateProduct is the resolver for the createProduct field.
func (r *mutationResolver) CreateProduct(ctx context.Context, input model.NewProduct) (*model.Product, error) {
	panic(fmt.Errorf("not implemented: CreateProduct - createProduct"))
}

// CreateTable is the resolver for the createTable field.
func (r *mutationResolver) CreateTable(ctx context.Context, input model.NewTable) (*model.Table, error) {
	panic(fmt.Errorf("not implemented: CreateTable - createTable"))
}

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, input *model.NewOrder) (*model.Order, error) {
	panic(fmt.Errorf("not implemented: CreateOrder - createOrder"))
}

// CreateFidelity is the resolver for the createFidelity field.
func (r *mutationResolver) CreateFidelity(ctx context.Context, input model.NewFidelity) (*model.Fidelity, error) {
	panic(fmt.Errorf("not implemented: CreateFidelity - createFidelity"))
}

// CreateClient is the resolver for the createClient field.
func (r *mutationResolver) CreateClient(ctx context.Context, input model.NewClient) (*model.Client, error) {
	panic(fmt.Errorf("not implemented: CreateClient - createClient"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// FindUser is the resolver for the findUser field.
func (r *queryResolver) FindUser(ctx context.Context, id string) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: FindUser - findUser"))
}

// Tables is the resolver for the tables field.
func (r *queryResolver) Tables(ctx context.Context) ([]*model.Table, error) {
	panic(fmt.Errorf("not implemented: Tables - tables"))
}

// FindTables is the resolver for the findTables field.
func (r *queryResolver) FindTables(ctx context.Context, id string) ([]*model.Table, error) {
	panic(fmt.Errorf("not implemented: FindTables - findTables"))
}

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	panic(fmt.Errorf("not implemented: Products - products"))
}

// FindProduct is the resolver for the findProduct field.
func (r *queryResolver) FindProduct(ctx context.Context, id string) ([]*model.Product, error) {
	panic(fmt.Errorf("not implemented: FindProduct - findProduct"))
}

// Clients is the resolver for the clients field.
func (r *queryResolver) Clients(ctx context.Context) ([]*model.Client, error) {
	panic(fmt.Errorf("not implemented: Clients - clients"))
}

// FindClient is the resolver for the findClient field.
func (r *queryResolver) FindClient(ctx context.Context, id string) ([]*model.Client, error) {
	panic(fmt.Errorf("not implemented: FindClient - findClient"))
}

// Orders is the resolver for the orders field.
func (r *queryResolver) Orders(ctx context.Context) ([]*model.Order, error) {
	panic(fmt.Errorf("not implemented: Orders - orders"))
}

// FindOrder is the resolver for the findOrder field.
func (r *queryResolver) FindOrder(ctx context.Context, id string) ([]*model.Order, error) {
	panic(fmt.Errorf("not implemented: FindOrder - findOrder"))
}

// Fidelity is the resolver for the fidelity field.
func (r *queryResolver) Fidelity(ctx context.Context) ([]*model.Fidelity, error) {
	panic(fmt.Errorf("not implemented: Fidelity - fidelity"))
}

// FindFidelity is the resolver for the findFidelity field.
func (r *queryResolver) FindFidelity(ctx context.Context, id string) ([]*model.Fidelity, error) {
	panic(fmt.Errorf("not implemented: FindFidelity - findFidelity"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) DeleteUser(ctx context.Context, id string) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: DeleteUser - deleteUser"))
}
