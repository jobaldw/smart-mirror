package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"what-to-watch/internal/controller"
	"what-to-watch/internal/model"
	"what-to-watch/internal/service"
	"what-to-watch/pkg/config"
	"what-to-watch/pkg/log"
	"what-to-watch/pkg/router"

	"github.com/gorilla/mux"
)

//Start creates and runs the app server
func Start(conf config.Config, ctrl controller.Controller) {
	log.Trace("initializing router...")

	router := mux.NewRouter()
	router.HandleFunc("/health", health()).Methods(http.MethodGet)
	router.HandleFunc("/ready", ready(ctrl)).Methods(http.MethodGet)

	router.HandleFunc("/movie", addMovie(ctrl)).Methods(http.MethodPost)

	log.Info(fmt.Sprintf("API is up and running on port :%d", conf.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), router))
}

func health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := struct {
			Status string
		}{
			Status: "Up",
		}

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

		if err := service.DependencyCheck(ctrl); err != nil {
			resp.Status = "Down"
			resp.Err = err.Error()
			router.Response(w, http.StatusGatewayTimeout, resp)

			return
		}

		resp.MSG = "API is ready"
		router.Response(w, http.StatusOK, resp)
	}
}

func addMovie(ctrl controller.Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var movie = model.Movie{}

		resp := struct {
			MSG string `json:"msg,omitempty"`
			Err string `json:"error,omitempty"`
		}{}

		if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
			resp.Err = err.Error()
			router.Response(w, http.StatusBadRequest, resp)
			return
		}

		if err := service.AddMovie(ctrl, movie); err != nil {
			resp.Err = err.Error()
			router.Response(w, http.StatusUnprocessableEntity, resp)
			return
		}

		resp.MSG = "added " + movie.Name
		router.Response(w, http.StatusOK, resp)
	}
}
