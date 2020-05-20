package show

import (
	"encoding/json"
	"fmt"
	"net/http"

	"what-to-watch/internal/controller/show"
	"what-to-watch/internal/model"

	"what-to-watch/pkg/log"
	"what-to-watch/pkg/router"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

//Create show
func Create(ctrl show.Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		show := model.Show{}
		resp := router.Resp{}

		if err := json.NewDecoder(r.Body).Decode(&show); err != nil {
			resp.Err = err.Error()
			log.Entry.WithFields(logrus.Fields{"method": "Create", "title": show.Name}).Error(err)
			router.Response(w, http.StatusBadRequest, resp)
			return
		}

		id, err := ctrl.Add(show)
		if err != nil {
			resp.Err = err.Error()
			log.Entry.WithFields(logrus.Fields{"method": "Create", "title": show.Name}).Error(err)
			router.Response(w, http.StatusUnprocessableEntity, resp)
			return
		}

		resp.ID = id
		resp.MSG = "created show"

		log.Entry.WithFields(logrus.Fields{"method": "Create", "id": id, "title": show.Name}).Info(resp.MSG)
		router.Response(w, http.StatusOK, resp)
	}
}

//Retrieve show
func Retrieve(ctrl show.Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		resp := router.Resp{}

		show, err := ctrl.Get(id)
		if err != nil {
			resp.Err = err.Error()
			log.Entry.WithFields(logrus.Fields{"method": "Retrieve"}).Error(err)
			router.Response(w, http.StatusUnprocessableEntity, resp)
			return
		}

		resp.ID = id
		resp.MSG = "retrieved show"
		resp.Show = show

		log.Entry.WithFields(logrus.Fields{"method": "Retrieve", "id": id}).Info(resp.MSG)
		router.Response(w, http.StatusOK, resp)
	}
}

//RetrieveAll show
func RetrieveAll(ctrl show.Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := router.Resp{}

		vars := mux.Vars(r)
		id := vars["id"]

		params := r.URL.Query()

		shows, count, err := ctrl.GetMany(params)
		if err != nil {
			resp.Err = err.Error()
			log.Entry.WithFields(logrus.Fields{"method": "Retrieve"}).Error(err)
			router.Response(w, http.StatusUnprocessableEntity, resp)
			return
		}

		if count == 0 {
			resp.Err = "no match found"
			log.Entry.WithFields(logrus.Fields{"method": "RetrieveAll"}).Error(resp.Err)
			router.Response(w, http.StatusNotFound, resp)
			return
		}

		resp.ID = id
		resp.MSG = fmt.Sprintf("retrieved %d shows", count)
		resp.Show = shows

		log.Entry.WithFields(logrus.Fields{"method": "Retrieve", "id": id}).Info(resp.MSG)
		router.Response(w, http.StatusOK, resp)
	}
}

//Delete show
func Delete(ctrl show.Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		resp := router.Resp{}

		count, err := ctrl.Remove(id)
		if err != nil {
			resp.Err = err.Error()
			log.Entry.WithFields(logrus.Fields{"method": "Delete"}).Error(err)
			router.Response(w, http.StatusUnprocessableEntity, resp)
			return
		}

		if count == 0 {
			resp.Err = "could not find show"
			log.Entry.WithFields(logrus.Fields{"method": "Delete"}).Error(resp.Err)
			router.Response(w, http.StatusNotFound, resp)
			return
		}

		resp.ID = id
		resp.MSG = fmt.Sprintf("deleted %v show(s)", count)

		log.Entry.WithFields(logrus.Fields{"method": "Delete", "id": id, "removed": count}).Info(resp.MSG)
		router.Response(w, http.StatusOK, resp)
	}
}
