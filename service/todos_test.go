package service

import (
	"context"
	"errors"
	"github.com/storyofhis/golang-crud/todos/entity"
	"github.com/storyofhis/golang-crud/todos/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	todoSvc     TodoSvc
	todoRepo    repository.TodosRepo
	createTodo  func(ctx context.Context, todo entity.Todos) (entity.Todos, error)
	getTodos    func(ctx context.Context) ([]entity.Todos, error)
	getTodoByID func(ctx context.Context, id uint) (entity.Todos, error)
	updateTodo  func(ctx context.Context, todo entity.Todos) (entity.Todos, error)
	deleteTodo  func(ctx context.Context, id uint) (entity.Todos, error)
)

type todosRepositoryMock struct{}

func (r *todosRepositoryMock) CreateTodo(ctx context.Context, todo entity.Todos) (entity.Todos, error) {
	return createTodo(ctx, todo)
}

func (r *todosRepositoryMock) UpdateTodo(ctx context.Context, todo entity.Todos) (entity.Todos, error) {
	return updateTodo(ctx, todo)
}

func (r *todosRepositoryMock) GetTodoByID(ctx context.Context, id uint) (entity.Todos, error) {
	return getTodoByID(ctx, id)
}

func (r *todosRepositoryMock) GetTodos(ctx context.Context) ([]entity.Todos, error) {
	return getTodos(ctx)
}

func (r *todosRepositoryMock) DeleteTodo(ctx context.Context, id uint) (entity.Todos, error) {
	return deleteTodo(ctx, id)
}

func TestTodoService_CreateTodo_Success(t *testing.T) {
	todoRepo = &todosRepositoryMock{}
	todoSvc = NewTodoService(todoRepo)
	ctx := context.TODO()

	params := entity.TodosParams{
		Title:       "My Title",
		Description: "My Description",
		IsDone:      true,
	}

	expectedVal := entity.Todos{
		Title:       "My Title",
		Description: "My Description",
		IsDone:      true,
	}
	expectedVal.ID = 1

	createTodo = func(ctx context.Context, todo entity.Todos) (entity.Todos, error) {
		return expectedVal, nil
	}

	todo, err := todoSvc.CreateTodo(ctx, params)
	assert.Nil(t, err)

	assert.NotNil(t, todo)
	assert.EqualValues(t, expectedVal.ID, todo.ID)
	assert.EqualValues(t, expectedVal.Title, todo.Title)
	assert.EqualValues(t, expectedVal.Description, todo.Description)
	assert.EqualValues(t, expectedVal.IsDone, todo.IsDone)
}

func TestTodosService_CreateTodo_ServerError(t *testing.T) {
	todoRepo = &todosRepositoryMock{}
	todoSvc = NewTodoService(todoRepo)
	ctx := context.TODO()

	params := entity.TodosParams{
		Title:       "My Title",
		Description: "My Description",
		IsDone:      true,
	}
	serverError := errors.New("internal server error")

	createTodo = func(ctx context.Context, todo entity.Todos) (entity.Todos, error) {
		return entity.Todos{}, serverError
	}

	_, err := todoSvc.CreateTodo(ctx, params)
	assert.NotNil(t, err)
	assert.EqualValues(t, serverError, err)
}

func TestTodoService_UpdateTodo_Success(t *testing.T) {
	todoRepo = &todosRepositoryMock{}
	todoSvc = NewTodoService(todoRepo)
	ctx := context.TODO()

	params := entity.TodosParams{
		Title:       "Updated Title",
		Description: "Updated Description",
		IsDone:      false,
	}

	expectedVal := entity.Todos{
		Title:       "Updated Title",
		Description: "Updated Description",
		IsDone:      false,
	}
	expectedVal.ID = 1

	updateTodo = func(ctx context.Context, todo entity.Todos) (entity.Todos, error) {
		return expectedVal, nil
	}

	todo, err := todoSvc.UpdateTodo(ctx, expectedVal.ID, params)
	assert.Nil(t, err)

	assert.NotNil(t, todo)
	assert.EqualValues(t, expectedVal.ID, todo.ID)
	assert.EqualValues(t, expectedVal.Title, todo.Title)
	assert.EqualValues(t, expectedVal.Description, todo.Description)
	assert.EqualValues(t, expectedVal.IsDone, todo.IsDone)
}

func TestTodosService_UpdateTodo_ServerError(t *testing.T) {
	todoRepo = &todosRepositoryMock{}
	todoSvc = NewTodoService(todoRepo)
	ctx := context.TODO()

	params := entity.TodosParams{
		Title:       "Updated Title",
		Description: "Updated Description",
		IsDone:      false,
	}
	serverError := errors.New("internal server error")

	updateTodo = func(ctx context.Context, todo entity.Todos) (entity.Todos, error) {
		return entity.Todos{}, serverError
	}

	_, err := todoSvc.UpdateTodo(ctx, 1, params)
	assert.NotNil(t, err)
	assert.EqualValues(t, serverError, err)
}

