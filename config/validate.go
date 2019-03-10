package config

import (
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func NewValidate(e *echo.Echo) *echo.Echo {
	e.Validator = &CustomValidator{validator: validator.New()}
	return e
}
