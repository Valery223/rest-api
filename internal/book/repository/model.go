package repository

type BookModel struct {
	ID     int    `db:"id"`
	Name   string `db:"name"`
	Author string `db:"author"`
}
