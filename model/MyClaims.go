package model

import (
	"github.com/golang-jwt/jwt/v4"
)

type MyClaims struct {
	ID    uint
	Name  string
	Phone string
	Email string
	jwt.RegisteredClaims
}
