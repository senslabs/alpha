package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/senslabs/alpha/sens/datastore/generated/models/fn"
	"github.com/senslabs/alpha/sens/logger"
	"github.com/senslabs/alpha/sens/types"
)

func GroupMain(r *mux.Router) {
	r.HandleFunc("/api/groups/create", CreateGroup)
	r.HandleFunc("/api/groups/update", UpdateGroup)
	r.HandleFunc("/api/groups/get/{id}", GetGroup)
	r.HandleFunc("/api/groups/find", FindGroup)
}

func CreateGroup(w http.ResponseWriter, r *http.Request) {
	if data, err := ioutil.ReadAll(r.Body); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if id, err := fn.InsertGroup(data); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		fmt.Fprintln(w, id)
	}
}

func UpdateGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if data, err := ioutil.ReadAll(r.Body); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if err := fn.UpdateGroup(id, data); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func GetGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if m, err := fn.SelectGroup(id); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if err := types.JsonMarshalToWrite(w, m); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func FindGroup(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	if limit == "" {
		http.Error(w, "Query param limit is mandatory", http.StatusBadRequest)
	} else if ms, err := fn.FindGroup(or, and, span, limit, column, order); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if err := types.JsonMarshalToWrite(w, ms); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
