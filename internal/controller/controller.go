package controller

import (
	"what-to-watch/internal/db"
	"what-to-watch/internal/db/mgo"
	"what-to-watch/pkg/config"

	"github.com/pkg/errors"
)

//Controller dependecies needed for the app
type Controller struct {
	Datasource mgo.Mongo
}

//New controller
func New(conf config.Config) (Controller, error) {
	datasource, err := db.New(conf)
	if err != nil {
		return Controller{}, errors.WithMessage(err, "could not create new datasource")
	}

	return Controller{Datasource: datasource}, err
}

//Check dependencies
func (c *Controller) Check() error {
	return c.Datasource.Ping()
}
