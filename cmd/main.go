package main

import (
	"Kingcrusher9764/go-server/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Static("/static", "static")

	handler.SetupAuthRoutes(e)
	handler.SetupAdminRoutes(e)

	e.Logger.Fatal(e.Start(":7000"))
}
