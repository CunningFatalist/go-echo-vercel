package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

var e *echo.Echo

// init is called before the main function and sets up the echo server
func init() {
	e = echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
}

// Handler is the entry point for the API when deployed on Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	e.ServeHTTP(w, r)
}

// Local is the entry point for the API when running locally
func Local() {
	e.Logger.Fatal(e.Start(":1323"))
}
