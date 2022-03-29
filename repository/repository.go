package repository

import (
	"fmt"
	"library-api/model"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]model.Book, error)
	FindById(ID int) (model.Book, error)
	Create(book *model.Book) (model.Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]model.Book, error) {
	var books []model.Book
	err := r.db.Debug().Find(&books).Error

	return books, err
}

func (r *repository) FindById(ID int) (model.Book, error) {
	var books model.Book
	err := r.db.Debug().First(&books, ID).Error

	return books, err
}

func (r *repository) Create(book *model.Book) (model.Book, error) {
	err := r.db.Debug().Create(&book).Error
	fmt.Println(book)
	return *book, err
}
