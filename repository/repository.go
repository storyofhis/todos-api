package repository

import (
	"context"

	"github.com/storyofhis/golang-crud/todos/entity"
)

type TodosRepo interface {
	CreateTodo(ctx context.Context, todo entity.Todos) (entity.Todos, error)
	GetTodos(ctx context.Context) ([]entity.Todos, error)
	GetTodoByID(ctx context.Context, id uint) (entity.Todos, error)
	UpdateTodo(ctx context.Context, todo entity.Todos) (entity.Todos, error)
}
