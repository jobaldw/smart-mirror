package main

import (
	"what-to-watch/internal/controller"
	"what-to-watch/internal/server"
	"what-to-watch/pkg/config"
	"what-to-watch/pkg/log"
)

func main() {
	if err := run(); err != nil {
		log.Error(err)
	}
}

func run() error {
	conf, err := config.Marshal()
	if err != nil {
		return err
	}

	ctrl, err := controller.New(conf)
	if err != nil {
		return err
	}

	server.Start(conf, ctrl)

	return err
}
