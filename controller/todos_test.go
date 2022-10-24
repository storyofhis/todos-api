package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/storyofhis/golang-crud/todos/common"
	"github.com/storyofhis/golang-crud/todos/entity"
	"github.com/storyofhis/golang-crud/todos/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var (
	tm             = time.Now()
	todoService    service.TodoSvc
	todoController TodosControllers
	createTodo     func(context.Context, entity.TodosParams) (*entity.TodosView, error)
	getTodos       func(context.Context) ([]entity.TodosView, error)
	getTodoByID    func(ctx context.Context, id uint) (*entity.TodosView, error)
	updateTodo     func(ctx context.Context, id uint, params entity.TodosParams) (*entity.TodosView, error)
	deleteTodo     func(ctx context.Context, id uint) (*entity.TodosView, error)
)

type TodoServiceMock struct{}

func (svc *TodoServiceMock) CreateTodo(ctx context.Context, params entity.TodosParams) (*entity.TodosView, error) {
	return createTodo(ctx, params)
}

func (svc *TodoServiceMock) UpdateTodo(ctx context.Context, id uint, params entity.TodosParams) (*entity.TodosView, error) {
	return updateTodo(ctx, id, params)
}

func (svc *TodoServiceMock) GetTodos(ctx context.Context) ([]entity.TodosView, error) {
	return getTodos(ctx)
}

func (svc *TodoServiceMock) GetTodoByID(ctx context.Context, id uint) (*entity.TodosView, error) {
	return getTodoByID(ctx, id)
}

func (svc *TodoServiceMock) DeleteTodo(ctx context.Context, id uint) (*entity.TodosView, error) {
	return deleteTodo(ctx, id)
}

func TestTodosController_CreateTodo_Success(t *testing.T) {
	todoService = &TodoServiceMock{}
	todoController = NewTodoController(todoService)

	requestBody := `
		{
			"title": "My title",
			"title": "My Description",
			"isDone": true
		}
	`

	expectedVal := &entity.TodosView{
		ID:          1,
		Title:       "My title",
		Description: "My Description",
		IsDone:      true,
	}

	createTodo = func(ctx context.Context, params entity.TodosParams) (*entity.TodosView, error) {
		return expectedVal, nil
	}

	req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewBufferString(requestBody))

	router := gin.Default()
	gin.SetMode(gin.TestMode)
	rr := httptest.NewRecorder()
	router.POST("/todos", todoController.CreateTodo)
	router.ServeHTTP(rr, req)
	result := rr.Result()
	body, err := ioutil.ReadAll(result.Body)
	defer result.Body.Close()

	require.Nil(t, err)
	var response common.Response

	err = json.Unmarshal(body, &response)
	assert.Nil(t, err)

	var todo entity.TodosView
	common.MapToStruct(response.Data, &todo)
	assert.NotNil(t, todo)

	assert.EqualValues(t, expectedVal.ID, todo.ID)
	assert.EqualValues(t, expectedVal.Title, todo.Title)
	assert.EqualValues(t, expectedVal.Description, todo.Description)
	assert.EqualValues(t, expectedVal.IsDone, todo.IsDone)
}

func TestTodosController_CreateTodo_ServerError(t *testing.T) {
	todoService = &TodoServiceMock{}
	todoController = NewTodoController(todoService)

	requestBody := `
		{
			"title": "My title",
			"title": "My Description",
			"isDone": true
		}
	`

	serverError := errors.New("internal server error")

	createTodo = func(ctx context.Context, params entity.TodosParams) (*entity.TodosView, error) {
		return nil, serverError
	}

	req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewBufferString(requestBody))

	router := gin.Default()
	gin.SetMode(gin.TestMode)
	rr := httptest.NewRecorder()
	router.POST("/todos", todoController.CreateTodo)
	router.ServeHTTP(rr, req)
	result := rr.Result()
	body, err := ioutil.ReadAll(result.Body)
	defer result.Body.Close()

	require.Nil(t, err)
	var response common.Response

	err = json.Unmarshal(body, &response)
	assert.Nil(t, err)

	assert.EqualValues(t, "internal server error", response.Message)
	assert.EqualValues(t, serverError.Error(), response.Errors)
	assert.EqualValues(t, http.StatusInternalServerError, result.StatusCode)
}

