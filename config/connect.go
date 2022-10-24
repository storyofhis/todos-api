package config

import (
	"fmt"
	"log"
	"os"

	"github.com/storyofhis/golang-crud/todos/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dbUser := os.Getenv("PGUSER")
	dbPass := os.Getenv("PGPASSWORD")
	dbHost := os.Getenv("PGHOST")
	dbName := os.Getenv("PGDATABASE")
	dbPort := os.Getenv("PGPORT")

	var err error
	dsn := fmt.Sprintf(
		"host=%s port=%s password=%s user=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbPass, dbUser, dbName,
	)

	entity.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("Cannot Open DB")
		return nil
	}
	return entity.DB
}
