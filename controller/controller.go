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
}

type controllers struct {
	service service.Service
}

func NewController(svc service.Service) Controllers {
	return &controllers {
		service: svc,
	}
}

func (control *controllers) GetTodos (c *gin.Context) {
	result, err := control.service.GetTodos(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status" : "Bad Request",
		})
		return
	}
	response := common.BuildResponse(true, "OK", result)

	c.JSON(http.StatusOK, gin.H{
		"data" : response,
	})
}

func (control *controllers) CreateTodos (c *gin.Context) {
	var input entity.TodosInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status" : err.Error(),
		})
		return
	}

	todos := entity.Todos{
		Title: input.Title,
		Description: input.Description,
		IsDone: input.IsDone,
	}
	entity.DB.Create(&todos)
	c.JSON(http.StatusOK, gin.H{
		"data" : todos,
	})
}