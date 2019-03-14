package entity

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model
	Name string `validate:"required"`
}
