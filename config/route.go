package config

import (
	"hashagime/handler"
	"github.com/labstack/echo"
)

func NewRoute(e *echo.Echo) {
	v1 := e.Group("/api/v1")
	{
		v1.POST("", handler.IndexHandler)
	}
}
