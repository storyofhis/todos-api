package router

import (
	"github.com/gin-gonic/gin"
	"github.com/storyofhis/golang-crud/todos/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SwaggerRoute(router *gin.Engine) {
	docs.SwaggerInfo.BasePath = ""

	swaggerRoutes := router.Group("/swagger")
	{
		swaggerRoutes.GET("/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
}
