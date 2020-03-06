package main

import (
	"portofolio-go/routes"
	"net/http"
	"github.com/labstack/echo"
)

func main() {
	e := routes.Web()
	// e.Static("/", "template")
	e.Static("/", "template")
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}