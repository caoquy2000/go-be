package controller

import (
	"go-be/internal/service"
	"net/http"

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
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello world!",
		"name": uc.userService.GetUserInfo(),
	})
}