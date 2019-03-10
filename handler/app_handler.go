package handler

import (
	"detoplan-go/repository"
	"github.com/labstack/echo"
	"net/http"
)

type AppHandler struct {
	PlanHandler
	EventHandler
	PlanEventHandler
}

type Response struct {
	IsSuccess bool
	Result    interface{}
}

func BadRequest(c echo.Context, err error) error {
	return c.JSON(
		http.StatusBadRequest,
		Response{
			IsSuccess: false,
			Result:    err.Error(),
		},
	)
}

func NotFound(c echo.Context, err error) error {
	return c.JSON(
		http.StatusNotFound,
		Response{
			IsSuccess: false,
			Result:    err.Error(),
		},
	)
}

func InternalServerError(c echo.Context, err error) error {
	return c.JSON(
		http.StatusInternalServerError,
		Response{
			IsSuccess: false,
			Result:    err.Error(),
		},
	)
}

func NewHandler(r *repository.AppRepository) *AppHandler {
	h := &AppHandler{}
	h.PlanHandler = PlanHandler{Repo: &r.PlanRepository}
	h.EventHandler = EventHandler{Repo: &r.EventRepository}
	h.PlanEventHandler = PlanEventHandler{Repo: &r.PlanEventRepository}
	return h
}
