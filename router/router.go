package router

import (
	"github.com/gin-gonic/gin"
	"im/global"
	"im/middlewares"
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
	{
		allowGroup := routerGroup.Group("/allow")
		{
			allowGroup.POST("user/login", service.Login)
			allowGroup.POST("user/register", service.RegisterUser)
		}

		OAuthGroup := routerGroup.Group("", middlewares.ValidationToken())
		{
			commonGroup := OAuthGroup.Group("/common")
			{
				commonGroup.GET("/district/list", service.GetDistrictList)
			}

			userGroup := OAuthGroup.Group("/user", middlewares.ValidationToken())
			{
				userGroup.GET("/list", service.GetUserList)
				userGroup.GET("/:id", service.GetUser)
				userGroup.POST("/update", service.UpdateUser)
				userGroup.POST("/updatePassword", service.UpdatePassword)
			}
		}
	}

}
