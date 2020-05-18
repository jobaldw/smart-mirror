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

//New creates a new controller
func New(conf config.Config) (Controller, error) {
	datasource, err := db.New(conf)
	if err != nil {
		return Controller{}, errors.WithMessage(err, "could not create new datasource")
	}

	return Controller{Datasource: datasource}, err
}

//Check test dependencies
func (c *Controller) Check() error {
	return c.Datasource.Ping()
}

//Add inserts movie or show
func (c *Controller) Add(object interface{}) (interface{}, error) {
	result, err := c.Datasource.Insert(object)

	if err != nil {
		return nil, errors.WithMessagef(err, "could not add %T", object)
	}

	return result.InsertedID, err
}

//Remove deletes movie or show
func (c *Controller) Remove(id, key string) (int, error) {
	result, err := c.Datasource.Delete(id, key)

	if err != nil {
		return 0, errors.WithMessagef(err, "could not remove %v", key)
	}

	return int(result.DeletedCount), err
}
