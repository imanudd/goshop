package middleware

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JwtCustomClaims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

var IsAuth = middleware.JWTWithConfig(middleware.JWTConfig{
	Claims:     &JwtCustomClaims{},
	SigningKey: []byte("secret")})

func GetID(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	user_id := claims.Id

	return user_id
}
