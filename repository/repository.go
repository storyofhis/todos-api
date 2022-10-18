package repository

import (
	"context"
	"log"

	"github.com/storyofhis/golang-crud/todos/entity"
	"gorm.io/gorm"
)

type Repository interface {
	CreateTodos(context.Context) (entity.Todos, error)
	GetTodos(context.Context) ([]entity.Todos, error)
}

type repositories struct {
	connection *gorm.DB
}

func NewRepository (db *gorm.DB) Repository {
	return &repositories{
		connection: db,
	}
}

func (db *repositories) CreateTodos (ctx context.Context) (entity.Todos, error) {
	var todos entity.Todos

	err := db.connection.WithContext(ctx).Create(&todos).Error
	if err != nil {
		log.Println(err)
		return todos, err
	}
	return todos, nil
}

func (db *repositories) GetTodos (ctx context.Context) ([]entity.Todos, error) {
	var todos []entity.Todos
	err := db.connection.WithContext(ctx).Find(&todos).Error
	if err != nil {
		log.Println(err)
		return todos, err
	}
	return todos, nil
}