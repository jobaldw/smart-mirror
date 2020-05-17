package service

import (
	"net/http"

	"what-to-watch/internal/controller"
	"what-to-watch/pkg/router"
)

//Health checks status of app
func Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := struct {
			Status string
		}{
			Status: "Up",
		}

		router.Response(w, http.StatusOK, resp)
	}
}

//Ready checks status of app
func Ready(ctrl controller.Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := struct {
			Status string
			MSG    string `json:"msg,omitempty"`
			Err    string `json:"error,omitempty"`
		}{
			Status: "Up",
		}

		if err := ctrl.Datasource.Ping(); err != nil {
			resp.Status = "Down"
			resp.Err = err.Error()
			router.Response(w, http.StatusGatewayTimeout, resp)

			return
		}

		resp.MSG = "API is ready"
		router.Response(w, http.StatusOK, resp)
	}
}
