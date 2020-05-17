package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"what-to-watch/internal/controller"
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

	log.Entry.WithFields(logrus.Fields{"method": "Start", "port": conf.Port}).Info("API is up and running...")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), router))
}

func health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := struct {
			Status string
		}{
			Status: "Up",
		}

		log.Entry.WithFields(logrus.Fields{"method": "health", "status": resp.Status}).Info("API is healthy")
		router.Response(w, http.StatusOK, resp)
	}
}

func ready(ctrl controller.Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := struct {
			Status string
			MSG    string `json:"msg,omitempty"`
			Err    string `json:"error,omitempty"`
		}{
			Status: "Up",
		}

		if err := ctrl.DependencyCheck(); err != nil {
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
		var movie = model.Movie{}

		resp := struct {
			ID  interface{} `json:"id,omitempty"`
			MSG string      `json:"msg,omitempty"`
			Err string      `json:"error,omitempty"`
		}{}

		if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
			resp.Err = err.Error()
			log.Entry.WithFields(logrus.Fields{"method": "addMovie", "title": movie.Name}).Error(err)
			router.Response(w, http.StatusBadRequest, resp)
			return
		}

		id, err := ctrl.AddMovie(movie)
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
