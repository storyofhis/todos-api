package router

import (
	"github.com/gin-gonic/gin"
	"github.com/storyofhis/golang-crud/todos/controller"
)

func CreateRoute(controller controller.TodosControllers) *gin.Engine {
	route := gin.Default()

	v1 := route.Group("/v1")
	v1.Use(cors)

	v1.POST("/todos", controller.CreateTodo)
	v1.GET("/todos", controller.GetTodos)
	v1.GET("/todos/:id", controller.GetTodoByID)
	v1.PUT("/todos/:id", controller.UpdateTodo)
	v1.DELETE("/todos/:id", controller.DeleteTodo)

	return route
}

func cors(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT, DELETE")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
}
