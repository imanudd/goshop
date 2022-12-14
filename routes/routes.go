package routes

import (
	"Go-Shop/controllers"
	"Go-Shop/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "WELCOME")
	})
	e.POST("/register", controllers.Register)
	e.POST("/login", controllers.CheckLogin)
	e.GET("/viewproduct", controllers.SearchByCategory)
	e.POST("/cart", controllers.AddToCart, middleware.IsAuth)
	e.GET("/cart", controllers.ShowMyCart, middleware.IsAuth)
	e.DELETE("/cart", controllers.DeleteCart, middleware.IsAuth)
	e.POST("/transaction", controllers.AddTransaction, middleware.IsAuth)

	// e.GET("/detailuser", controllers.FetchAllData, middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("secret")}))

	return e

}
