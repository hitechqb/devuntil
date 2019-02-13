package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	server := echo.New()
	server.Use(middleware.Logger())
	server.Use(middleware.Gzip())

	// load all routes to server
	loadRoutes(server)

	// Start server and log error
	server.Logger.Fatal(server.Start(":2019"))
}
