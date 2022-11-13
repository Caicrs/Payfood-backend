package graph
<<<<<<< HEAD
=======

import (
	"context"
	"fmt"

	"time"

	"github.com/Caicrs/Payfood-backend/common"
	"github.com/Caicrs/Payfood-backend/graph/generated"
	model "github.com/Caicrs/Payfood-backend/graph/model"
	uuid "github.com/satori/go.uuid"
)

// CreateMarketplace is the resolver for the createMarketplace field.
func (r *mutationResolver) CreateMarketplace(ctx context.Context, input model.NewMarketplace) (*model.Marketplaces, error) {
	context := common.GetContext(ctx)
	marketplace := model.Marketplaces{
		ID:          fmt.Sprintln(uuid.NewV4()),
		Name:        input.Name,
		Description: input.Description,
		Cnpj:        input.Cnpj,
		Email:       input.Email,
		Password:    input.Password,
		Version:     input.Version,
	}
	err := context.Database.Create(&marketplace).Error
	if err != nil {
		return nil, err
	}
	return &marketplace, nil
}

// GetMarketplaces is the resolver for the getMarketplaces field.
func (r *queryResolver) GetMarketplaces(ctx context.Context) ([]*model.Marketplaces, error) {
	panic(fmt.Errorf("not implemented: GetMarketplaces - getMarketplaces"))
}

// GetOneMarketplace is the resolver for the getOneMarketplace field.
func (r *queryResolver) GetOneMarketplace(ctx context.Context, id string) (*model.Marketplaces, error) {
	panic(fmt.Errorf("not implemented: GetOneMarketplace - getOneMarketplace"))
}

// UpdateMarketplace is the resolver for the updateMarketplace field.
func (r *mutationResolver) UpdateMarketplace(ctx context.Context, input model.UpdateMarketplace, id string) (*model.Marketplaces, error) {
	panic(fmt.Errorf("not implemented: UpdateMarketplace - updateMarketplace"))
}

// DeleteMarketplace is the resolver for the deleteMarketplace field.
func (r *mutationResolver) DeleteMarketplace(ctx context.Context, id string) (*model.Marketplaces, error) {
	panic(fmt.Errorf("not implemented: DeleteMarketplace - deleteMarketplace"))
}

