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
}

type controllers struct {
	service service.Service
}

func NewController(svc service.Service) Controllers {
	return &controllers{
		service: svc,
	}
}

func (control *controllers) GetTodos(c *gin.Context) {
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

func (control *controllers) CreateTodos(c *gin.Context) {
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

func (control *controllers) GetTodoByID(c *gin.Context) {
	id := c.Param("id")
	// strID, err := strconv.ParseUint(id)
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
	entity.DB.Model(&todos).Updates(input)
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
