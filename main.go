package main

import (
	"Go-Shop/connection"
	"Go-Shop/routes"

	_ "Go-Shop/docs"
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

// @host 127.0.0.1:2020
// @BasePath /v1

func main() {
	e := routes.Init()
	e.Logger.Fatal(e.Start(":2020"))
	connection.Init()
}
