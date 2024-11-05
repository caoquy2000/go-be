package controller

import (
	"go-be/internal/service"
	"go-be/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	response.SuccessResponse(c, 20001, uc.userService.GetUserInfo())
}