package server

import (
	"fmt"
	"net/http"

	"what-to-watch/internal/controller"
	"what-to-watch/internal/service"
	"what-to-watch/pkg/config"
	"what-to-watch/pkg/log"

	"github.com/gorilla/mux"
)

//Start creates and runs the app server
func Start(conf config.Config, ctrl controller.Controller) {
	log.Trace("initializing router...")

	router := mux.NewRouter()
	router.HandleFunc("/health", service.Health()).Methods(http.MethodGet)

	log.Info(fmt.Sprintf("API is up and running on port :%d", conf.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), router))
}
