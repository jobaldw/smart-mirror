package movie

import (
	"encoding/json"
	"fmt"
	"net/http"

	"what-to-watch/internal/controller/movie"
	"what-to-watch/internal/model"

	"what-to-watch/pkg/log"
	"what-to-watch/pkg/router"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

//Create movie
func Create(ctrl movie.Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		movie := model.Movie{}
		resp := router.Resp{}

		if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
			resp.Err = err.Error()
			log.Entry.WithFields(logrus.Fields{"method": "Create", "title": movie.Name}).Error(err)
			router.Response(w, http.StatusBadRequest, resp)
			return
		}

		id, err := ctrl.Add(movie)
		if err != nil {
			resp.Err = err.Error()
			log.Entry.WithFields(logrus.Fields{"method": "Create", "title": movie.Name}).Error(err)
			router.Response(w, http.StatusUnprocessableEntity, resp)
			return
		}

		resp.ID = id
		resp.MSG = "created movie"

		log.Entry.WithFields(logrus.Fields{"method": "Create", "id": id, "title": movie.Name}).Info(resp.MSG)
		router.Response(w, http.StatusOK, resp)
	}
}

//Retrieve movie
func Retrieve(ctrl movie.Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		resp := router.Resp{}

		movie, err := ctrl.Get(id)
		if err != nil {
			resp.Err = err.Error()
			log.Entry.WithFields(logrus.Fields{"method": "Retrieve", "id": id}).Error(err)
			router.Response(w, http.StatusUnprocessableEntity, resp)
			return
		}

		resp.ID = id
		resp.MSG = "retrieved movie"
		resp.Movie = movie

		log.Entry.WithFields(logrus.Fields{"method": "Retrieve", "id": id}).Info(resp.MSG)
		router.Response(w, http.StatusOK, resp)
	}
}

//RetrieveAll movie
func RetrieveAll(ctrl movie.Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := router.Resp{}
		vars := mux.Vars(r)
		params := r.URL.Query()

		id := vars["id"]

		movies, count, err := ctrl.GetMany(params)
		if err != nil {
			resp.Err = err.Error()
			log.Entry.WithFields(logrus.Fields{"method": "RetrieveAll", "id": id}).Error(err)
			router.Response(w, http.StatusUnprocessableEntity, resp)
			return
		}

		if count == 0 {
			resp.MSG = "no match found"
			log.Entry.WithFields(logrus.Fields{"method": "RetrieveAll", "id": id}).Warn(resp.MSG)
			router.Response(w, http.StatusNoContent, resp)
			return
		}

		resp.ID = id
		resp.MSG = fmt.Sprintf("retrieved %d movie(s)", count)
		resp.Movie = movies

		log.Entry.WithFields(logrus.Fields{"method": "RetrieveAll", "id": id}).Info(resp.MSG)
		router.Response(w, http.StatusOK, resp)
	}
}

//Update movie
func Update(ctrl movie.Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		movie := model.Movie{}
		resp := router.Resp{}

		if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
			resp.Err = err.Error()
			log.Entry.WithFields(logrus.Fields{"method": "Update", "id": id, "title": movie.Name}).Error(err)
			router.Response(w, http.StatusBadRequest, resp)
			return
		}

		updated, err := ctrl.Update(id, movie)
		if err != nil {
			resp.Err = err.Error()
			log.Entry.WithFields(logrus.Fields{"method": "Update", "id": id, "title": movie.Name}).Error(err)
			router.Response(w, http.StatusUnprocessableEntity, resp)
			return
		}

		if updated == 0 {
			resp.MSG = "no changes detected"
			log.Entry.WithFields(logrus.Fields{"method": "Update", "id": id, "title": movie.Name}).Warn(resp.MSG)
			router.Response(w, http.StatusNoContent, resp)
			return
		}

		resp.ID = id
		resp.MSG = fmt.Sprintf("updated %d movie(s)", updated)

		log.Entry.WithFields(logrus.Fields{"method": "Update", "id": id, "title": movie.Name}).Info(resp.MSG)
		router.Response(w, http.StatusOK, resp)
	}
}

//Delete movie
func Delete(ctrl movie.Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		resp := router.Resp{}

		count, err := ctrl.Remove(id)
		if err != nil {
			resp.Err = err.Error()
			log.Entry.WithFields(logrus.Fields{"method": "Delete", "id": id}).Error(err)
			router.Response(w, http.StatusUnprocessableEntity, resp)
			return
		}

		if count == 0 {
			resp.MSG = "could not find movie"
			log.Entry.WithFields(logrus.Fields{"method": "Delete", "id": id}).Warn(resp.MSG)
			router.Response(w, http.StatusNoContent, resp)
			return
		}

		resp.ID = id
		resp.MSG = fmt.Sprintf("deleted %d movie(s)", count)

		log.Entry.WithFields(logrus.Fields{"method": "Delete", "id": id, "removed": count}).Info(resp.MSG)
		router.Response(w, http.StatusOK, resp)
	}
}
