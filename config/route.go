package config

import (
	_ "detoplan-go/docs"
	"detoplan-go/handler"
	"github.com/labstack/echo"
)

func NewRoute(e *echo.Echo, h *handler.AppHandler) {
	v1 := e.Group("/api/v1")
	{
		p := v1.Group("/plans")
		{
			p.GET("", h.PlanHandler.Index)
			p.GET("/:id", h.PlanHandler.Show)
			p.POST("", h.PlanHandler.Create)
			p.PUT("/:id", h.PlanHandler.Update)
			p.DELETE("/:id", h.PlanHandler.Delete)
		}
		e := v1.Group("/events")
		{
			e.GET("", h.EventHandler.Index)
			e.GET("/:id", h.EventHandler.Show)
			e.POST("", h.EventHandler.Create)
			e.PUT("/:id", h.EventHandler.Update)
			e.DELETE("/:id", h.EventHandler.Delete)
		}
	}
}
