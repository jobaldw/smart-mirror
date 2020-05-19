package db

import (
	"encoding/base64"
	"net/url"

	"what-to-watch/internal/db/mgo"

	"what-to-watch/pkg/config"

	"github.com/pkg/errors"
)

//New creates the mongo instance to be used throughout the app
func New(conf config.Config) (ds mgo.Mongo, err error) {
	rawURI, err := base64.RawStdEncoding.DecodeString(conf.MongoURI)
	if err != nil {
		return ds, err
	}

	uri, err := url.Parse(string(rawURI))
	if err != nil {
		return ds, err
	}

	ds = mgo.Init(uri, conf.Database, conf.Collections)
	if err := ds.Connect(); err != nil {
		return ds, errors.WithMessage(err, "could not connect to mongo")
	}

	return ds, nil
}
