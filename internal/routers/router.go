package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1") 
	{
		v1.GET("/hello", firstRoute)
	}

	return r
}

func firstRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "World!",
	})
}