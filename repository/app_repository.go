package repository

import (
	"github.com/jinzhu/gorm"
)

type AppRepository struct {
	PlanRepository
	EventRepository
	PlanEventRepository
}

func NewRepository(db *gorm.DB) *AppRepository {
	r := &AppRepository{}
	r.PlanRepository = PlanRepository{Conn: db}
	r.EventRepository = EventRepository{Conn: db}
	r.PlanEventRepository = PlanEventRepository{Conn: db}
	return r
}
