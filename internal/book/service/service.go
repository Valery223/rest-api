package service

import (
	"fmt"
	"learn/rest-api/internal/book/repository"
)

type Storage interface {
	GetBookById(id int) (repository.BookEntity, error)
	SaveBook(repository.BookEntity) (int, error)
}

type BookService struct {
	db Storage
}

func NewBookService(db Storage) *BookService {
	return &BookService{db: db}
}

func (bs *BookService) GetBook(book BookInputDTO) (BookOutputDTO, error) {
	var id int = book.ID
	bookEntity, err := bs.db.GetBookById(id)
	if err != nil {
		return BookOutputDTO{}, fmt.Errorf("error get book from repo: %w", err)
	}

	return BookOutputDTO{
		ID:     bookEntity.ID,
		Name:   bookEntity.Name,
		Author: bookEntity.Author,
	}, nil
}

func (bs *BookService) PostBook(book CreateBookInputDTO) (CreateBookOutputDTO, error) {
	var bookEntity = repository.BookEntity{Name: book.Name, Author: book.Author}
	id, err := bs.db.SaveBook(bookEntity)
	if err != nil {
		return CreateBookOutputDTO{}, fmt.Errorf("error create book from repo: %w", err)
	}

	return CreateBookOutputDTO{
		ID: id,
	}, nil
}
