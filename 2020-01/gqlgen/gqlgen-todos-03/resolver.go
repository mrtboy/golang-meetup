package gqlgen_todos

import (
	"context"

	"github.com/friendsofgo/errors"
)

type Resolver struct {
	todosRepo *TodosRepository
}

var _ ResolverRoot = &Resolver{}

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

func (r *mutationResolver) MarkDone(ctx context.Context, todoID string) (*Result, error) {
	todo := r.todosRepo.FindTodoByID(todoID)
	if todo == nil {
		return &Result{Error: strVal("todo not found")}, nil
	}
	todo.Done = true
	err := r.todosRepo.UpdateTodo(todo)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update todo")
	}
	return &Result{}, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]*Todo, error) {
	return r.todosRepo.FindAllTodos(), nil
}

func (r *Resolver) Todo() TodoResolver {
	return &todoResolver{
		todosRepo: r.todosRepo,
	}
}

type todoResolver struct {
	todosRepo *TodosRepository
}

func (t *todoResolver) User(ctx context.Context, todo *Todo) (*User, error) {
	return t.todosRepo.FindUserByID(todo.UserID), nil
}

func strVal(s string) *string {
	return &s
}
