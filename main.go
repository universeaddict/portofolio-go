package main

import (
	"portofolio-go/routes"
	"net/http"
	"os"
	"github.com/labstack/echo"
)

func main() {
	port := os.Getenv("PORT")
	e := routes.Web()
	// e.Static("/", "template")
	e.Static("/", "template")
	// Start server
	e.Logger.Fatal(e.Start(port))
}

// Handler
func hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}