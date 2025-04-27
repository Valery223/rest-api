package service

type BookInputDTO struct {
	ID int
}

type BookOutputDTO struct {
	ID     int
	Name   string
	Author string
}
