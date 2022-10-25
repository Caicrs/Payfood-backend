package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/Caicrs/Payfood-backend/graph/store"
	"github.com/Caicrs/Payfood-backend/graph/generated"
	"github.com/Caicrs/Payfood-backend/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	db := store.GetStoreFromContext(ctx)
	err := db.AddTodo(&input)
	if err != nil {
		return nil, err
	}

	return &model.Todo{
		ID: "1",
		Text: input.Text,
		Done:false,
		User: &model.User{
		ID: input.UserID,
		Name: "caicdev",
		},
		}, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	db := store.GetStoreFromContext(ctx)
    return db.Todos,nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
