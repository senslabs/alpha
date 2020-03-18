package api

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/senslabs/alpha/sens/datastore/generated/models/fn"
	"github.com/senslabs/alpha/sens/logger"
	"github.com/senslabs/alpha/sens/types"
)

func OrgAuthMain(r *mux.Router) {
	r.HandleFunc("/api/org-auths/create", CreateOrgAuth)
	r.HandleFunc("/api/org-auths/batch/create", BatchCreateOrgAuth)
	r.HandleFunc("/api/org-auths/update", UpdateOrgAuth)
	r.HandleFunc("/api/org-auths/get/{id}", GetOrgAuth)
	r.HandleFunc("/api/org-auths/find", FindOrgAuth)
}

func CreateOrgAuth(w http.ResponseWriter, r *http.Request) {
	if data, err := ioutil.ReadAll(r.Body); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if id, err := fn.InsertOrgAuth(data); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		fmt.Fprintln(w, id)
	}
}

func BatchCreateOrgAuth(w http.ResponseWriter, r *http.Request) {
	if data, err := ioutil.ReadAll(r.Body); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if id, err := fn.BatchInsertOrgAuth(data); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		fmt.Fprintln(w, id)
	}
}

func UpdateOrgAuth(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if data, err := ioutil.ReadAll(r.Body); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if err := fn.UpdateOrgAuth(id, data); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func GetOrgAuth(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if m, err := fn.SelectOrgAuth(id); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if err := types.JsonMarshalToWriter(w, m); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func FindOrgAuth(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	if limit == "" {
		http.Error(w, "Query param limit is mandatory", http.StatusBadRequest)
	} else if ms, err := fn.FindOrgAuth(or, and, span, limit, column, order); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if err := types.JsonMarshalToWriter(w, ms); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
