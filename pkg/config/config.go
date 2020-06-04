package config

import (
	"what-to-watch/pkg/log"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/tkanos/gonfig"
)

//Config application
type Config struct {
	AppName    string       `json:"app"`
	Port       int          `json:"port"`
	LogLevel   logrus.Level `json:"log_level"`
	StackTrace bool         `json:"stack_trace"`

	MongoURI    string            `json:"mongo_uri"`
	Database    string            `json:"database"`
	Collections map[string]string `json:"collections"`

	Clients map[string]Client `json:"clients"`
}

//Client API
type Client struct {
	URL string `json:"url"`
}

//Marshal configurables for the application
func Marshal() (conf Config, err error) {
	err = gonfig.GetConf("config/application.json", &conf)
	if err != nil {
		return conf, errors.WithMessage(err, "could not read config")
	}

	log.Configure(log.Logger{
		App:       conf.AppName,
		Level:     conf.LogLevel,
		ShowStack: conf.StackTrace,
	})

	return conf, err
}
