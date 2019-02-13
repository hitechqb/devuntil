package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func GetArticleListWithPaginationHandler(c echo.Context) error {
	return c.String(http.StatusOK, "This is the Posts")
}
