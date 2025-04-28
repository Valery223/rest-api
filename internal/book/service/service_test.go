package service_test

import (
	"errors"
	"learn/rest-api/internal/book/repository"
	"learn/rest-api/internal/book/service"
	"testing"
)

type mockStrorage struct{}

func (m *mockStrorage) GetBookById(id int) (repository.BookEntity, error) {
	if id == 9 {
		return repository.BookEntity{}, ErrNotFound
	}
	return repository.BookEntity{ID: id, Name: "Test Book", Author: "Test Author"}, nil
}

// TODO
func (m *mockStrorage) SaveBook(repository.BookEntity) (int, error) {
	return 0, nil
}

var ErrNotFound = errors.New("book not found")

func TestGetBook_Success(t *testing.T) {
	storage := &mockStrorage{}
	svc := service.NewBookService(storage)

	input := service.BookInputDTO{ID: 1}
	got, err := svc.GetBook(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := service.BookOutputDTO{ID: 1, Name: "Test Book", Author: "Test Author"}
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestGetBook_NotFound(t *testing.T) {
	storage := &mockStrorage{}
	svc := service.NewBookService(storage)

	input := service.BookInputDTO{ID: 9}
	_, err := svc.GetBook(input)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, ErrNotFound) {
		t.Errorf("expected ErrNotFound, got %v", err)
	}
}
