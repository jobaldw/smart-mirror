package mgo

import (
	"context"
	"encoding/base64"
	"time"

	"what-to-watch/internal/model"
	"what-to-watch/pkg/log"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//DBO configurations for mongo
type DBO struct {
	URI         string
	Database    string
	Collections map[string]string
	Ctx         context.Context
	Client      *mongo.Client
	Session     mongo.Session
}

//Connect attempt to connect to a database
func (m *DBO) Connect(uri string) error {
	log.Trace("connecting to mongo database...")

	decodedURI, err := base64.StdEncoding.DecodeString(uri)
	if err != nil {
		return errors.Wrapf(err, "could not decode uri")
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(string(decodedURI)))
	if err != nil {
		return errors.Wrapf(err, "could not establish a connection to mongo")
	}

	m.URI = uri
	m.Ctx = ctx
	m.Client = client

	return err
}

//Ping test connection to database
func (m *DBO) Ping() error {
	log.Trace("pinging database...")

	err := m.Client.Ping(m.Ctx, readpref.Primary())
	if err != nil {
		return errors.Wrapf(err, "could not ping database")
	}

	return err
}

//Insert adds a new record to the database
func (m *DBO) Insert(movie model.Movie) error {
	log.Trace("inserting new movie...")

	movie.Created = time.Now()
	movie.Updated = movie.Created
	movie.Watched = false

	_, err := m.Client.Database(m.Database).Collection(m.Collections["movies"]).InsertOne(m.Ctx, movie)
	if err != nil {
		return errors.Wrapf(err, "could not insert new movie")
	}

	return err
}
