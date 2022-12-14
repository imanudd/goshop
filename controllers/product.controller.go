package controllers

import (
	"Go-Shop/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SearchByCategory(c echo.Context) error {
	category := c.FormValue("category")
	fmt.Println(category)

	res, err := models.SearchByCategory(category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, res)
}
