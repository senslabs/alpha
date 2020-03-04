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

func SleepMain(r *mux.Router) {
	r.HandleFunc("/api/sleeps/create", CreateSleep)
	r.HandleFunc("/api/sleeps/update", UpdateSleep)
	r.HandleFunc("/api/sleeps/get/{id}", GetSleep)
	r.HandleFunc("/api/sleeps/find", FindSleep)
}

func CreateSleep(w http.ResponseWriter, r *http.Request) {
	var sleep models.Sleep
	if err := types.JsonUnmarshelFromReader(r.Body, &sleep); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if id, err := db.CreateSleep(sleep); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		fmt.Fprintln(w, id)
	}
}

func UpdateSleep(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var sleep models.Sleep
	if err := types.JsonUnmarshelFromReader(r.Body, &sleep); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if err := db.UpdateSleep(id, sleep); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func GetSleep(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if sleep, err := db.GetSleepById(id); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if err := types.JsonMarshalToWrite(w, *sleep); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func FindSleep(w http.ResponseWriter, r *http.Request) {
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

	if sleeps, err := db.FindSleep(m, batch, limit); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if err := types.JsonMarshalToWrite(w, sleeps); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
