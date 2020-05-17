package service

import (
	"fmt"
	"what-to-watch/internal/controller"
	"what-to-watch/internal/model"
	"what-to-watch/pkg/log"

	"github.com/pkg/errors"
)

//DependencyCheck test dependencies
func DependencyCheck(ctrl controller.Controller) error {
	return ctrl.Datasource.Ping()
}

//AddMovie inserts movie
func AddMovie(ctrl controller.Controller, movie model.Movie) error {
	log.Trace("inserting  new movie...")

	err := ctrl.Datasource.Insert(movie)
	if err != nil {
		return errors.Wrapf(err, "could not add "+movie.Name)
	}

	log.Info(fmt.Sprintf("added %v to database", movie.Name))

	return err
}
