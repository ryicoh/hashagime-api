package main

import (
	"github.com/labstack/echo"
	"net/http"
	"server/handler"
)

func main() {
	h := func(c echo.Context) error {
		return c.String(http.StatusOK, "hello")
	}
	e := echo.New()

	g := e.Group("/api/v1")
	g.GET("/", h)

	e.Logger.Fatal(e.Start(":1323"))
}