func TestTodosController_UpdateTodo_Success(t *testing.T) {
	todoService = &TodoServiceMock{}
	todoController = NewTodoController(todoService)

	requestBody := `
		{
			"title": "Updated title",
			"title": "Updated Description",
			"isDone": false
		}
	`

	expectedVal := &entity.TodosView{
		ID:          1,
		Title:       "Updated title",
		Description: "Updated Description",
		IsDone:      false,
	}

	updateTodo = func(ctx context.Context, id uint, params entity.TodosParams) (*entity.TodosView, error) {
		return expectedVal, nil
	}

	req := httptest.NewRequest(http.MethodPut, "/todos/1", bytes.NewBufferString(requestBody))

	router := gin.Default()
	gin.SetMode(gin.TestMode)
	rr := httptest.NewRecorder()
	router.PUT("/todos/:id", todoController.UpdateTodo)
	router.ServeHTTP(rr, req)
	result := rr.Result()
	body, err := ioutil.ReadAll(result.Body)
	defer result.Body.Close()

	require.Nil(t, err)
	var response common.Response

	err = json.Unmarshal(body, &response)
	assert.Nil(t, err)

	var todo entity.TodosView
	common.MapToStruct(response.Data, &todo)
	assert.NotNil(t, todo)

	assert.EqualValues(t, expectedVal.ID, todo.ID)
	assert.EqualValues(t, expectedVal.Title, todo.Title)
	assert.EqualValues(t, expectedVal.Description, todo.Description)
	assert.EqualValues(t, expectedVal.IsDone, todo.IsDone)
}

func TestTodosController_UpdateTodo_ServerError(t *testing.T) {
	todoService = &TodoServiceMock{}
	todoController = NewTodoController(todoService)

	requestBody := `
		{
			"title": "Updated title",
			"title": "Updated Description",
			"isDone": false
		}
	`

	serverError := errors.New("internal server error")

	updateTodo = func(ctx context.Context, id uint, params entity.TodosParams) (*entity.TodosView, error) {
		return nil, serverError
	}

	req := httptest.NewRequest(http.MethodPut, "/todos/1", bytes.NewBufferString(requestBody))

	router := gin.Default()
	gin.SetMode(gin.TestMode)
	rr := httptest.NewRecorder()
	router.PUT("/todos/:id", todoController.UpdateTodo)
	router.ServeHTTP(rr, req)
	result := rr.Result()
	body, err := ioutil.ReadAll(result.Body)
	defer result.Body.Close()

	require.Nil(t, err)
	var response common.Response

	err = json.Unmarshal(body, &response)
	assert.Nil(t, err)

	assert.EqualValues(t, "internal server error", response.Message)
	assert.EqualValues(t, serverError.Error(), response.Errors)
	assert.EqualValues(t, http.StatusInternalServerError, result.StatusCode)
}

func TestTodosController_DeleteTodo_Success(t *testing.T) {
	todoService = &TodoServiceMock{}
	todoController = NewTodoController(todoService)

	expectedVal := &entity.TodosView{
		ID:          1,
		Title:       "My title",
		Description: "My Description",
		IsDone:      false,
	}

	deleteTodo = func(ctx context.Context, id uint) (*entity.TodosView, error) {
		return expectedVal, nil
	}

	req := httptest.NewRequest(http.MethodDelete, "/todos/1", nil)

	router := gin.Default()
	gin.SetMode(gin.TestMode)
	rr := httptest.NewRecorder()
	router.DELETE("/todos/:id", todoController.DeleteTodo)
	router.ServeHTTP(rr, req)
	result := rr.Result()
	body, err := ioutil.ReadAll(result.Body)
	defer result.Body.Close()

	require.Nil(t, err)
	var response common.Response

	err = json.Unmarshal(body, &response)
	assert.Nil(t, err)

	var todo entity.TodosView
	common.MapToStruct(response.Data, &todo)
	assert.NotNil(t, todo)

	assert.EqualValues(t, expectedVal.ID, todo.ID)
	assert.EqualValues(t, expectedVal.Title, todo.Title)
	assert.EqualValues(t, expectedVal.Description, todo.Description)
	assert.EqualValues(t, expectedVal.IsDone, todo.IsDone)
}

func TestTodosController_DeleteTodo_ServerError(t *testing.T) {
	todoService = &TodoServiceMock{}
	todoController = NewTodoController(todoService)

	serverError := errors.New("internal server error")

	deleteTodo = func(ctx context.Context, id uint) (*entity.TodosView, error) {
		return nil, serverError
	}

	req := httptest.NewRequest(http.MethodDelete, "/todos/1", nil)

	router := gin.Default()
	gin.SetMode(gin.TestMode)
	rr := httptest.NewRecorder()
	router.DELETE("/todos/:id", todoController.DeleteTodo)
	router.ServeHTTP(rr, req)
	result := rr.Result()
	body, err := ioutil.ReadAll(result.Body)
	defer result.Body.Close()

	require.Nil(t, err)
	var response common.Response

	err = json.Unmarshal(body, &response)
	assert.Nil(t, err)

	assert.EqualValues(t, "internal server error", response.Message)
	assert.EqualValues(t, serverError.Error(), response.Errors)
	assert.EqualValues(t, http.StatusInternalServerError, result.StatusCode)
}

