package main

import (
	"fmt"
	"net/http"

	"github.com/senslabs/alpha/sens/datastore/api/db"
	"github.com/senslabs/alpha/sens/datastore/models"
	"github.com/senslabs/alpha/sens/logger"
	"github.com/senslabs/alpha/sens/types"
	"github.com/gorilla/mux"
)

func {{.Model}}Main(r *mux.Router) {
	r.HandleFunc("/api/{{.Object}}s/create", Create{{.Model}})
	r.HandleFunc("/api/{{.Object}}s/update", Update{{.Model}})
	r.HandleFunc("/api/{{.Object}}s/get/{id}", Get{{.Model}})
	r.HandleFunc("/api/{{.Object}}s/find", Find{{.Model}})
}

func Create{{.Model}}(w http.ResponseWriter, r *http.Request) {
	var {{.Object}} models.{{.Model}}
	if err := types.JsonUnmarshelFromReader(r.Body, &{{.Object}}); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if id, err := db.Create{{.Model}}({{.Object}}); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		fmt.Fprintln(w, id)
	}
}

func Update{{.Model}}(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var {{.Object}} models.{{.Model}}
	if err := types.JsonUnmarshelFromReader(r.Body, &{{.Object}}); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if err := db.Update{{.Model}}(id, {{.Object}}); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func Get{{.Model}}(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if {{.Object}}, err := db.Get{{.Model}}ById(id); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if err := types.JsonMarshalToWrite(w, *{{.Object}}); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Find{{.Model}}(w http.ResponseWriter, r *http.Request) {
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

	if {{.Object}}s, err := db.Find{{.Model}}(m, batch, limit); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if err := types.JsonMarshalToWrite(w, {{.Object}}s); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
