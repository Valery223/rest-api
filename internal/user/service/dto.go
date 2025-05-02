package service

type RegistateUserInputDTO struct {
	Name     string
	Email    string
	Password string
}

type LoginUserDTO struct {
	Email    string
	Password string
}

type RegistateUserOutputDTO struct {
	ID int
}

type UserDTO struct {
	ID    int
	Name  string
	Email string
}
