package repository

import (
	"database/sql"
	"fmt"
	"learn/rest-api/internal/errdefs"
)

type BookStorage struct {
	db *sql.DB
}

func NewBookStorage(db *sql.DB) *BookStorage {
	return &BookStorage{db: db}
}

func (bs *BookStorage) GetBookById(id int) (BookModel, error) {
	stmt, err := bs.db.Prepare("SELECT * FROM book WHERE id = ?")
	if err != nil {
		return BookModel{}, fmt.Errorf("failed to prepare select statement: %w", err)
	}
	defer stmt.Close()

	var bE BookModel
	err = stmt.QueryRow(id).Scan(&bE.ID, &bE.Name, &bE.Author)
	if err != nil {
		if err == sql.ErrNoRows {
			return BookModel{}, errdefs.ErrNotFound
		}
		return BookModel{}, fmt.Errorf("failed to scan book: %w", err)
	}

	return bE, nil
}

func (bs *BookStorage) SaveBook(book BookModel) (int, error) {
	stmt, err := bs.db.Prepare("INSERT INTO book (name, author) VALUES (?, ?) RETURNING id")
	if err != nil {
		return 0, fmt.Errorf("failed to prepare insert statement: %w", err)
	}
	defer stmt.Close()

	var id int
	err = stmt.QueryRow(book.Name, book.Author).Scan(&id)
	if err != nil {
		// Возможно ошибки на уникальность проверять, но как?
		return 0, fmt.Errorf("failed to execute insert: %w", err)
	}

	return id, nil
}
