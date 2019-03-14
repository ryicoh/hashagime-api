package entity

import "github.com/jinzhu/gorm"

type Plan struct {
	gorm.Model
	Title       string `validate:"required"`
	Description string `gorm:"type:text" validate:"required"`
	Markdown    string `gorm:"type:text" validate:"required"`
	Image       string
	Events      []*Event
}
