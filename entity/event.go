package entity

import (
	"github.com/jinzhu/gorm"
)

type Event struct {
	gorm.Model
	Name        string `validate:"required"`
	Description string `gorm:"type:text" validate:"required"`
	Image       string
}
