package controllers

import (
	"Go-Shop/middleware"
	"Go-Shop/models"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func AddToCart(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middleware.JwtCustomClaims)
	userid := claims.Id

	productId := c.FormValue("productid")
	qty := c.FormValue("qty")
	qtyint, _ := strconv.Atoi(qty)
	productIdint, _ := strconv.Atoi(productId)
	result, err := models.AddToCart(userid, qtyint, productIdint)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ShowMyCart(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middleware.JwtCustomClaims)
	userid := claims.Id

	res, err := models.ShowMyCart(userid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, res)
}

func DeleteCart(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middleware.JwtCustomClaims)
	userid := claims.Id

	id := c.FormValue("cart_id")
	idint, _ := strconv.Atoi(id)
	result, err := models.DeleteCart(idint, userid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, result)
}
