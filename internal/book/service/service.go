package service

import (
	"fmt"
	"learn/rest-api/internal/book/repository"
)

type Storage interface {
	GetBookById(id int) (repository.BookModel, error)
	SaveBook(repository.BookModel) (int, error)
}

type BookService struct {
	db Storage
}

func NewBookService(db Storage) *BookService {
	return &BookService{db: db}
}

func (bs *BookService) GetBook(book BookInputDTO) (BookOutputDTO, error) {
	var id int = book.ID
	bookModel, err := bs.db.GetBookById(id)
	if err != nil {
		return BookOutputDTO{}, fmt.Errorf("error get book from repo: %w", err)
	}

	return BookOutputDTO{
		ID:     bookModel.ID,
		Name:   bookModel.Name,
		Author: bookModel.Author,
	}, nil
}

func (bs *BookService) CreateBook(book CreateBookInputDTO) (CreateBookOutputDTO, error) {
	var bookModel = repository.BookModel{Name: book.Name, Author: book.Author}
	id, err := bs.db.SaveBook(bookModel)
	if err != nil {
		return CreateBookOutputDTO{}, fmt.Errorf("error create book from repo: %w", err)
	}

	return CreateBookOutputDTO{
		ID: id,
	}, nil
}
