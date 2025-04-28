package repository

type BookEntity struct {
	ID     int    `db:"id"`
	Name   string `db:"name"`
	Author string `db:"author"`
}