func TestTodoService_GetTodos_Success(t *testing.T) {
	todoRepo = &todosRepositoryMock{}
	todoSvc = NewTodoService(todoRepo)
	ctx := context.TODO()

	expectedValues := []entity.Todos{
		{
			Title:       "Updated Title",
			Description: "Updated Description",
			IsDone:      false,
		},
		{
			Title:       "Updated Title 2",
			Description: "Updated Description 2",
			IsDone:      true,
		},
	}
	expectedValues[0].ID = 1
	expectedValues[1].ID = 1

	getTodos = func(ctx context.Context) ([]entity.Todos, error) {
		return expectedValues, nil
	}

	todos, err := todoSvc.GetTodos(ctx)
	assert.Nil(t, err)

	assert.NotNil(t, todos)
	for i, _ := range expectedValues {
		assert.EqualValues(t, expectedValues[i].ID, todos[i].ID)
		assert.EqualValues(t, expectedValues[i].Title, todos[i].Title)
		assert.EqualValues(t, expectedValues[i].Description, todos[i].Description)
		assert.EqualValues(t, expectedValues[i].IsDone, todos[i].IsDone)
	}
}

func TestTodosService_GetTodos_ServerError(t *testing.T) {
	todoRepo = &todosRepositoryMock{}
	todoSvc = NewTodoService(todoRepo)
	ctx := context.TODO()

	serverError := errors.New("internal server error")

	getTodos = func(ctx context.Context) ([]entity.Todos, error) {
		return nil, serverError
	}

	_, err := todoSvc.GetTodos(ctx)
	assert.NotNil(t, err)
	assert.EqualValues(t, serverError, err)
}

func TestTodoService_GetTodoByID_Success(t *testing.T) {
	todoRepo = &todosRepositoryMock{}
	todoSvc = NewTodoService(todoRepo)
	ctx := context.TODO()

	expectedVal := entity.Todos{
		Title:       "My Title",
		Description: "My Description",
		IsDone:      true,
	}
	expectedVal.ID = 1

	getTodoByID = func(ctx context.Context, id uint) (entity.Todos, error) {
		return expectedVal, nil
	}

	todo, err := todoSvc.GetTodoByID(ctx, 1)
	assert.Nil(t, err)

	assert.NotNil(t, todo)
	assert.EqualValues(t, expectedVal.ID, todo.ID)
	assert.EqualValues(t, expectedVal.Title, todo.Title)
	assert.EqualValues(t, expectedVal.Description, todo.Description)
	assert.EqualValues(t, expectedVal.IsDone, todo.IsDone)
}

func TestTodosService_GetTodoByID_ServerError(t *testing.T) {
	todoRepo = &todosRepositoryMock{}
	todoSvc = NewTodoService(todoRepo)
	ctx := context.TODO()

	serverError := errors.New("internal server error")

	getTodoByID = func(ctx context.Context, id uint) (entity.Todos, error) {
		return entity.Todos{}, serverError
	}

	_, err := todoSvc.GetTodoByID(ctx, 1)
	assert.NotNil(t, err)
	assert.EqualValues(t, serverError, err)
}

func TestTodoService_DeleteTodo_Success(t *testing.T) {
	todoRepo = &todosRepositoryMock{}
	todoSvc = NewTodoService(todoRepo)
	ctx := context.TODO()

	expectedVal := entity.Todos{
		Title:       "My Title",
		Description: "My Description",
		IsDone:      true,
	}
	expectedVal.ID = 1

	deleteTodo = func(ctx context.Context, id uint) (entity.Todos, error) {
		return expectedVal, nil
	}

	todo, err := todoSvc.DeleteTodo(ctx, 1)
	assert.Nil(t, err)

	assert.NotNil(t, todo)
	assert.EqualValues(t, expectedVal.ID, todo.ID)
	assert.EqualValues(t, expectedVal.Title, todo.Title)
	assert.EqualValues(t, expectedVal.Description, todo.Description)
	assert.EqualValues(t, expectedVal.IsDone, todo.IsDone)
}

func TestTodosService_DeleteTodo_ServerError(t *testing.T) {
	todoRepo = &todosRepositoryMock{}
	todoSvc = NewTodoService(todoRepo)
	ctx := context.TODO()

	serverError := errors.New("internal server error")

	deleteTodo = func(ctx context.Context, id uint) (entity.Todos, error) {
		return entity.Todos{}, serverError
	}

	_, err := todoSvc.DeleteTodo(ctx, 1)
	assert.NotNil(t, err)
	assert.EqualValues(t, serverError, err)
}
