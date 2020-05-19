package router

import (
	"encoding/json"
	"net/http"

	"what-to-watch/internal/controller"
	"what-to-watch/pkg/log"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

//Resp client response
type Resp struct {
	ID     interface{} `json:"id,omitempty"`
	Status string      `json:"status,omitempty"`
	MSG    string      `json:"msg,omitempty"`
	Err    string      `json:"error,omitempty"`

	Movie interface{} `json:"movie,omitempty"`
	Show  interface{} `json:"show,omitempty"`
}

//Health of API
func Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := Resp{Status: "Up"}

		log.Entry.WithFields(logrus.Fields{"method": "health", "status": resp.Status}).Info("API is healthy")
		Response(w, http.StatusOK, resp)
	}
}

//Ready state of API
func Ready(ctrl controller.Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := Resp{Status: "Up"}

		if err := ctrl.Check(); err != nil {
			resp.Status = "Down"
			resp.Err = err.Error()
			log.Entry.WithField("method", "ready").Error(err)
			Response(w, http.StatusGatewayTimeout, resp)

			return
		}

		resp.MSG = "API is ready"
		log.Entry.WithFields(logrus.Fields{"method": "ready", "status": resp.Status}).Info(resp.MSG)
		Response(w, http.StatusOK, resp)
	}
}

//Response writes to client
func Response(w http.ResponseWriter, code int, payload interface{}) error {
	response, err := json.Marshal(payload)
	if err != nil {
		return errors.WithMessage(err, "could not marshal payload")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	_, err = w.Write(response)
	if err != nil {
		return errors.WithMessage(err, "could not write response")
	}

	return err
}
