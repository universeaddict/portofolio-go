package controllers

import (
	"net/http"
	// "io/ioutil"
	// "fmt"
	"github.com/labstack/echo"
)

func Home (c echo.Context) error {
	return c.Render(http.StatusOK, "index", "hello")
}

func About (c echo.Context) error {
	return c.Render(http.StatusOK, "about", "hello")
}