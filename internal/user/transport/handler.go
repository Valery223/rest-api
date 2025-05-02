package transport

import (
	"errors"
	"learn/rest-api/internal/errdefs"
	"learn/rest-api/internal/user/service"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	RegistrationUser(service.RegistateUserInputDTO) (service.RegistateUserOutputDTO, error)
	GetUserByID(int) (service.UserDTO, error)
}

type UserHandler struct {
	service UserService
}

func NewUserHandler(uS UserService) *UserHandler {
	return &UserHandler{service: uS}
}

func (uh *UserHandler) PostUser(c *gin.Context) {
	var cUser CreateUserRequestDTO
	if err := c.ShouldBindJSON(&cUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	input := service.RegistateUserInputDTO{
		Name:     cUser.Name,
		Email:    cUser.Email,
		Password: cUser.Password,
	}

	sID, err := uh.service.RegistrationUser(input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	responceUser := CreateUserResponseDTO{ID: sID.ID}

	c.JSON(200, responceUser)
}

func (uh *UserHandler) GetUser(c *gin.Context) {
	var req struct {
		ID int `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := uh.service.GetUserByID(req.ID)
	if err != nil {
		if errors.Is(err, errdefs.ErrNotFound) {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	resp := GetUserResponseDTO{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	c.JSON(200, resp)
}
