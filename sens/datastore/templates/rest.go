package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/senslabs/alpha/sens/datastore/api/db"
	"github.com/senslabs/alpha/sens/datastore/models"
	"github.com/senslabs/alpha/sens/logger"
	"github.com/senslabs/alpha/sens/types"
)

func AuthMain(r *mux.Router) {
	r.HandleFunc("/api/auths/create", CreateAuth)
	r.HandleFunc("/api/auths/update", UpdateAuth)
	r.HandleFunc("/api/auths/get/{id}", GetAuth)
	r.HandleFunc("/api/auths/find", FindAuth)
}

func CreateAuth(w http.ResponseWriter, r *http.Request) {
	var auth models.Auth
	if err := types.JsonUnmarshelFromReader(r.Body, &auth); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if id, err := db.CreateAuth(auth); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		fmt.Fprintln(w, id)
	}

}

func UpdateAuth(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var auth models.Auth
	if err := types.JsonUnmarshelFromReader(r.Body, &auth); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if err := db.UpdateAuth(id, auth); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func GetAuth(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if auth, err := db.GetAuthById(id); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if err := types.JsonMarshalToWrite(w, *auth); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func FindAuth(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	batch := values["batch"]
	limit := values.Get("limit")
	delete(values, "batch")
	delete(values, "limit")

	var m types.Map
	for k, _ := range values {
		m[k] = values.Get(k)
	}

	if limit == "" {
		limit = "10"
	}

	if auths, err := db.FindAuth(m, batch, limit); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if err := types.JsonMarshalToWrite(w, auths); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
