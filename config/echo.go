package config

import (
	"github.com/labstack/echo"
)

func Start() {
	e := echo.New()
	NewMiddleware(e)
	NewRoute(e)
	e.Logger.Fatal(e.Start(":1323"))
}
