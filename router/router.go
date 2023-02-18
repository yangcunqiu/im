package router

import (
	"github.com/gin-gonic/gin"
	"im/global"
	"im/service"
	"net/http"
)

func RegisterRouter(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	routerGroup := r.Group(global.Config.Server.ContextPath)

	userGroup := routerGroup.Group("/user")
	userGroup.GET("/list", service.GetUserList)
	userGroup.GET("/:id", service.GetUser)
	userGroup.POST("/register", service.RegisterUser)

}
