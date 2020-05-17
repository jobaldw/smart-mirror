package service

import (
	"net/http"

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
