package controller

import (
	"what-to-watch/internal/db"
	"what-to-watch/internal/db/mgo"
	"what-to-watch/internal/model"

	"what-to-watch/pkg/config"

	"github.com/pkg/errors"
)

//Controller dependecies needed for the app
type Controller struct {
	Datasource mgo.Mongo
}

//New creates a new controller
func New(conf config.Config) (Controller, error) {
	datasource, err := db.New(conf)
	if err != nil {
		return Controller{}, errors.WithMessage(err, "could not create new datasource")
	}

	return Controller{Datasource: datasource}, err
}

//DependencyCheck test dependencies
func (c *Controller) DependencyCheck() error {
	return c.Datasource.Ping()
}

//AddMovie inserts movie
func (c *Controller) AddMovie(movie model.Movie) (id interface{}, err error) {
	result, err := c.Datasource.Insert(movie)
	if err != nil {
		return id, errors.WithMessage(err, "could not add movie")
	}

	return result.InsertedID, err
}
