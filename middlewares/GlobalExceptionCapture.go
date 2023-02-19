package middlewares

import (
	"github.com/gin-gonic/gin"
	"im/handler"
	"log"
)

func GlobalExceptionCapture() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				handler.Fail(c, handler.InternalServerError, "")
				log.Println(r)
				return
			}
		}()
		c.Next()
	}
}
