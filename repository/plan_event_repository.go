package repository

import (
	"detoplan-go/entity"
	"github.com/jinzhu/gorm"
)

type PlanEventRepository struct {
	Conn *gorm.DB
}

func (r *PlanEventRepository) FindAll() (plan_events []*entity.PlanEvent, err error) {
	err = r.Conn.Find(&plan_events).Error
	return
}

func (r *PlanEventRepository) FindOne(id uint) (*entity.PlanEvent, error) {
	var plan_event entity.PlanEvent
	err := r.Conn.First(&plan_event, id).Error
	return &plan_event, err
}

func (r *PlanEventRepository) Create(plan_event *entity.PlanEvent) error {
	err := r.Conn.Create(plan_event).Error
	return err
}

func (r *PlanEventRepository) Update(id uint, plan_event *entity.PlanEvent) (err error) {
	before, err := r.FindOne(id)
	if err != nil {
		return
	}
	err = r.Conn.Model(before).Updates(plan_event).Error
	return
}

func (r *PlanEventRepository) Delete(id uint) (err error) {
	plan_event, err := r.FindOne(id)
	if err != nil {
		return
	}
	err = r.Conn.Delete(&plan_event).Error
	return
}
