package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

// e is the echo server
var e *echo.Echo

// init is called before the main function and sets up the echo server
func init() {
	e = echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// This only works in combination with the `vercel.json` routes.
	e.Use(middleware.Static("public"))

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<p>Hello, World!</p><p><img src='/example.png' alt='a blue square' /></p>")
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
