package router

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

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
