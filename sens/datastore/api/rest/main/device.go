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

func DeviceMain(r *mux.Router) {
	r.HandleFunc("/api/devices/create", CreateDevice)
	r.HandleFunc("/api/devices/update", UpdateDevice)
	r.HandleFunc("/api/devices/get/{id}", GetDevice)
	r.HandleFunc("/api/devices/find", FindDevice)
}

func CreateDevice(w http.ResponseWriter, r *http.Request) {
	var device models.Device
	if err := types.JsonUnmarshelFromReader(r.Body, &device); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if id, err := db.CreateDevice(device); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		fmt.Fprintln(w, id)
	}
}

func UpdateDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var device models.Device
	if err := types.JsonUnmarshelFromReader(r.Body, &device); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if err := db.UpdateDevice(id, device); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func GetDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if device, err := db.GetDeviceById(id); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if err := types.JsonMarshalToWrite(w, *device); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func FindDevice(w http.ResponseWriter, r *http.Request) {
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

	if devices, err := db.FindDevice(m, batch, limit); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if err := types.JsonMarshalToWrite(w, devices); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
