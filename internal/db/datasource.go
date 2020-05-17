package db

import (
	"what-to-watch/internal/db/mgo"
	"what-to-watch/pkg/config"
	"what-to-watch/pkg/log"
)

//New creates the mongo instance to be used throughout the app
func New(conf config.Config) (mgo.DBO, error) {
	log.Trace("initializing database...")

	mongo := mgo.DBO{
		URI:         conf.MongoURI,
		Database:    conf.Database,
		Collections: conf.Collections,
		Ctx:         nil,
		Client:      nil,
		Session:     nil,
	}

	return mongo, mongo.Connect(mongo.URI)
}
