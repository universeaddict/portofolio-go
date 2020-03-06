package routes

import (
	"net/http"
	"io/ioutil"
	"github.com/labstack/echo"
	"fmt"
)


func Web() *echo.Echo {
	e := echo.New()

	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "home")
	// })

	// e.GET("/home", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "welcome home")
	// })

	// home
	e.GET("/", func(c echo.Context) error {
		content, _ := ioutil.ReadFile("frontend/index.html")
		fmt.Fprintf(c.Response(), "%s", content)
		return c.String(http.StatusBadRequest, "")
	})

	return e
}

