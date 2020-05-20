package mgo

import (
	"context"
	"net/url"
	"time"

	"what-to-watch/pkg/log"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"gopkg.in/mgo.v2/bson"
)

var (
	//MovieKey for movie collection
	MovieKey = "movie"

	//ShowKey for show collection
	ShowKey = "show"
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
func Init(uri *url.URL, database string, collections map[string]string) Mongo {
	return Mongo{
		URI:         uri,
		Name:        database,
		Collections: collections,
	}
}

//Connect to mongo
func (m *Mongo) Connect() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(m.URI.String()))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	m.Database = client.Database(m.Name)
	m.Host, m.User = m.URI.Hostname(), m.URI.User.Username()

	log.Entry.WithFields(logrus.Fields{"database": m.Name, "method": "Connect"}).Debug("connected to mongo")
	return err
}

//Ping mongo
func (m *Mongo) Ping() error {
	if m.Database == nil {
		return errors.New("could not connect to database")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := m.Database.Client().Ping(ctx, readpref.Primary()); err != nil {
		return err
	}

	log.Entry.WithFields(logrus.Fields{"database": m.Name, "method": "Ping"}).Debug("pinged mongo")
	return nil
}

//Insert record in mongo
func (m *Mongo) Insert(doc interface{}, title, key string) (*mongo.InsertOneResult, error) {
	log.Entry.WithFields(logrus.Fields{
		key:          title,
		"database":   m.Name,
		"collection": m.Collections[key],
		"method":     "Insert"},
	).Debug("inserting...")

	collection := m.Database.Collection(m.Collections[key])
	return collection.InsertOne(context.Background(), doc, options.InsertOne())
}

//FindOne record in mongo
func (m *Mongo) FindOne(id primitive.ObjectID, key string) *mongo.SingleResult {
	log.Entry.WithFields(logrus.Fields{
		"id":         id.Hex(),
		"database":   m.Name,
		"collection": m.Collections[key],
		"method":     "FindOne"},
	).Debug("finding...")

	filter := bson.M{"_id": id}
	collection := m.Database.Collection(m.Collections[key])
	return collection.FindOne(context.Background(), filter, options.FindOne())
}

//FindMany record in mongo
func (m *Mongo) FindMany(filter []bson.M, key string) (*mongo.Cursor, error) {
	log.Entry.WithFields(logrus.Fields{
		"database":   m.Name,
		"collection": m.Collections[key],
		"method":     "FindMany"},
	).Debug("finding...")

	// filter := bson.M{}
	collection := m.Database.Collection(m.Collections[key])
	return collection.Aggregate(context.Background(), filter)
}

//Delete record in mongo
func (m *Mongo) Delete(id primitive.ObjectID, key string) (*mongo.DeleteResult, error) {
	log.Entry.WithFields(logrus.Fields{
		"id":         id.Hex(),
		"database":   m.Name,
		"collection": m.Collections[key],
		"method":     "Delete"},
	).Debug("deleteing...")

	filter := bson.M{"_id": id}
	collection := m.Database.Collection(m.Collections[key])
	return collection.DeleteOne(context.Background(), filter, options.Delete())

}
