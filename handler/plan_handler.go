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

type PlanHandler struct {
	Repo *repository.PlanRepository
}

func (h *PlanHandler) Index(c echo.Context) (err error) {
	plans, err := h.Repo.FindAll()
	if err != nil {
		return BadRequest(c, err)
	}
	return c.JSON(http.StatusOK, Response{IsSuccess: true, Result: plans})
}

func (h *PlanHandler) Show(c echo.Context) (err error) {
	id := c.Param("id")
	if id == "" {
		return BadRequest(c, errors.New("id not found."))
	}
	plan, err := h.Repo.FindOne(cast.ToUint(id))
	if err != nil {
		return NotFound(c, err)
	}
	return c.JSON(http.StatusOK, Response{IsSuccess: true, Result: plan})
}

func (h *PlanHandler) Create(c echo.Context) (err error) {
	var plan entity.Plan
	if err = c.Bind(&plan); err != nil {
		return BadRequest(c, err)
	}
	if err := c.Validate(&plan); err != nil {
		return BadRequest(c, err.(validator.ValidationErrors))
	}

	if err = h.Repo.Create(&plan); err != nil {
		return InternalServerError(c, err)
	}
	return c.JSON(http.StatusOK, Response{IsSuccess: true, Result: nil})
}

func (h *PlanHandler) Update(c echo.Context) (err error) {
	id := c.Param("id")
	if id == "" {
		return BadRequest(c, errors.New("id not found."))
	}

	var plan entity.Plan
	if err = c.Bind(&plan); err != nil {
		return BadRequest(c, err)
	}
	if err := c.Validate(&plan); err != nil {
		return BadRequest(c, err.(validator.ValidationErrors))
	}

	if err = h.Repo.Update(cast.ToUint(id), &plan); err != nil {
		return InternalServerError(c, err)
	}
	return c.JSON(http.StatusOK, Response{IsSuccess: true, Result: nil})
}

func (h *PlanHandler) Delete(c echo.Context) (err error) {
	id := c.Param("id")
	if id == "" {
		return BadRequest(c, errors.New("id not found."))
	}
	if err = h.Repo.Delete(cast.ToUint(id)); err != nil {
		return InternalServerError(c, err)
	}
	return c.JSON(http.StatusOK, Response{IsSuccess: true, Result: nil})
}
