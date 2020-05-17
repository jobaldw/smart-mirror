package mgo

import (
	"context"
	"encoding/base64"
	"net/url"
	"time"

	"what-to-watch/internal/model"

	"what-to-watch/pkg/config"
	"what-to-watch/pkg/log"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	movieKey = "movies"
	showKey  = "shows"

	//ErrModelNotSupported can not handle request for given model
	ErrModelNotSupported = errors.New("model no supported")
)

//Mongo configurations for mongo
type Mongo struct {
	Host string
	Name string
	User string

	URI         *url.URL
	Database    *mongo.Database
	Collections map[string]string
}

//Init mongo instance
func Init(conf config.Config) (Mongo, error) {
	rawURI, err := base64.RawStdEncoding.DecodeString(conf.MongoURI)
	if err != nil {
		return Mongo{}, err
	}

	uri, err := url.Parse(string(rawURI))
	if err != nil {
		return Mongo{}, err
	}

	return Mongo{URI: uri, Name: conf.Database}, err
}

//Connect attempt to connect to a database
func (m *Mongo) Connect() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(m.URI.String()))
	if err != nil {
		return err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	m.Database = client.Database(m.Name)
	m.Host, m.User = m.URI.Hostname(), m.URI.User.Username()

	log.Entry.WithFields(logrus.Fields{"database": m.Name, "method": "Connect"}).Debug("connected to mongo")
	return err
}

//Ping test connection to database
func (m *Mongo) Ping() error {
	if m.Database == nil {
		return errors.New("could not connect to database")
	}

	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	if err := m.Database.Client().Ping(ctx, readpref.Primary()); err != nil {
		return err
	}

	log.Entry.WithFields(logrus.Fields{"database": m.Name, "method": "Ping"}).Debug("pinged mongo")
	return nil
}

//Insert adds a new record to the database
func (m *Mongo) Insert(i interface{}) (*mongo.InsertOneResult, error) {
	now := time.Now()

	switch v := i.(type) {
	case model.Movie:
		v.Created = now
		v.Updated = v.Created
		v.Watched = false

		log.Entry.WithFields(logrus.Fields{"movie": v.Name, "database": m.Name, "method": "Insert"}).Debug("inserting movie...")
		return m.Database.Collection(m.Collections[movieKey]).InsertOne(context.Background(), v)

	case model.Show:
		v.Created = now
		v.Updated = v.Created
		v.Watched = false

		log.Entry.WithFields(logrus.Fields{"show": v.Name, "database": m.Name, "method": "Insert"}).Debug("inserting show...")
		return m.Database.Collection(m.Collections[showKey]).InsertOne(context.Background(), v)

	default:
		log.Entry.WithFields(logrus.Fields{"database": m.Name, "method": "Insert"}).Error(ErrModelNotSupported)
		return nil, ErrModelNotSupported
	}
}
