package service

import (
	"context"
	"log"

	"github.com/storyofhis/golang-crud/todos/entity"
	"github.com/storyofhis/golang-crud/todos/repository"
)

type Service interface {
	CreateTodos(context.Context) (entity.Todos, error)
	GetTodos(context.Context) ([]entity.Todos, error)
	GetTodoByID(ctx context.Context, id string) (entity.Todos, error)
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{
		repo: repo,
	}
}

func (svc *service) GetTodos(ctx context.Context) ([]entity.Todos, error) {
	res, err := svc.repo.GetTodos(ctx)
	if err != nil {
		log.Println(err)
	}
	return res, nil
}

func (svc *service) CreateTodos(ctx context.Context) (entity.Todos, error) {
	res, err := svc.repo.CreateTodos(ctx)
	if err != nil {
		log.Println(err)
	}
	return res, nil
}

func (svc *service) GetTodoByID(ctx context.Context, id string) (entity.Todos, error) {
	res, err := svc.repo.GetTodoByID(ctx, id)
	if err != nil {
		log.Println(err)
		return res, err
	}
	return res, nil
}
