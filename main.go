package main

import (
	"github.com/hermandev/go-restapi-echo/routes"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := routes.Init()

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.Logger.Fatal(app.Start(":3000"))
}
