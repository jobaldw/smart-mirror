package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"what-to-watch/internal/controller"
	"what-to-watch/internal/db/mgo"
	"what-to-watch/internal/model"

	"what-to-watch/pkg/config"
	"what-to-watch/pkg/log"
	"what-to-watch/pkg/router"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

//Start creates and runs the app server
func Start(conf config.Config, ctrl controller.Controller) {
	log.Trace("initializing router...")

	router := mux.NewRouter()
	router.HandleFunc("/health", health()).Methods(http.MethodGet)
	router.HandleFunc("/ready", ready(ctrl)).Methods(http.MethodGet)

	router.HandleFunc("/movie", addMovie(ctrl)).Methods(http.MethodPost)
	router.HandleFunc("/movie/{id}", deleteMovie(ctrl)).Methods(http.MethodDelete)

	router.HandleFunc("/show", addShow(ctrl)).Methods(http.MethodPost)
	router.HandleFunc("/show/{id}", deleteShow(ctrl)).Methods(http.MethodDelete)

	log.Entry.WithFields(logrus.Fields{"method": "Start", "port": conf.Port}).Info("API is up and running...")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), router))
}

type resp struct {
	ID     interface{} `json:"id,omitempty"`
	Status string      `json:"status,omitempty"`
	MSG    string      `json:"msg,omitempty"`
	Err    string      `json:"error,omitempty"`
}

func health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := resp{Status: "Up"}

		log.Entry.WithFields(logrus.Fields{"method": "health", "status": resp.Status}).Info("API is healthy")
		router.Response(w, http.StatusOK, resp)
	}
}

func ready(ctrl controller.Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := resp{Status: "Up"}

		if err := ctrl.Check(); err != nil {
			resp.Status = "Down"
			resp.Err = err.Error()
			log.Entry.WithField("method", "ready").Error(err)
			router.Response(w, http.StatusGatewayTimeout, resp)

			return
		}

		resp.MSG = "API is ready"
		log.Entry.WithField("method", "ready").Info(resp.MSG)
		router.Response(w, http.StatusOK, resp)
	}
}

func addMovie(ctrl controller.Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		movie := model.Movie{}
		resp := resp{}

		if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
			resp.Err = err.Error()
			log.Entry.WithFields(logrus.Fields{"method": "addMovie", "title": movie.Name}).Error(err)
			router.Response(w, http.StatusBadRequest, resp)
			return
		}

		id, err := ctrl.Add(movie)
		if err != nil {
			resp.Err = err.Error()
			log.Entry.WithFields(logrus.Fields{"method": "addMovie", "title": movie.Name}).Error(err)
			router.Response(w, http.StatusUnprocessableEntity, resp)
			return
		}

		resp.ID = id
		resp.MSG = "added movie"

		log.Entry.WithFields(logrus.Fields{"method": "addMovie", "id": id, "title": movie.Name}).Info(resp.MSG)
		router.Response(w, http.StatusOK, resp)
	}
}

func deleteMovie(ctrl controller.Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		resp := resp{}

		count, err := ctrl.Remove(id, mgo.MovieKey)
		if err != nil {
			resp.Err = err.Error()
			log.Entry.WithFields(logrus.Fields{"method": "deleteMovie"}).Error(err)
			router.Response(w, http.StatusUnprocessableEntity, resp)
			return
		}

		if count == 0 {
			resp.Err = "could not find movie"
			log.Entry.WithFields(logrus.Fields{"method": "deleteMovie"}).Error(resp.Err)
			router.Response(w, http.StatusNotFound, resp)
			return
		}

		resp.ID = id
		resp.MSG = fmt.Sprintf("deleted %v movie(s)", count)

		log.Entry.WithFields(logrus.Fields{"method": "deleteMovie", "removed": count}).Info(resp.MSG)
		router.Response(w, http.StatusOK, resp)
	}
}

func addShow(ctrl controller.Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		show := model.Show{}
		resp := resp{}

		if err := json.NewDecoder(r.Body).Decode(&show); err != nil {
			resp.Err = err.Error()
			log.Entry.WithFields(logrus.Fields{"method": "addShow", "title": show.Name}).Error(err)
			router.Response(w, http.StatusBadRequest, resp)
			return
		}

		id, err := ctrl.Add(show)
		if err != nil {
			resp.Err = err.Error()
			log.Entry.WithFields(logrus.Fields{"method": "addShow", "title": show.Name}).Error(err)
			router.Response(w, http.StatusUnprocessableEntity, resp)
			return
		}

		resp.ID = id
		resp.MSG = "added show"

		log.Entry.WithFields(logrus.Fields{"method": "addShow", "id": id, "title": show.Name}).Info(resp.MSG)
		router.Response(w, http.StatusOK, resp)
	}
}

func deleteShow(ctrl controller.Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		resp := resp{}

		count, err := ctrl.Remove(id, mgo.ShowKey)
		if err != nil {
			resp.Err = err.Error()
			log.Entry.WithFields(logrus.Fields{"method": "deleteShow"}).Error(err)
			router.Response(w, http.StatusUnprocessableEntity, resp)
			return
		}

		if count == 0 {
			resp.Err = "could not find show"
			log.Entry.WithFields(logrus.Fields{"method": "deleteShow"}).Error(resp.Err)
			router.Response(w, http.StatusNotFound, resp)
			return
		}

		resp.ID = id
		resp.MSG = fmt.Sprintf("deleted %v show(s)", count)

		log.Entry.WithFields(logrus.Fields{"method": "deleteShow", "removed": count}).Info(resp.MSG)
		router.Response(w, http.StatusOK, resp)
	}
}
