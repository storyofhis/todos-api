package entity

import (
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

type BaseModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

// Todos Represents the model for an todos
type Todos struct {
	BaseModel
	Title       string `json:"title"`
	Description string `json:"description"`
	IsDone      bool   `json:"isDone"`
}

type TodosParams struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	IsDone      bool   `json:"isDone"`
}

type TodosView struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsDone      bool   `json:"isDone"`
}
