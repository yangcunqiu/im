package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"im/global"
	"im/handler"
	"im/model"
	"log"
)

func ValidationToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			handler.Fail(c, handler.TokenEmptyError, "")
			c.Abort()
			return
		}
		// jwt解析token
		myClaims, err := parseToken(tokenString)
		if err != nil {
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorMalformed != 0 {
					handler.Fail(c, handler.TokenMalformedError, "")
					c.Abort()
					return
				} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
					handler.Fail(c, handler.TokenExpiredError, "")
					c.Abort()
					return
				} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
					handler.Fail(c, handler.TokenNoTActiveError, "")
					c.Abort()
					return
				} else {
					handler.Fail(c, handler.TokenParseError, "")
					c.Abort()
					return
				}
			} else {
				handler.Fail(c, handler.TokenParseError, "")
				c.Abort()
				return
			}
		}
		// user存到全局变量中, 方便后面直接用
		user := model.User{
			Model: gorm.Model{
				ID: myClaims.ID,
			},
			Name:  myClaims.Name,
			Phone: myClaims.Phone,
			Email: myClaims.Email,
		}
		global.User = user
		c.Next()
	}
}

func parseToken(tokenString string) (model.MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.Config.Jwt.Secret), nil
	})
	if err != nil {
		return model.MyClaims{}, err
	}
	claims, ok := token.Claims.(*model.MyClaims)
	if !ok || !token.Valid {
		return *claims, errors.New("parse token fail")
	}
	log.Printf(claims.ExpiresAt.String())
	return *claims, nil
}
