package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/storyofhis/golang-crud/todos/controller"
)

func CreateRoute(controller controller.Controllers) *gin.Engine {
	route := gin.Default()

	v1 := route.Group("/v1")
	v1.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "get success",
		})
	})

	v1.POST("/todos", controller.CreateTodos)
	v1.GET("/todos", controller.GetTodos)
	v1.GET("/todos/:id", controller.GetTodoByID)
	v1.PUT("/todos/:id", controller.EditTodos)

	return route
}
