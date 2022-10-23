package router

import (
	"github.com/gin-gonic/gin"
	"github.com/storyofhis/golang-crud/todos/controller"
)

func CreateRoute(controller controller.Controllers) *gin.Engine {
	route := gin.Default()

	v1 := route.Group("/v1")
	v1.Use(cors)

	v1.POST("/todos", controller.CreateTodos)
	v1.GET("/todos", controller.GetTodos)
	v1.GET("/todos/:id", controller.GetTodoByID)
	v1.PUT("/todos/:id", controller.EditTodos)

	return route
}

func cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
}
