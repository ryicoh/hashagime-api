package repository

import (
	"detoplan-go/entity"
	"github.com/jinzhu/gorm"
)

type EventRepository struct {
	Conn *gorm.DB
}

func (r *EventRepository) FindAll() (events []*entity.Event, err error) {
	err = r.Conn.Find(&events).Error
	return
}

func (r *EventRepository) FindOne(id uint) (*entity.Event, error) {
	var event entity.Event
	err := r.Conn.First(&event, id).Error
	return &event, err
}

func (r *EventRepository) Create(event *entity.Event) error {
	err := r.Conn.Create(event).Error
	return err
}

func (r *EventRepository) Update(id uint, event *entity.Event) (err error) {
	before, err := r.FindOne(id)
	if err != nil {
		return
	}
	err = r.Conn.Model(before).Updates(event).Error
	return
}

func (r *EventRepository) Delete(id uint) (err error) {
	event, err := r.FindOne(id)
	if err != nil {
		return
	}
	err = r.Conn.Delete(&event).Error
	return
}
