package server

import (
	"fmt"
	"net/http"

	"what-to-watch/internal/controller"
	movieController "what-to-watch/internal/controller/movie"
	showController "what-to-watch/internal/controller/show"
	"what-to-watch/internal/server/movie"
	"what-to-watch/internal/server/show"

	"what-to-watch/pkg/config"
	"what-to-watch/pkg/log"
	"what-to-watch/pkg/router"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

//Start creates and runs the app server
func Start(conf config.Config, ctrl controller.Controller) {
	mCtrl, sCtrl := movieController.Init(ctrl), showController.Init(ctrl)

	routes := mux.NewRouter()
	routes.HandleFunc("/health", router.Health()).Methods(http.MethodGet)
	routes.HandleFunc("/ready", router.Ready(ctrl)).Methods(http.MethodGet)

	routes.HandleFunc("/movie", movie.Create(mCtrl)).Methods(http.MethodPost)
	routes.HandleFunc("/movie/{id}", movie.Retrieve(mCtrl)).Methods(http.MethodGet)
	routes.HandleFunc("/movie", movie.RetrieveAll(mCtrl)).Methods(http.MethodGet)
	routes.HandleFunc("/movie/{id}", movie.Delete(mCtrl)).Methods(http.MethodDelete)

	routes.HandleFunc("/show", show.Create(sCtrl)).Methods(http.MethodPost)
	routes.HandleFunc("/show/{id}", show.Retrieve(sCtrl)).Methods(http.MethodGet)
	routes.HandleFunc("/show", show.RetrieveAll(sCtrl)).Methods(http.MethodGet)
	routes.HandleFunc("/show/{id}", show.Delete(sCtrl)).Methods(http.MethodDelete)

	log.Entry.WithFields(logrus.Fields{"method": "Start", "port": conf.Port}).Info("API is up and running...")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), routes))
}
