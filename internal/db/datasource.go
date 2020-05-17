package db

import (
	"what-to-watch/internal/db/mgo"

	"what-to-watch/pkg/config"

	"github.com/pkg/errors"
)

//New creates the mongo instance to be used throughout the app
func New(conf config.Config) (mgo.Mongo, error) {
	ds, err := mgo.Init(conf)
	if err != nil {
		return ds, errors.WithMessage(err, "could not initialize mongo")
	}

	if err := ds.Connect(); err != nil {
		return ds, errors.WithMessage(err, "could not connect to mongo")
	}

	return ds, nil
}
