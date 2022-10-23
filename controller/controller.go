package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/storyofhis/golang-crud/todos/common"
	"github.com/storyofhis/golang-crud/todos/entity"
	"github.com/storyofhis/golang-crud/todos/service"
)

type Controllers interface {
	CreateTodos(c *gin.Context)
	GetTodos(c *gin.Context)
	GetTodoByID(c *gin.Context)
	EditTodos(c *gin.Context)
	DeleteTodo(c *gin.Context)
}

type controllers struct {
	service service.Service
}

func NewController(svc service.Service) Controllers {
	return &controllers{
		service: svc,
	}
}

// GetTodos godoc
// @Summary      Show all todos data
// @Description  get todos data
// @Tags         todos
// @Accept       json
// @Produce      json
// @Success      200  {object}  entity.Todos
// @Failure      400  {object}  http.Header
// @Failure      404  {object}  http.Header
// @Failure      500  {object}  http.Header
// @Router       /v1/todos [get]
func (control *controllers) GetTodos(c *gin.Context) {
	// try to solve CORS error conditions
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	result, err := control.service.GetTodos(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Bad Request",
		})
		return
	}
	response := common.BuildResponse(true, "OK", result)

	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

// CreateTodos godoc
// @Summary Create a Todos List
// @Description Create a new Todos With the Input Payload
// @Tags todos
// @Accept json
// @Produce json
// @Success 200 {object} entity.Todos
// @Router /v1/todos [post]
func (control *controllers) CreateTodos(c *gin.Context) {
	// try to solve CORS error conditions
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	var input entity.TodosInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	todos := entity.Todos{
		Title:       input.Title,
		Description: input.Description,
		IsDone:      input.IsDone,
	}
	entity.DB.Create(&todos)
	c.JSON(http.StatusOK, gin.H{
		"data": todos,
	})
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
// @Router       /v1/todos/{id} [get]
func (control *controllers) GetTodoByID(c *gin.Context) {
	// try to solve CORS error conditions
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	id := c.Param("id")
	// strID, err := strconv.ParseUint(id)
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "id cannot empty",
		})
	}
	result, err := control.service.GetTodoByID(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}
	// id := c.Param("id")
	response := common.BuildResponse(true, "OK", result)
	// err = entity.DB.Where("id = ?", id).Find(&response).Error
	// fmt.Println(id)
	// if err != nil {
	// 	log.Println("id not found")
	// }
	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

func (control *controllers) EditTodos(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	var input entity.TodosInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	todos := entity.Todos{
		Title:       input.Title,
		Description: input.Description,
		IsDone:      input.IsDone,
	}
	id := c.Param("id")
	result, err := control.service.EditTodos(c, id)
	entity.DB.Model(&todos).Updates(&todos)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}
	response := common.BuildResponse(true, "OK", result)
	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

// DeleteTodo godoc
// @Summary      Delete an todos by id
// @Description  Delete todos
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Todos ID"
// @Success      200  {object}  	entity.Todos
// @Failure      400  {object}  http.Header
// @Failure      404  {object}  http.Header
// @Failure      500  {object}  http.Header
// @Router       /v1/todos/{id} [delete]
func (control *controllers) DeleteTodo(c *gin.Context) {
	var todos entity.Todos
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "id cannot empty",
		})
	}
	result, err := control.service.DeleteTodo(c, id)
	entity.DB.Delete(&todos)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}
	response := common.BuildResponse(true, "OK", result)
	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})

}
