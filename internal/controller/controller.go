package controller

import "what-to-watch/pkg/config"

//Controller dependecies needed for the app
type Controller struct {
	Datasource string
}

//New creates a new controller
func New(conf config.Config) (ctrl Controller, err error) {

	return ctrl, err
}
