package transport

type BookRequestDTO struct {
	ID int `uri:"id" binding:"required,min=1"`
}

type BookResponceDTO struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}
