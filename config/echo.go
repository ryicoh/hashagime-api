package config

import (
	"detoplan-go/handler"
	"detoplan-go/repository"
	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo"
)

func Start() {
	db, err := NewConnection()
	if err != nil {
		spew.Dump(err)
		return
	}
	r := repository.NewRepository(db)
	h := handler.NewHandler(r)

	e := echo.New()
	NewValidate(e)
	NewRoute(e, h)
	e.Logger.Fatal(e.Start(":1323"))
}
