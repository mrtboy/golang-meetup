package gqlgen_todos

import (
	"context"

	"github.com/friendsofgo/errors"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	todosRepo *TodosRepository
}

func NewResolver(todosRepo *TodosRepository) *Resolver {
	return &Resolver{
		todosRepo: todosRepo,
	}
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (*Todo, error) {
	todo, err := r.todosRepo.InsertTodo(input)
	if err != nil {
		return nil, errors.Wrap(err, "failed to insert todo")
	}

	return todo, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]*Todo, error) {
	return r.todosRepo.FindAllTodos(), nil
}
