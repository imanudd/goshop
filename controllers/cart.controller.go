package controllers

import (
	"Go-Shop/middleware"
	"Go-Shop/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddToCart(c echo.Context) error {
	user_id := middleware.GetID(c)
	product_id := c.FormValue("product_id")
	qty := c.FormValue("qty")
	conv_qty, _ := strconv.Atoi(qty)
	conv_product_id, _ := strconv.Atoi(product_id)
	result, err := models.AddToCart(user_id, conv_qty, conv_product_id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, result)
}

func ShowMyCart(c echo.Context) error {
	user_id := middleware.GetID(c)
	res, err := models.ShowMyCart(user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, res)
}

func DeleteCart(c echo.Context) error {
	user_id := middleware.GetID(c)
	cart_id := c.FormValue("cart_id")
	conv_cart_id, _ := strconv.Atoi(cart_id)
	result, err := models.DeleteCart(conv_cart_id, user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, result)
}
