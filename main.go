package main

import (
	"Go-Shop/connection"
	"Go-Shop/routes"
)

func main() {
	e := routes.Init()
	e.Logger.Fatal(e.Start(":2020"))
	connection.Init()
}
