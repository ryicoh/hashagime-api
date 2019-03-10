package repository

import (
	"detoplan-go/entity"
	"fmt"
	"github.com/jinzhu/gorm"
)

type PlanRepository struct {
	Conn *gorm.DB
}

func (r *PlanRepository) FindAll() (plans []*entity.Plan, err error) {
	err = r.Conn.Find(&plans).Error
	return
}

func (r *PlanRepository) FindOne(id uint) (plan entity.Plan, err error) {
	if err = r.Conn.First(&plan, id).Error; err != nil {
		return
	}
	fmt.Printf("%#v", err)
	var events []*entity.Event
	err = r.Conn.Table("events").
		Joins("inner join `plan_events` on `events`.`id` = `plan_events`.`event_id`").
		Where("`plan_id` = ?", plan.ID).
		Order("`order` asc").Find(&events).Error
	plan.Events = events
	return
}

func (r *PlanRepository) Create(plan *entity.Plan) (err error) {
	err = r.Conn.Create(plan).Error
	for index, event := range plan.Events {
		pe := entity.PlanEvent{
			PlanID:  plan.ID,
			EventID: event.ID,
			Order:   uint(index + 1),
		}
		if err = r.Conn.Create(&pe).Error; err != nil {
			return err
		}
	}
	return
}

func (r *PlanRepository) Update(id uint, plan *entity.Plan) (err error) {
	before, err := r.FindOne(id)
	if err != nil {
		return
	}
	err = r.Conn.Model(before).Updates(plan).Error

	var plan_events []entity.PlanEvent
	r.Conn.Where(entity.PlanEvent{PlanID: id}).Find(&plan_events)

	for _, plan_event := range plan_events {
		is_contain := false
		for _, event := range plan.Events {
			if plan_event.EventID == event.ID {
				is_contain = true
				break
			}
		}
		if !is_contain {
			if err = r.Conn.Delete(&plan_event).Error; err != nil {
				return err
			}
		}
	}

	for index, event := range plan.Events {
		is_contain := false
		for _, plan_event := range plan_events {
			if plan_event.EventID == event.ID {
				is_contain = true
				break
			}
		}
		if !is_contain {
			pe := entity.PlanEvent{
				PlanID:  before.ID,
				EventID: event.ID,
				Order:   uint(index + 1),
			}
			if err = r.Conn.Create(&pe).Error; err != nil {
				return err
			}
		}
	}

	for index, event := range plan.Events {
		pe := entity.PlanEvent{
			PlanID:  before.ID,
			EventID: event.ID,
		}

		if err = r.Conn.Where(&pe).First(&pe).Error; err != nil {
			return err
		}

		if pe.Order != uint(index+1) {
			var pes []entity.PlanEvent
			r.Conn.Find(&pes)

			if err = r.Conn.Model(&pe).Update("order", uint(index+1)).Error; err != nil {
				return err
			}
		}
	}
	return
}

func (r *PlanRepository) Delete(id uint) (err error) {
	plan, err := r.FindOne(id)
	if err != nil {
		return
	}
	err = r.Conn.Delete(&plan).Error
	return
}
