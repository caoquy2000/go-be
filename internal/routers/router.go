package routers

import (
	"go-be/internal/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1") 
	{
		v1.GET("/hello", firstRoute)
		v1.GET("/user/1", controller.NewUserController().GetUserByID)
	}

	return r
}

func firstRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "World!",
	})
}