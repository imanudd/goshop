package controllers

import (
	"Go-Shop/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func SearchByCategory(c echo.Context) error {
	category_id := c.FormValue("category_id")
	conv_category_id, _ := strconv.Atoi(category_id)

	res, err := models.SearchByCategory(conv_category_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, res)
}
