package model

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name      string    `json:"name_book"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
