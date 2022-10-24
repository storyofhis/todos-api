package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/storyofhis/golang-crud/todos/common"
	"github.com/storyofhis/golang-crud/todos/entity"
	"github.com/storyofhis/golang-crud/todos/service"
)

type todosController struct {
	svc service.TodoSvc
}

func NewTodoController(svc service.TodoSvc) TodosControllers {
	return &todosController{
		svc: svc,
	}
}

// GetTodos godoc
// @Summary      Show all todos data
// @Description  get todos data
// @Tags         todos
// @Accept       json
// @Produce      json
// @Success      200  {object}  []entity.Todos
// @Failure      400  {object}  http.Header
// @Failure      404  {object}  http.Header
// @Failure      500  {object}  http.Header
// @Router       /todos [get]
func (control *todosController) GetTodos(c *gin.Context) {
	result, err := control.svc.GetTodos(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.BuildErrorResponse(
			"internal server error",
			err))
		return
	}
	c.JSON(http.StatusOK, common.BuildResponse("OK", result))
}

// CreateTodo godoc
// @Summary Create a Todos List
// @Description Create a new Todos With the Input Payload
// @Tags todos
// @Accept json
// @Produce json
// @Success 200 {object} entity.Todos
// @Router /todos [post]
func (control *todosController) CreateTodo(c *gin.Context) {
	var params entity.TodosParams

	if err := c.ShouldBindJSON(&params); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, common.BuildErrorResponse(
			"bad request",
			err))
		return
	}

	result, err := control.svc.CreateTodo(c, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.BuildErrorResponse(
			"internal server error",
			err))
		return
	}

	response := common.BuildResponse("ok", result)
	c.JSON(http.StatusCreated, response)
}

// GetTodoByID godoc
// @Summary      Show an todos by id
// @Description  get string by ID
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Todos ID"
// @Success      200  {object}  	entity.Todos
// @Failure      400  {object}  http.Header
// @Failure      404  {object}  http.Header
// @Failure      500  {object}  http.Header
// @Router       /todos/{id} [get]
func (control *todosController) GetTodoByID(c *gin.Context) {
	paramsId := c.Param("id")
	id64, err := strconv.ParseUint(paramsId, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.BuildErrorResponse(
			"bad request",
			err))
		return
	}
	id := uint(id64)
	result, err := control.svc.GetTodoByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.BuildErrorResponse(
			"internal server error",
			err))
		return
	}
	response := common.BuildResponse("ok", result)
	c.JSON(http.StatusCreated, response)
}

// UpdateTodo	 godoc
// @Summary      Edit todos by id
// @Description  edit uint by ID
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        id   path      uint  true  "Todos ID"
// @Success      200  {object}  entity.Todos
// @Failure      400  {object}  http.Header
// @Failure      404  {object}  http.Header
// @Failure      500  {object}  http.Header
// @Router       /todos/{id} [put]
func (control *todosController) UpdateTodo(c *gin.Context) {
	var params entity.TodosParams

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, common.BuildErrorResponse(
			"invalid json body",
			err))
		return
	}

	paramsId := c.Param("id")
	id64, err := strconv.ParseUint(paramsId, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.BuildErrorResponse(
			"invalid id",
			err))
		return
	}
	id := uint(id64)

	result, err := control.svc.UpdateTodo(c, id, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.BuildErrorResponse(
			"internal server error",
			err))
		return
	}
	response := common.BuildResponse("ok", result)
	c.JSON(http.StatusCreated, response)
}

// DeleteTodo	 godoc
// @Summary      Delete todos by id
// @Description  Delete uint by ID
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        id   path      uint  true  "Todos ID"
// @Success      200  {object}  entity.Todos
// @Failure      400  {object}  http.Header
// @Failure      404  {object}  http.Header
// @Failure      500  {object}  http.Header
// @Router       /todos/{id} [delete]
func (control *todosController) DeleteTodo(c *gin.Context) {
	paramsId := c.Param("id")
	id64, err := strconv.ParseUint(paramsId, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.BuildErrorResponse(
			"invalid id",
			err))
		return
	}
	id := uint(id64)

	result, err := control.svc.DeleteTodo(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.BuildErrorResponse(
			"internal server error",
			err))
		return
	}
	response := common.BuildResponse("ok", result)
	c.JSON(http.StatusCreated, response)
}
