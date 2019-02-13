package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func GetHomePageHandler(c echo.Context) error {
	return c.HTML(http.StatusOK, "<div><h2>Golang Blog</h2></div>")
}
