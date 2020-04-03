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

func UserEndpointAccessGroupMain(r *mux.Router) {
	r.HandleFunc("/api/user-endpoint-access-groups/create", CreateUserEndpointAccessGroup)
	r.HandleFunc("/api/user-endpoint-access-groups/batch/create", BatchCreateUserEndpointAccessGroup)
	
	r.HandleFunc("/api/user-endpoint-access-groups/update", UpdateUserEndpointAccessGroupWhere)
	r.HandleFunc("/api/user-endpoint-access-groups/find", FindUserEndpointAccessGroup)
}

func CreateUserEndpointAccessGroup(w http.ResponseWriter, r *http.Request) {
	if data, err := ioutil.ReadAll(r.Body); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if id, err := fn.InsertUserEndpointAccessGroup(data); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		fmt.Fprint(w, id)
	}
}

func BatchCreateUserEndpointAccessGroup(w http.ResponseWriter, r *http.Request) {
	if data, err := ioutil.ReadAll(r.Body); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if id, err := fn.BatchInsertUserEndpointAccessGroup(data); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		fmt.Fprint(w, id)
	}
}



func UpdateUserEndpointAccessGroupWhere(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]

	if data, err := ioutil.ReadAll(r.Body); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if err := fn.UpdateUserEndpointAccessGroupWhere(or, and, span, data); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func FindUserEndpointAccessGroup(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	if limit == "" {
		http.Error(w, "Query param limit is mandatory", http.StatusBadRequest)
	} else if ms, err := fn.FindUserEndpointAccessGroup(or, and, span, limit, column, order); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if err := types.JsonMarshalToWriter(w, ms); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
