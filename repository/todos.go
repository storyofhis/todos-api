package repository

import (
	"context"
	"github.com/storyofhis/golang-crud/todos/entity"
	"gorm.io/gorm"
	"log"
)

type todosRepository struct {
	connection *gorm.DB
}

func NewTodosRepo(db *gorm.DB) TodosRepo {
	return &todosRepository{
		connection: db,
	}
}

func (db *todosRepository) CreateTodo(ctx context.Context, todo entity.Todos) (entity.Todos, error) {
	err := db.connection.WithContext(ctx).Create(&todo).Error
	if err != nil {
		log.Println(err)
		return todo, err
	}
	return todo, nil
}

func (db *todosRepository) GetTodos(ctx context.Context) ([]entity.Todos, error) {
	var todos []entity.Todos
	err := db.connection.WithContext(ctx).Find(&todos).Error
	if err != nil {
		log.Println(err)
		return todos, err
	}
	return todos, nil
}

func (db *todosRepository) GetTodoByID(ctx context.Context, id uint) (entity.Todos, error) {
	var todo entity.Todos
	err := db.connection.WithContext(ctx).Where("id = ?", id).Find(&todo).Error
	if err != nil {
		log.Println(err)
		return todo, err
	}
	return todo, nil
}

func (db *todosRepository) UpdateTodo(ctx context.Context, todo entity.Todos) (entity.Todos, error) {
	err := db.connection.WithContext(ctx).Where("id = ?", todo.ID).Updates(&todo).Error
	if err != nil {
		log.Println(err)
		return todo, err
	}
	return todo, nil
}
