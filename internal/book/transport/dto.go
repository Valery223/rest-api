package transport

type BookRequestDTO struct {
	ID int `uri:"id" binding:"required,min=1"`
}

type BookResponseDTO struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

type CreateBookRequestDTO struct {
	Name   string `json:"name" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type CreateBookResponseDTO struct {
	ID int `json:"id"`
}
