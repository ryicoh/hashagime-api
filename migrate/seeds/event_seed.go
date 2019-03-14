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

	for i := 0; i < 10; i++ {
		event := entity.Event{
			Title:       "Sample Title " + string(i+1),
			Description: "Sample Description",
			Image:       "Sample Image",
		}
		err = db.Create(&event).Error
		if err != nil {
			spew.Dump(err)
		}
	}
	db.Close()
}
