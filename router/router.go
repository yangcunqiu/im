package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRouter(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

}
