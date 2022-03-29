package model

import (
	"gorm.io/gorm"
)

type Book struct {
	BookResponse
	gorm.Model
}

type BookResponse struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	Rating      int    `json:"rating" validate:"required"`
}
