package controllers

import (
	"Go-Shop/middleware"
	"Go-Shop/models"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func AddTransaction(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middleware.JwtCustomClaims)
	userid := claims.Id

	paymentId, _ := strconv.Atoi(c.FormValue("payment_id"))

	result, err := models.AddTransaction(userid, paymentId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
