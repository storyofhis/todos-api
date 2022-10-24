package controller

import (
	"github.com/gin-gonic/gin"
)

type TodosControllers interface {
	CreateTodo(c *gin.Context)
	GetTodos(c *gin.Context)
	GetTodoByID(c *gin.Context)
	UpdateTodo(c *gin.Context)
}
