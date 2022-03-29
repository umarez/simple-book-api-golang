package service

import (
	"library-api/model"
	"library-api/repository"
)

type Service interface {
	FindAll() ([]model.Book, error)
	FindById(ID int) (model.Book, error)
	Create(book model.Book) (model.Book, error)
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]model.Book, error) {
	return s.repository.FindAll()
}

func (s *service) FindById(ID int) (model.Book, error) {
	book, err := s.repository.FindById(ID)
	return book, err
}

func (s *service) Create(book model.Book) (model.Book, error) {

	// books := model.Book{
	// 	BookResponse: model.BookResponse{
	// 		Title:       book.Title,
	// 		Description: book.Description,
	// 		Price:       book.Price,
	// 		Rating:      book.Rating,
	// 	},
	// }

	books, err := s.repository.Create(&book)
	return books, err
}
