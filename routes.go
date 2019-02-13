package main

import (
	"github.com/labstack/echo"
	"github.com/lyquocnam/go-blog/controllers"
)

// Map routes to handlers
func loadRoutes(server *echo.Echo) {
	server.GET("/", controllers.GetHomePageHandler)
	server.GET("/posts", controllers.GetArticleListWithPaginationHandler)
}