func TestTodosController_GetTodoByID_Success(t *testing.T) {
	todoService = &TodoServiceMock{}
	todoController = NewTodoController(todoService)

	expectedVal := &entity.TodosView{
		ID:          1,
		Title:       "My title",
		Description: "My Description",
		IsDone:      false,
	}

	getTodoByID = func(ctx context.Context, id uint) (*entity.TodosView, error) {
		return expectedVal, nil
	}

	req := httptest.NewRequest(http.MethodGet, "/todos/1", nil)

	router := gin.Default()
	gin.SetMode(gin.TestMode)
	rr := httptest.NewRecorder()
	router.GET("/todos/:id", todoController.GetTodoByID)
	router.ServeHTTP(rr, req)
	result := rr.Result()
	body, err := ioutil.ReadAll(result.Body)
	defer result.Body.Close()

	require.Nil(t, err)
	var response common.Response

	err = json.Unmarshal(body, &response)
	assert.Nil(t, err)

	var todo entity.TodosView
	common.MapToStruct(response.Data, &todo)
	assert.NotNil(t, todo)

	assert.EqualValues(t, expectedVal.ID, todo.ID)
	assert.EqualValues(t, expectedVal.Title, todo.Title)
	assert.EqualValues(t, expectedVal.Description, todo.Description)
	assert.EqualValues(t, expectedVal.IsDone, todo.IsDone)
}

func TestTodosController_GetTodoByID_ServerError(t *testing.T) {
	todoService = &TodoServiceMock{}
	todoController = NewTodoController(todoService)

	serverError := errors.New("internal server error")

	getTodoByID = func(ctx context.Context, id uint) (*entity.TodosView, error) {
		return nil, serverError
	}

	req := httptest.NewRequest(http.MethodGet, "/todos/1", nil)

	router := gin.Default()
	gin.SetMode(gin.TestMode)
	rr := httptest.NewRecorder()
	router.GET("/todos/:id", todoController.GetTodoByID)
	router.ServeHTTP(rr, req)
	result := rr.Result()
	body, err := ioutil.ReadAll(result.Body)
	defer result.Body.Close()

	require.Nil(t, err)
	var response common.Response

	err = json.Unmarshal(body, &response)
	assert.Nil(t, err)

	assert.EqualValues(t, "internal server error", response.Message)
	assert.EqualValues(t, serverError.Error(), response.Errors)
	assert.EqualValues(t, http.StatusInternalServerError, result.StatusCode)
}

func TestTodosController_GetTodos_Success(t *testing.T) {
	todoService = &TodoServiceMock{}
	todoController = NewTodoController(todoService)

	expectedValues := []entity.TodosView{
		{
			ID:          1,
			Title:       "My title",
			Description: "My Description",
			IsDone:      false,
		},
		{
			ID:          2,
			Title:       "My title 2",
			Description: "My Description 2",
			IsDone:      true,
		},
	}

	getTodos = func(ctx context.Context) ([]entity.TodosView, error) {
		return expectedValues, nil
	}

	req := httptest.NewRequest(http.MethodGet, "/todos", nil)

	router := gin.Default()
	gin.SetMode(gin.TestMode)
	rr := httptest.NewRecorder()
	router.GET("/todos", todoController.GetTodos)
	router.ServeHTTP(rr, req)
	result := rr.Result()
	body, err := ioutil.ReadAll(result.Body)
	defer result.Body.Close()

	require.Nil(t, err)
	var response common.Response

	err = json.Unmarshal(body, &response)
	assert.Nil(t, err)

	var todos []entity.TodosView
	common.MapToStruct(response.Data, &todos)
	assert.NotNil(t, todos)

	for i, _ := range expectedValues {
		assert.EqualValues(t, expectedValues[i].ID, todos[i].ID)
		assert.EqualValues(t, expectedValues[i].Title, todos[i].Title)
		assert.EqualValues(t, expectedValues[i].Description, todos[i].Description)
		assert.EqualValues(t, expectedValues[i].IsDone, todos[i].IsDone)
	}
}

func TestTodosController_GetTodos_ServerError(t *testing.T) {
	todoService = &TodoServiceMock{}
	todoController = NewTodoController(todoService)

	serverError := errors.New("internal server error")

	getTodos = func(ctx context.Context) ([]entity.TodosView, error) {
		return []entity.TodosView{}, serverError
	}

	req := httptest.NewRequest(http.MethodGet, "/todos", nil)

	router := gin.Default()
	gin.SetMode(gin.TestMode)
	rr := httptest.NewRecorder()
	router.GET("/todos", todoController.GetTodos)
	router.ServeHTTP(rr, req)
	result := rr.Result()
	body, err := ioutil.ReadAll(result.Body)
	defer result.Body.Close()

	require.Nil(t, err)
	var response common.Response

	err = json.Unmarshal(body, &response)
	assert.Nil(t, err)

	assert.EqualValues(t, "internal server error", response.Message)
	assert.EqualValues(t, serverError.Error(), response.Errors)
	assert.EqualValues(t, http.StatusInternalServerError, result.StatusCode)
}
