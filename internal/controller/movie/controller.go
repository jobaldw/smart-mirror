package movie

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

// Init movie controller
func Init(ctrl controller.Controller) Controller {
	return Controller{
		Datasource: ctrl.Datasource,
	}
}

//Add movie
func (c *Controller) Add(movie model.Movie) (interface{}, error) {
	movie.Created = time.Now()
	movie.Updated = movie.Created
	movie.Watched = false

	result, err := c.Datasource.Insert(movie, movie.Name, mgo.MovieKey)
	if err != nil {
		return nil, err
	}

	return result.InsertedID, err
}

//Get movie
func (c *Controller) Get(id string) (movie model.Movie, err error) {
	hex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return movie, err
	}

	result := c.Datasource.FindOne(hex, mgo.MovieKey)
	if err != nil {
		return movie, err
	}

	if err = result.Decode(&movie); err != nil {
		return movie, err
	}

	return movie, err
}

//GetMany movie
func (c *Controller) GetMany() (movies []model.Movie, count int, err error) {
	movie := model.Movie{}
	ctx := context.TODO()

	results, err := c.Datasource.FindMany(mgo.MovieKey)
	if err != nil {
		return movies, count, err
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		err = results.Decode(&movie)
		if err != nil {
			return movies, count, err
		}

		count++
		movies = append(movies, movie)
	}
	if err = results.Err(); err != nil {
		return movies, count, err
	}

	return movies, count, err
}

//Remove movie
func (c *Controller) Remove(id string) (int, error) {
	hex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}

	result, err := c.Datasource.Delete(hex, mgo.MovieKey)
	if err != nil {
		return 0, err
	}

	return int(result.DeletedCount), err
}
