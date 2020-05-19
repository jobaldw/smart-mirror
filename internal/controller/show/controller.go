package show

import (
	"context"
	"time"

	"what-to-watch/internal/controller"
	"what-to-watch/internal/db/mgo"
	"what-to-watch/internal/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Controller dependecies needed for the app
type Controller struct {
	Datasource mgo.Mongo
}

// Init show controller
func Init(ctrl controller.Controller) Controller {
	return Controller{
		Datasource: ctrl.Datasource,
	}
}

//Add show
func (c *Controller) Add(show model.Show) (interface{}, error) {
	show.Created = time.Now()
	show.Updated = show.Created
	show.Watched = false

	result, err := c.Datasource.Insert(show, show.Name, mgo.ShowKey)
	if err != nil {
		return nil, err
	}

	return result.InsertedID, err
}

//Get show
func (c *Controller) Get(id string) (show model.Show, err error) {
	hex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return show, err
	}

	result := c.Datasource.FindOne(hex, mgo.ShowKey)
	if err != nil {
		return show, err
	}

	if err = result.Decode(&show); err != nil {
		return show, err
	}

	return show, err
}

//GetMany show
func (c *Controller) GetMany() (shows []model.Show, count int, err error) {
	show := model.Show{}

	results, err := c.Datasource.FindMany(mgo.ShowKey)
	if err != nil {
		return shows, count, err
	}

	defer results.Close(context.TODO())
	for results.Next(context.TODO()) {
		err = results.Decode(&show)
		if err != nil {
			return shows, count, err
		}

		count++
		shows = append(shows, show)
	}
	if err = results.Err(); err != nil {
		return shows, count, err
	}

	return shows, count, err
}

//Remove show
func (c *Controller) Remove(id string) (int, error) {
	hex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}

	result, err := c.Datasource.Delete(hex, mgo.ShowKey)
	if err != nil {
		return 0, err
	}

	return int(result.DeletedCount), err
}
