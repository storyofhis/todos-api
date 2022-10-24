package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/storyofhis/golang-crud/todos/config"
	"github.com/storyofhis/golang-crud/todos/controller"
	"github.com/storyofhis/golang-crud/todos/docs" // docs is generated by Swag CLI, you have to import it.
	"github.com/storyofhis/golang-crud/todos/entity"
	"github.com/storyofhis/golang-crud/todos/repository"
	"github.com/storyofhis/golang-crud/todos/router"
	"github.com/storyofhis/golang-crud/todos/service"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
		return
	}
}

// @Title Todos API
// @Description This is a simple API for managing Create, Read, Update and Delete Data
// @Version 1.0
// @Host localhost:8080
// @BasePath /v1/
// @Schemes http
// @contact.name API Support
// @contact.email azizi.maula@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "ToDos API"
	docs.SwaggerInfo.Description = "This is a simple API for managing Create, Read, Update and Delete Data"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/v1/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	var (
		db              = config.ConnectDB()
		PORT            = os.Getenv("PORT")
		todosRepo       = repository.NewTodosRepo(db)
		todosSvc        = service.NewTodoService(todosRepo)
		todosController = controller.NewTodoController(todosSvc)
	)
	entity.DB.AutoMigrate(&entity.Todos{})
	app := router.CreateRoute(todosController)

	app.Run(":" + PORT)
}
