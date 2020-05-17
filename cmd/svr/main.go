package main

import (
	"what-to-watch/internal/controller"
	"what-to-watch/internal/server"
	"what-to-watch/pkg/config"
	"what-to-watch/pkg/log"

	"github.com/pkg/errors"
)

func main() {
	if err := run(); err != nil {
		log.Error(errors.Wrapf(err, "could not run app"))
	}
}

func run() error {
	conf, err := config.Marshal()
	if err != nil {
		return errors.Wrapf(err, "could not marshal config")
	}

	ctrl, err := controller.New(conf)
	if err != nil {
		return errors.Wrapf(err, "could not create new controller")
	}

	server.Start(conf, ctrl)

	return err
}
