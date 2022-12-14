package controllers

import (
	"Go-Shop/helper"
	"Go-Shop/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	username := c.FormValue("username")
	email := c.FormValue("email")
	phonenumber := c.FormValue("phonenumber")
	address := c.FormValue("address")
	password := c.FormValue("password")
	hashpass, _ := helper.HashPassword(password)
	result, err := models.Register(username, hashpass, email, phonenumber, address)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
