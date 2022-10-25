package store
import (
	"github.com/Caicrs/Payfood-backend/graph/model"
	"net/http"
	"context"
)


type Store struct {
	Todos []*model.Todo
}

func NewStore() *Store{
	todos := make([]*model.Todo,0)
	return &Store{
	Todos:todos,
	}
}

func(s *Store) AddTodo(t *model.NewTodo) error{
	s.Todos = append(s.Todos, &model.Todo{
		ID: "1",
		Text: t.Text,
		Done:false,
		User: &model.User{
		ID: t.UserID,
		Name: "caicdev",
		},
	})
return nil
} 

type StoreKeyType string

var StoreKey StoreKeyType = "STORE"


func WithStore(store *Store,next http.Handler) http.Handler{
return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
	reqWithStore := r.WithContext(context.WithValue(r.Context(), StoreKey, store))
	next.ServeHTTP(w, reqWithStore)
})
}

func GetStoreFromContext(ctx context.Context) *Store{
	store, ok := ctx.Value(StoreKey).(*Store)
	if !ok{
	panic("couldn't find the store")
	}
	return store
	}