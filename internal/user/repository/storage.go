package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"learn/rest-api/internal/errdefs"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(u UserModel) (int, error) {
	stmt, err := r.db.Prepare("INSERT INTO user (name, email, password) VALUES (?, ?, ?) RETURNING id")
	if err != nil {
		return 0, fmt.Errorf("failed to prepare insert for user %w", err)
	}
	defer stmt.Close()
	var id int
	err = stmt.QueryRow(u.Name, u.Email, u.Password).Scan(&id)
	if err != nil {
		// Возможно ошибки на уникальность проверять, но как?
		return 0, fmt.Errorf("failed to insert into user %w", err)
	}

	return id, nil
}

func (r *UserRepository) GetUserByID(id int) (UserModel, error) {
	stmt, err := r.db.Prepare("SELECT id, name, email FROM user WHERE id = ?")

	if err != nil {
		return UserModel{}, fmt.Errorf("failed to prerape select for user: %w", err)
	}
	defer stmt.Close()

	var uM UserModel
	err = stmt.QueryRow(id).Scan(&uM.ID, &uM.Name, &uM.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return UserModel{}, errdefs.ErrNotFound
		}
		return UserModel{}, fmt.Errorf("failed select user %w", err)
	}

	return uM, nil

}

// func (r *UserRepository) DeleteUser(ctx context.Context, id string) error {
// 	query := "DELETE FROM users WHERE id = ?"
// 	_, err := r.db.ExecContext(ctx, query, id)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
