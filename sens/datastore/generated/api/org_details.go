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

func OrgDetailMain(r *mux.Router) {
	r.HandleFunc("/api/org-details/create", CreateOrgDetail)
	r.HandleFunc("/api/org-details/batch/create", BatchCreateOrgDetail)
	r.HandleFunc("/api/org-details/update", UpdateOrgDetail)
	r.HandleFunc("/api/org-details/get/{id}", GetOrgDetail)
	r.HandleFunc("/api/org-details/find", FindOrgDetail)
}

func CreateOrgDetail(w http.ResponseWriter, r *http.Request) {
	if data, err := ioutil.ReadAll(r.Body); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if id, err := fn.InsertOrgDetail(data); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		fmt.Fprintln(w, id)
	}
}

func BatchCreateOrgDetail(w http.ResponseWriter, r *http.Request) {
	if data, err := ioutil.ReadAll(r.Body); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if id, err := fn.BatchInsertOrgDetail(data); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		fmt.Fprintln(w, id)
	}
}

func UpdateOrgDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if data, err := ioutil.ReadAll(r.Body); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if err := fn.UpdateOrgDetail(id, data); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func GetOrgDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if m, err := fn.SelectOrgDetail(id); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if err := types.JsonMarshalToWriter(w, m); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func FindOrgDetail(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	if limit == "" {
		http.Error(w, "Query param limit is mandatory", http.StatusBadRequest)
	} else if ms, err := fn.FindOrgDetail(or, and, span, limit, column, order); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if err := types.JsonMarshalToWriter(w, ms); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
