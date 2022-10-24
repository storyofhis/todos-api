package service

import (
	"context"
	"log"

	"github.com/storyofhis/golang-crud/todos/entity"
	"github.com/storyofhis/golang-crud/todos/repository"
)

type todoService struct {
	repo repository.TodosRepo
}

func NewService(repo repository.TodosRepo) TodoSvc {
	return &todoService{
		repo: repo,
	}
}

func (svc *todoService) GetTodos(ctx context.Context) ([]entity.TodosView, error) {
	res, err := svc.repo.GetTodos(ctx)
	if err != nil {
		log.Println(err)
	}

	todoViews := make([]entity.TodosView, 0)
	for _, v := range res {
		todoViews = append(todoViews, entity.TodosView{
			ID:          v.ID,
			Title:       v.Title,
			Description: v.Description,
			IsDone:      v.IsDone,
		})
	}
	return todoViews, nil
}

func (svc *todoService) CreateTodo(ctx context.Context, params entity.TodosParams) (*entity.TodosView, error) {
	todo := entity.Todos{
		Title:       params.Title,
		Description: params.Description,
		IsDone:      params.IsDone,
	}
	res, err := svc.repo.CreateTodo(ctx, todo)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &entity.TodosView{
		ID:          res.ID,
		Title:       res.Title,
		Description: res.Description,
		IsDone:      res.IsDone,
	}, nil
}

func (svc *todoService) GetTodoByID(ctx context.Context, id uint) (*entity.TodosView, error) {
	res, err := svc.repo.GetTodoByID(ctx, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &entity.TodosView{
		ID:          res.ID,
		Title:       res.Title,
		Description: res.Description,
		IsDone:      res.IsDone,
	}, nil
}

func (svc *todoService) UpdateTodo(ctx context.Context, id uint, params entity.TodosParams) (*entity.TodosView, error) {
	todo := entity.Todos{
		Title:       params.Title,
		Description: params.Description,
		IsDone:      params.IsDone,
	}
	todo.ID = id
	res, err := svc.repo.UpdateTodo(ctx, todo)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &entity.TodosView{
		ID:          res.ID,
		Title:       res.Title,
		Description: res.Description,
		IsDone:      res.IsDone,
	}, nil
}

func (svc *todoService) DeleteTodo(ctx context.Context, id uint) (*entity.TodosView, error) {
	res, err := svc.repo.DeleteTodo(ctx, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &entity.TodosView{
		ID:          res.ID,
		Title:       res.Title,
		Description: res.Description,
		IsDone:      res.IsDone,
	}, nil
}
