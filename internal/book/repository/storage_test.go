package repository_test

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"learn/rest-api/internal/book/repository"
)

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open in-memory sqlite: %v", err)
	}
	_, err = db.Exec(`CREATE TABLE book (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		author TEXT NOT NULL
	)`)
	if err != nil {
		t.Fatalf("failed to create table: %v", err)
	}
	return db
}

func TestSaveBookAndGetBookById(t *testing.T) {
	db := setupTestDB(t)
	storage := repository.NewBookStorage(db)

	book := repository.BookModel{
		Name:   "Test Book",
		Author: "Test Author",
	}
	id, err := storage.SaveBook(book)
	if err != nil {
		t.Fatalf("SaveBook failed: %v", err)
	}
	if id == 0 {
		t.Fatalf("expected non-zero id, got %d", id)
	}

	got, err := storage.GetBookById(id)
	if err != nil {
		t.Fatalf("GetBookById failed: %v", err)
	}
	if got.Name != book.Name || got.Author != book.Author {
		t.Errorf("got %+v, want %+v", got, book)
	}
}

func TestGetBookById_NotFound(t *testing.T) {
	db := setupTestDB(t)
	storage := repository.NewBookStorage(db)

	_, err := storage.GetBookById(999)
	if err == nil {
		t.Error("expected error for non-existent book, got nil")
	}
}
