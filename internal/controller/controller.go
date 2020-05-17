package controller

import (
	"what-to-watch/internal/db"
	"what-to-watch/internal/db/mgo"

	"what-to-watch/pkg/config"
	"what-to-watch/pkg/log"

	"github.com/pkg/errors"
)

//Controller dependecies needed for the app
type Controller struct {
	Datasource mgo.DBO
}

//New creates a new controller
func New(conf config.Config) (ctrl Controller, err error) {
	log.Trace("creating new controller...")

	datasource, err := db.New(conf)
	if err != nil {
		return ctrl, errors.Wrapf(err, "could not create new datasource")
	}

	ctrl.Datasource = datasource

	return ctrl, err
}
