package config

import (
	"what-to-watch/pkg/log"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/tkanos/gonfig"
)

//Config used to configure app
type Config struct {
	AppName    string       `json:"app"`
	Port       int          `json:"port"`
	LogLevel   logrus.Level `json:"log_level"`
	StackTrace bool         `json:"stack_trace"`
}

//Marshal retrieves configurables for the app
func Marshal() (conf Config, err error) {
	log.Trace("marshaling config(s)...")

	err = gonfig.GetConf("config/application.json", &conf)
	if err != nil {
		return conf, errors.Wrap(err, "could not read config")
	}

	log.Configure(log.Logger{
		App:       conf.AppName,
		Level:     conf.LogLevel,
		ShowStack: conf.StackTrace,
	})

	return conf, err
}
