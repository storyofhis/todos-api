package service

import (
	"context"
	"github.com/storyofhis/golang-crud/todos/entity"
)

type TodoSvc interface {
	CreateTodo(context.Context, entity.TodosParams) (*entity.TodosView, error)
	GetTodos(context.Context) ([]entity.TodosView, error)
	GetTodoByID(ctx context.Context, id uint) (*entity.TodosView, error)
	UpdateTodo(ctx context.Context, id uint, params entity.TodosParams) (*entity.TodosView, error)
	DeleteTodo(ctx context.Context, id uint) (*entity.TodosView, error)
}
