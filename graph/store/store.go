package store

import (
	"context"
	"net/http"

	"github.com/Caicrs/Payfood-backend/graph/model"
)

type Store struct {
	Clients []*model.Client
}

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
