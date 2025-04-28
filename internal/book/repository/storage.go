package repository

import (
	"database/sql"
	"fmt"
)

type BookStorage struct {
	db *sql.DB
}

func NewBookStorage(db *sql.DB) *BookStorage {
	return &BookStorage{db: db}
}

func (bs *BookStorage) GetBookById(id int) (BookEntity, error) {
	stmt, err := bs.db.Prepare("SELECT * FROM book WHERE id = ?")
	if err != nil {
		return BookEntity{}, fmt.Errorf("failed to prepare select statement: %w", err)
	}

	var bE BookEntity
	err = stmt.QueryRow(id).Scan(&bE.ID, &bE.Name, &bE.Author)
	if err != nil {
		if err == sql.ErrNoRows {
			return BookEntity{}, err
		}
		return BookEntity{}, err
	}

	return bE, nil
}

func (bs *BookStorage) SaveBook(book BookEntity) (int, error) {
	stmt, err := bs.db.Prepare("INSERT INTO book (name, author) VALUES (?, ?) RETURNING id")
	if err != nil {
		return 0, fmt.Errorf("failed to prepare insert statement: %w", err)
	}
	defer stmt.Close()

	var id int
	err = stmt.QueryRow(book.Name, book.Author).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to execute insert: %w", err)
	}

	return id, nil
}
