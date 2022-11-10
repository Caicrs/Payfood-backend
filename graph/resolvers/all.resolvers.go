package graph

import (
	"context"
	"fmt"

	"github.com/Caicrs/Payfood-backend/common"
	"github.com/Caicrs/Payfood-backend/graph/generated"
	"github.com/Caicrs/Payfood-backend/graph/model"
)

func (r *mutationResolver) CreateOrder(ctx context.Context, input *model.NewOrder, products []*model.NewProduct) (*model.Order, error) {

	// context
	context := common.GetContext(ctx)
	order := model.Order{
		MarketplaceUserID: input.MarketplaceUserID,
		ClientID:          input.ClientID,
		TableID:           input.TableID,
		TotalPrice:        input.TotalPrice,
	}

	order.Products = make([]*model.Product, len(products))

	for index, item := range products {
		order.Products[index] = &model.Product{
			MarketplaceID: item.MarketplaceID,
			Name:          item.Name,
			Description:   item.Description,
			Price:         item.Price,
		}
	}

	err := context.Database.Create(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *mutationResolver) CreateClient(ctx context.Context, input model.NewClient) (*model.Client, error) {
	context := common.GetContext(ctx)
	Clients := &model.Client{
		Name:     &input.Name,
		Email:    &input.Email,
		Password: &input.Password,
		Cpf:      &input.Cpf,
		Phone:    &input.Phone,
	}
	err := context.Database.Create(&Clients).Error
	if err != nil {
		return nil, err
	}
	return Clients, nil
}

func (r *mutationResolver) CreateFidelity(ctx context.Context, input model.NewFidelity) (*model.Fidelity, error) {
	context := common.GetContext(ctx)
	Fidelity := &model.Fidelity{
		ClientID: &input.ClientID,
		Score:    &input.Score,
	}
	err := context.Database.Create(&Fidelity).Error
	if err != nil {
		return nil, err
	}
	return Fidelity, nil
}

func (r *mutationResolver) CreateProduct(ctx context.Context, input model.NewProduct) (*model.Product, error) {
	context := common.GetContext(ctx)
	Products := &model.Product{
		MarketplaceID: input.MarketplaceID,
		Name:          input.Name,
		Description:   input.Description,
		Price:         input.Price,
	}
	err := context.Database.Create(&Products).Error
	if err != nil {
		return nil, err
	}
	return Products, nil
}

func (r *mutationResolver) CreateTable(ctx context.Context, input model.NewTable) (*model.Table, error) {
	context := common.GetContext(ctx)
	Tables := &model.Table{
		Number: &input.Number,
		Qrcode: &input.Qrcode,
	}
	err := context.Database.Create(&Tables).Error
	if err != nil {
		return nil, err
	}
	return Tables, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	context := common.GetContext(ctx)
	Users := &model.User{
		Name:        &input.Name,
		Description: &input.Description,
		Cnpj:        &input.Cnpj,
		Email:       &input.Email,
		Password:    &input.Password,
		Image:       &input.Image,
		Superadmin:  &input.Superadmin,
		Admin:       &input.Admin,
	}
	err := context.Database.Create(&Users).Error
	if err != nil {
		return nil, err
	}
	return Users, nil
}

func (r *mutationResolver) DeleteOrder(ctx context.Context, OrderID string) (*model.Order, error) {
	context := common.GetContext(ctx)
	var Order *model.Order
	err := context.Database.Where("id = ?", OrderID).Delete(&Order).Error
	if err != nil {
		return nil, err
	}
	return Order, nil
}

func (r *queryResolver) Clients(ctx context.Context) ([]*model.Client, error) {
	context := common.GetContext(ctx)
	var Clients []*model.Client
	err := context.Database.Find(&Clients).Error
	if err != nil {
		return nil, err
	}
	return Clients, nil
}

func (r *queryResolver) Fidelity(ctx context.Context) ([]*model.Fidelity, error) {
	context := common.GetContext(ctx)
	var Fidelitys []*model.Fidelity
	err := context.Database.Find(&Fidelitys).Error
	if err != nil {
		return nil, err
	}
	return Fidelitys, nil
}

func (r *queryResolver) Orders(ctx context.Context) ([]*model.Order, error) {
	context := common.GetContext(ctx)
	var Orders []*model.Order
	err := context.Database.Find(&Orders).Error
	if err != nil {
		return nil, err
	}
	return Orders, nil
}

func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	context := common.GetContext(ctx)
	var Products []*model.Product
	err := context.Database.Find(&Products).Error
	if err != nil {
		return nil, err
	}
	return Products, nil
}

func (r *queryResolver) Tables(ctx context.Context) ([]*model.Table, error) {
	context := common.GetContext(ctx)
	var Tables []*model.Table
	err := context.Database.Find(&Tables).Error
	if err != nil {
		return nil, err
	}
	return Tables, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	context := common.GetContext(ctx)
	var Users []*model.User
	err := context.Database.Find(&Users).Error
	if err != nil {
		return nil, err
	}
	return Users, nil
}

func (r *queryResolver) FindUser(ctx context.Context, UserID string) ([]*model.User, error) {
	context := common.GetContext(ctx)
	var User []*model.User
	err := context.Database.Where("id = ?", UserID).Find(&User).Error
	if err != nil {
		return nil, err
	}
	return User, nil
}

func (r *queryResolver) FindTables(ctx context.Context, TableID string) ([]*model.Table, error) {
	context := common.GetContext(ctx)
	var Table []*model.Table
	err := context.Database.Where("id = ?", TableID).Find(&Table).Error
	if err != nil {
		return nil, err
	}
	return Table, nil
}

func (r *queryResolver) FindProduct(ctx context.Context, ProductID string) ([]*model.Product, error) {
	context := common.GetContext(ctx)
	var Product []*model.Product
	err := context.Database.Where("id = ?", ProductID).Find(&Product).Error
	if err != nil {
		return nil, err
	}
	return Product, nil
}

func (r *queryResolver) GetOrder(ctx context.Context, OrderID string) (*model.Order, error) {
	context := common.GetContext(ctx)
	var Order *model.Order
	err := context.Database.Where("id = ?", OrderID).Find(&Order).Error
	if err != nil {
		return nil, err
	}
	return Order, nil
}

func (r *queryResolver) FindFidelity(ctx context.Context, FidelityID string) ([]*model.Fidelity, error) {
	context := common.GetContext(ctx)
	var Fidelity []*model.Fidelity
	err := context.Database.Where("id = ?", FidelityID).Find(&Fidelity).Error
	if err != nil {
		return nil, err
	}
	return Fidelity, nil
}

func (r *queryResolver) FindClient(ctx context.Context, ClientID string) ([]*model.Client, error) {
	context := common.GetContext(ctx)
	var Client []*model.Client
	err := context.Database.Where("id = ?", ClientID).Find(&Client).Error
	if err != nil {
		return nil, err
	}
	return Client, nil
}

// FindOrder is the resolver for the findOrder field.
func (r *queryResolver) FindOrder(ctx context.Context, id string) ([]*model.Order, error) {
	panic(fmt.Errorf("not implemented: FindOrder - findOrder"))
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteUser - deleteUser"))
}

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
