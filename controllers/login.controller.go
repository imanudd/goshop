package controllers

import (
	"Go-Shop/middleware"
	"Go-Shop/models"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CheckLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	id, res, err := models.CheckLogin(username, password)
	fmt.Println(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	Claims := &middleware.JwtCustomClaims{
		Id:       id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	Token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)
	t, err := Token.SignedString([]byte("secret"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	if !res {
		return echo.ErrUnauthorized
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
