package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Caicrs/Payfood-backend/graph/generated"
	"github.com/Caicrs/Payfood-backend/graph/model"
	"github.com/Caicrs/Payfood-backend/graph/store"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	db := store.GetStoreFromContext(ctx)
	result := db.AddUser(&input)
	if result != nil {
		return nil, result
	}
	return &model.User{
		ID:          "1",
		Name:        &input.Name,
		Description: &input.Description,
		Cnpj:        &input.Cnpj,
		Email:       &input.Email,
		Password:    &input.Password,
		Image:       &input.Image,
	}, nil
}

// CreateProduct is the resolver for the createProduct field.
func (r *mutationResolver) CreateProduct(ctx context.Context, input model.NewProduct) (*model.Product, error) {
	db := store.GetStoreFromContext(ctx)
	result := db.AddProduct(&input)
	if result != nil {
		return nil, result
	}
	return &model.Product{
		ID:            "1",
		MarketplaceID: &input.MarketplaceID,
		Name:          &input.Name,
		Description:   &input.Description,
		Price:         &input.Price,
	}, nil
}

// CreateTable is the resolver for the createTable field.
func (r *mutationResolver) CreateTable(ctx context.Context, input model.NewTable) (*model.Table, error) {
	db := store.GetStoreFromContext(ctx)
	result := db.AddTable(&input)
	if result != nil {
		return nil, result
	}
	return &model.Table{
		ID:     "1",
		Number: &input.Number,
		Qrcode: &input.Qrcode,
	}, nil
}

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, input model.NewOrder) (*model.Order, error) {
	db := store.GetStoreFromContext(ctx)
	result := db.AddOrder(&input)
	if result != nil {
		return nil, result
	}
	return &model.Order{
		ID:                "1",
		MarketplaceUserID: &input.MarketplaceUserID,
		ClientID:          &input.ClientID,
		TableID:           &input.TableID,
		TotalPrice:        &input.TotalPrice,
	}, nil
}

// CreateFidelity is the resolver for the createFidelity field.
func (r *mutationResolver) CreateFidelity(ctx context.Context, input model.NewFidelity) (*model.Fidelity, error) {
	db := store.GetStoreFromContext(ctx)
	result := db.AddFidelity(&input)
	if result != nil {
		return nil, result
	}
	return &model.Fidelity{
		ID:       "1",
		ClientID: &input.ClientID,
		Score:    &input.Score,
	}, nil
}

// CreateClient is the resolver for the createClient field.
func (r *mutationResolver) CreateClient(ctx context.Context, input model.NewClient) (*model.Client, error) {
	db := store.GetStoreFromContext(ctx)
	err := db.AddClient(&input)
	if err != nil {
		return nil, err
	}

	return &model.Client{
		ID:       "1",
		Name:     &input.Name,
		Email:    &input.Email,
		Password: &input.Password,
		Cpf:      &input.Cpf,
		Phone:    &input.Phone,
	}, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	db := store.GetStoreFromContext(ctx)
	return db.Users, nil
}

// Tables is the resolver for the tables field.
func (r *queryResolver) Tables(ctx context.Context) ([]*model.Table, error) {
	db := store.GetStoreFromContext(ctx)
	return db.Tables, nil
}

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	db := store.GetStoreFromContext(ctx)
	return db.Products, nil
}

// Clients is the resolver for the clients field.
func (r *queryResolver) Clients(ctx context.Context) ([]*model.Client, error) {
	db := store.GetStoreFromContext(ctx)
	return db.Clients, nil
}

// Orders is the resolver for the orders field.
func (r *queryResolver) Orders(ctx context.Context) ([]*model.Order, error) {
	db := store.GetStoreFromContext(ctx)
	return db.Orders, nil
}

// Fidelity is the resolver for the fidelity field.
func (r *queryResolver) Fidelity(ctx context.Context) ([]*model.Fidelity, error) {
	db := store.GetStoreFromContext(ctx)
	return db.Fidelitys, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
