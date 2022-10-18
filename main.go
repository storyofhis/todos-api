package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/storyofhis/golang-crud/todos/config"
	"github.com/storyofhis/golang-crud/todos/controller"
	"github.com/storyofhis/golang-crud/todos/repository"
	"github.com/storyofhis/golang-crud/todos/router"
	"github.com/storyofhis/golang-crud/todos/service"
	"gorm.io/gorm"
)


func init () {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
		return
	}
}
func main() {
	var (
		db *gorm.DB                    = config.ConnectDB()
		repo repository.Repository = repository.NewRepository(db)
		svc service.Service = service.NewService(repo)


		controller controller.Controllers = controller.NewController(svc)
	)
	app := router.CreateRoute(controller)
	app.Run(":8080")
}