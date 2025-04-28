package transport

import (
	"learn/rest-api/internal/book/service"

	"github.com/gin-gonic/gin"
)

type Service interface {
	GetBook(service.BookInputDTO) (service.BookOutputDTO, error)
	PostBook(service.CreateBookInputDTO) (service.CreateBookOutputDTO, error)
}

type BookHandler struct {
	service Service
}

func NewBookHandler(s Service) *BookHandler {
	return &BookHandler{service: s}
}

func (h *BookHandler) GetBook(c *gin.Context) {
	var query BookRequestDTO
	if err := c.ShouldBindUri(&query); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	input := service.BookInputDTO{ID: query.ID}
	book, err := h.service.GetBook(input)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	responseBook := BookResponseDTO{ID: book.ID, Name: book.Name, Author: book.Author}

	c.JSON(200, responseBook)
}

func (h *BookHandler) PostBook(c *gin.Context) {
	var query CreateBookRequestDTO
	if err := c.ShouldBindJSON(&query); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	input := service.CreateBookInputDTO{Name: query.Name, Author: query.Author}
	sOut, err := h.service.PostBook(input)
	if err != nil {
		// FIX status code  & error(make castom)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	responseBook := CreateBookResponseDTO{ID: sOut.ID}

	c.JSON(200, responseBook)
}