// CreateProduct is the resolver for the createProduct field.
func (r *mutationResolver) CreateProduct(ctx context.Context, input model.NewProduct) (*model.Products, error) {
	context := common.GetContext(ctx)
	product := model.Products{
		ID:            fmt.Sprintln(uuid.NewV4()),
		Name:          input.Name,
		Description:   input.Description,
		Price:         input.Price,
		Images:        input.Images,
		MarketplaceID: input.MarketplaceID,
	}
	err := context.Database.Create(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// UpdateProduct is the resolver for the updateProduct field.
func (r *mutationResolver) UpdateProduct(ctx context.Context, input model.UpdateProduct, id string) (*model.Products, error) {
	panic(fmt.Errorf("not implemented: UpdateProduct - updateProduct"))
}

// DeleteProduct is the resolver for the deleteProduct field.
func (r *mutationResolver) DeleteProduct(ctx context.Context, id string) (*model.Products, error) {
	panic(fmt.Errorf("not implemented: DeleteProduct - deleteProduct"))
}

// GetProducts is the resolver for the getProducts field.
func (r *queryResolver) GetProducts(ctx context.Context) ([]*model.Products, error) {
	panic(fmt.Errorf("not implemented: GetProducts - getProducts"))
}

// GetOneProduct is the resolver for the getOneProduct field.
func (r *queryResolver) GetOneProduct(ctx context.Context, id string) (*model.Products, error) {
	panic(fmt.Errorf("not implemented: GetOneProduct - getOneProduct"))
}

// GetClients is the resolver for the getClients field.
func (r *queryResolver) GetClients(ctx context.Context) ([]*model.Clients, error) {
	panic(fmt.Errorf("not implemented: GetOneClient - getOneClient"))
}

// GetOneClient is the resolver for the getOneClient field.
func (r *queryResolver) GetOneClient(ctx context.Context, id string) (*model.Clients, error) {
	panic(fmt.Errorf("not implemented: GetOneClient - getOneClient"))
}

// CreateClient is the resolver for the createClient field.
func (r *mutationResolver) CreateClient(ctx context.Context, input model.NewClient) (*model.Clients, error) {
	context := common.GetContext(ctx)
	clients := model.Clients{
		ID:       fmt.Sprintln(uuid.NewV4()),
		Name:     input.Name,
		Phone:    input.Phone,
		Cpf:      input.Cpf,
		Email:    input.Email,
		Password: input.Password,
		MiniPass: input.MiniPass,
	}
	err := context.Database.Create(&clients).Error
	if err != nil {
		return nil, err
	}
	return &clients, nil
}

// UpdateClient is the resolver for the updateClient field.
func (r *mutationResolver) UpdateClient(ctx context.Context, input model.UpdateClient, id string) (*model.Clients, error) {
	panic(fmt.Errorf("not implemented: UpdateClient - updateClient"))
}

// DeleteClient is the resolver for the deleteClient field.
func (r *mutationResolver) DeleteClient(ctx context.Context, id string) (*model.Clients, error) {
	panic(fmt.Errorf("not implemented: DeleteClient - deleteClient"))
}

// GetOrders is the resolver for the getOrders field.
func (r *queryResolver) GetOrders(ctx context.Context) ([]*model.Orders, error) {
	panic(fmt.Errorf("not implemented: GetOrders - getOrders"))
}

// GetOneOrder is the resolver for the getOneOrder field.
func (r *queryResolver) GetOneOrder(ctx context.Context, id string) (*model.Orders, error) {
	panic(fmt.Errorf("not implemented: GetOneOrder - getOneOrder"))
}

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, input model.NewOrder) (*model.Orders, error) {
	context := common.GetContext(ctx)
	clients := model.Orders{
		ID:            fmt.Sprintln(uuid.NewV4()),
		ClientID:      input.ClientID,
		MarketplaceID: input.MarketplaceID,
		ProductsID:    input.ProductsID,
		TableID:       input.TableID,
		TotalPrice:    input.TotalPrice,
		CreatedAt:     fmt.Sprintln(time.Now()),
	}
	err := context.Database.Create(&clients).Error
	if err != nil {
		return nil, err
	}
	return &clients, nil
}

// UpdateOrder is the resolver for the updateOrder field.
func (r *mutationResolver) UpdateOrder(ctx context.Context, input model.UpdateOrder, id string) (*model.Orders, error) {
	panic(fmt.Errorf("not implemented: UpdateOrder - updateOrder"))
}

// DeleteOrder is the resolver for the deleteOrder field.
func (r *mutationResolver) DeleteOrder(ctx context.Context, id string) (*model.Orders, error) {
	panic(fmt.Errorf("not implemented: DeleteOrder - deleteOrder"))
}

// CreateTable is the resolver for the createTable field.
func (r *mutationResolver) CreateTable(ctx context.Context, input model.NewTable) (*model.Tables, error) {
	panic(fmt.Errorf("not implemented: CreateTable - createTable"))
}

// UpdateTable is the resolver for the updateTable field.
func (r *mutationResolver) UpdateTable(ctx context.Context, input model.UpdateTable, id string) (*model.Tables, error) {
	panic(fmt.Errorf("not implemented: UpdateTable - updateTable"))
}

// DeleteTable is the resolver for the deleteTable field.
func (r *mutationResolver) DeleteTable(ctx context.Context, id string) (*model.Tables, error) {
	panic(fmt.Errorf("not implemented: DeleteTable - deleteTable"))
}

// GetTables is the resolver for the getTables field.
func (r *queryResolver) GetTables(ctx context.Context) ([]*model.Tables, error) {
	panic(fmt.Errorf("not implemented: GetTables - getTables"))
}

// GetOneTable is the resolver for the getOneTable field.
func (r *queryResolver) GetOneTable(ctx context.Context, id string) (*model.Tables, error) {
	panic(fmt.Errorf("not implemented: GetOneTable - getOneTable"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
>>>>>>> main
