package transport

type CreateUserRequestDTO struct {
	Name     string `json:"name"  binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateUserResponseDTO struct {
	ID int `json:"id"`
}

type GetUserResponseDTO struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateUserRequestDTO struct {
	Name     *string `json:"name,omitempty"`
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
}

type UpdateUserResponseDTO struct {
	ID int `json:"id"`
}

type DeleteUserResponseDTO struct {
	Success bool `json:"success"`
}
