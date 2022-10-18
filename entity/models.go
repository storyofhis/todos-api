package entity

import "gorm.io/gorm"

var DB *gorm.DB

type Todos struct {
	gorm.Model
	Title string `json:"title"`
	Description string `json:"description"`
	IsDone bool `json:"isDone"`
}

type TodosInput struct {
	Title string `json:"title"`
	Description string `json:"description"`
	IsDone bool `json:"isDone"`
}