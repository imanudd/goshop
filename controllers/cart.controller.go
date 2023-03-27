package controllers

import (
	"Go-Shop/middleware"
	"Go-Shop/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddToCart(c echo.Context) error {
	claims := middleware.GetID(c)
	product_id, _ := strconv.Atoi(c.FormValue("product_id"))
	qty, _ := strconv.Atoi(c.FormValue("qty"))

	result, err := models.AddToCart(claims.Id, qty, product_id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, result)
}

func ShowMyCart(c echo.Context) error {
	claims := middleware.GetID(c)
	res, err := models.ShowMyCart(claims.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, res)
}

func DeleteCart(c echo.Context) error {
	claims := middleware.GetID(c)
	cart_id, _ := strconv.Atoi(c.FormValue("cart_id"))
	result, err := models.DeleteCart(cart_id, claims.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusNoContent, result)
}
