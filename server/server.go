package server

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// Start runs the webserver
func Start(config Configuration) {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(config.Port)))
}
