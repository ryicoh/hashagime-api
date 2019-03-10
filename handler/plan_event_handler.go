package handler

import (
	"detoplan-go/entity"
	"detoplan-go/repository"
	"errors"
	"github.com/labstack/echo"
	"github.com/spf13/cast"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

type PlanEventHandler struct {
	Repo *repository.PlanEventRepository
}

func (h *PlanEventHandler) Index(c echo.Context) (err error) {
	events, err := h.Repo.FindAll()
	if err != nil {
		return BadRequest(c, err)
	}
	return c.JSON(http.StatusOK, Response{IsSuccess: true, Result: events})
}

func (h *PlanEventHandler) Show(c echo.Context) (err error) {
	id := c.Param("id")
	if id == "" {
		return BadRequest(c, errors.New("id not found."))
	}
	event, err := h.Repo.FindOne(cast.ToUint(id))
	if err != nil {
		return NotFound(c, err)
	}
	return c.JSON(http.StatusOK, Response{IsSuccess: true, Result: event})
}

func (h *PlanEventHandler) Create(c echo.Context) (err error) {
	var event entity.PlanEvent
	if err = c.Bind(&event); err != nil {
		return BadRequest(c, err)
	}
	if err := c.Validate(&event); err != nil {
		return BadRequest(c, err.(validator.ValidationErrors))
	}

	if err = h.Repo.Create(&event); err != nil {
		return InternalServerError(c, err)
	}
	return c.JSON(http.StatusOK, Response{IsSuccess: true, Result: nil})
}

func (h *PlanEventHandler) Update(c echo.Context) (err error) {
	id := c.Param("id")
	if id == "" {
		return BadRequest(c, errors.New("id not found."))
	}

	var event entity.PlanEvent
	if err = c.Bind(&event); err != nil {
		return BadRequest(c, err)
	}
	if err := c.Validate(&event); err != nil {
		return BadRequest(c, err.(validator.ValidationErrors))
	}

	if err = h.Repo.Update(cast.ToUint(id), &event); err != nil {
		return InternalServerError(c, err)
	}
	return c.JSON(http.StatusOK, Response{IsSuccess: true, Result: nil})
}

func (h *PlanEventHandler) Delete(c echo.Context) (err error) {
	id := c.Param("id")
	if id == "" {
		return BadRequest(c, errors.New("id not found."))
	}
	if err = h.Repo.Delete(cast.ToUint(id)); err != nil {
		return InternalServerError(c, err)
	}
	return c.JSON(http.StatusOK, Response{IsSuccess: true, Result: nil})
}
