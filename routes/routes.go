package routes

import (
	"Go-Shop/controllers"
	"Go-Shop/middleware"
	"net/http"

	_ "Go-Shop/docs"

	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Go-Shop server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:2020
// @BasePath /
// @schemes http

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", HealthCheck)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.POST("/register", controllers.Register)
	e.POST("/login", controllers.CheckLogin)
	e.GET("/viewproduct", controllers.SearchByCategory)

	e.POST("/cart", controllers.AddToCart, middleware.IsAuth)
	e.GET("/cart", controllers.ShowMyCart, middleware.IsAuth)
	e.DELETE("/cart", controllers.DeleteCart, middleware.IsAuth)

	e.POST("/transaction", controllers.AddTransaction, middleware.IsAuth)
	return e
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]

func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": "Server is up and running",
	})
}
