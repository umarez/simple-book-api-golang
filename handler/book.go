package handler

import (
	"fmt"
	"library-api/model"
	"library-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type bookHandler struct {
	bookService service.Service
}

func NewBookHandler(bookService service.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (handler *bookHandler) HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "umar Izzuddin",
		"bio":  "I'm software Engineer!",
	})
}

func (handler *bookHandler) BookHandler(c *gin.Context) {
	var book []model.Book

	book, err := handler.bookService.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, book)

}

func (handler *bookHandler) QueryHandler(c *gin.Context) {
	var book model.Book

	id := c.Param("id")
	idNum, _ := strconv.Atoi(id)

	book, err := handler.bookService.FindById(idNum)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, book)
}

func (handler *bookHandler) CreateBookHandler(c *gin.Context) {
	var bookInput model.Book
	validate = validator.New()

	c.ShouldBindJSON(&bookInput)
	errs := validate.Struct(bookInput)

	if errs != nil {

		for _, e := range errs.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintln(e.Error())

			c.JSON(http.StatusBadRequest, errorMessage)
			return
		}

	}

	book, err := handler.bookService.Create(bookInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	res := &model.BookResponse{
		Title:       book.Title,
		Description: book.Description,
		Price:       book.Price,
		Rating:      book.Price,
	}

	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})

}
