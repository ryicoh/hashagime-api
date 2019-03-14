package entity

import "github.com/jinzhu/gorm"

type Event struct {
	gorm.Model
	Title       string `validate:"required"`
	Description string `gorm:"type:text" validate:"required"`
	MarkDown    string `gorm:"type:longtext"`
	Image       string
	Priority    *Priority
	Tags        []*Tag `gorm:"many2many:event_tags"`
}
