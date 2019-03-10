package main

import (
	"detoplan-go/config"
	"detoplan-go/entity"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	db, err := config.NewConnection()
	if err != nil {
		spew.Dump(err)
	}

	err = db.DropTableIfExists(
		&entity.Event{},
		&entity.Plan{},
		&entity.PlanEvent{},
	).Error
	if err != nil {
		spew.Dump(err)
	}

	err = db.AutoMigrate(
		&entity.Event{},
		&entity.Plan{},
		&entity.PlanEvent{},
	).Error
	if err != nil {
		spew.Dump(err)
	}
}
