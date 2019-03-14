package entity

import "github.com/jinzhu/gorm"

type Priority struct {
	gorm.Model
	Name  string `validate:"required"`
	Order uint   `validate:"required"`
}
