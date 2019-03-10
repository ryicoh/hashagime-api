package entity

import (
	"github.com/jinzhu/gorm"
)

type PlanEvent struct {
	gorm.Model
	PlanID  uint
	EventID uint
	Order   uint
	StartAt string `gorm:"size:255"`
	EndAt   string `gorm:"size:255"`
}
