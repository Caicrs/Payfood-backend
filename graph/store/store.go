package store

import (
	"context"
	"net/http"

	"github.com/Caicrs/Payfood-backend/graph/model"
)

type Store struct {
	Clients   []*model.Client
	Orders    []*model.Order
	Fidelitys []*model.Fidelity

	Users    []*model.User
	Products []*model.Product
	Tables   []*model.Table
}

// Client Setup _______________________________________________

func NewClient() *Store {
	clients := make([]*model.Client, 0)
	return &Store{
		Clients: clients,
	}
}

func (s *Store) AddClient(t *model.NewClient) error {
	s.Clients = append(s.Clients, &model.Client{
		ID:       "1",
		Name:     &t.Name,
		Email:    &t.Email,
		Password: &t.Password,
		Cpf:      &t.Cpf,
		Phone:    &t.Phone,
	})
	return nil

}

// Order Setup _______________________________________________

func NewOrder() *Store {
	orders := make([]*model.Order, 0)
	return &Store{Orders: orders}
}

func (s *Store) AddOrder(t *model.NewOrder) error {
	s.Orders = append(s.Orders, &model.Order{
		ID:                "1",
		MarketplaceUserID: &t.MarketplaceUserID,
		ClientID:          &t.ClientID,
		TableID:           &t.TableID,
		TotalPrice:        &t.TotalPrice,
	})
	return nil
}

// Fidelity Setup _______________________________________________

func Fidelity() *Store {
	fidelitys := make([]*model.Fidelity, 0)
	return &Store{Fidelitys: fidelitys}
}

func (s *Store) AddFidelity(t *model.NewFidelity) error {
	s.Fidelitys = append(s.Fidelitys, &model.Fidelity{
		ID:       "1",
		ClientID: &t.ClientID,
		Score:    &t.Score,
	})
	return nil
}

// User Setup _______________________________________________

func NewUser() *Store {
	users := make([]*model.User, 0)
	return &Store{
		Users: users,
	}
}

func (s *Store) AddUser(t *model.NewUser) error {
	s.Users = append(s.Users, &model.User{
		ID:          "1",
		Name:        &t.Name,
		Description: &t.Description,
		Cnpj:        &t.Cnpj,
		Email:       &t.Email,
		Password:    &t.Password,
		Image:       &t.Image,
		Superadmin:  &t.Superadmin,
		Admin:       &t.Admin,
	})
	return nil

}

// Product Setup _______________________________________________
func NewProduct() *Store {
	products := make([]*model.Product, 0)
	return &Store{
		Products: products,
	}
}

func (s *Store) AddProduct(t *model.NewProduct) error {
	s.Products = append(s.Products, &model.Product{
		ID:            "1",
		MarketplaceID: &t.MarketplaceID,
		Name:          &t.Name,
		Description:   &t.Description,
		Price:         &t.Price,
	})
	return nil
}

// Table Setup _______________________________________________
func NewTable() *Store {
	tables := make([]*model.Table, 0)
	return &Store{Tables: tables}
}

func (s *Store) AddTable(t *model.NewTable) error {
	s.Tables = append(s.Tables, &model.Table{
		ID:     "1",
		Number: &t.Number,
		Qrcode: &t.Qrcode,
	})
	return nil
}

// Store Configs _______________________________________________

type StoreKeyType string

var StoreKey StoreKeyType = "STORE"

func WithStore(store *Store, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqWithStore := r.WithContext(context.WithValue(r.Context(), StoreKey, store))
		next.ServeHTTP(w, reqWithStore)
	})
}

func GetStoreFromContext(ctx context.Context) *Store {
	store, ok := ctx.Value(StoreKey).(*Store)
	if !ok {
		panic("couldn't find the store")
	}
	return store
}
