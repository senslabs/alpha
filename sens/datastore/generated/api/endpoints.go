package api

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/senslabs/alpha/sens/datastore/generated/models/fn"
	"github.com/senslabs/alpha/sens/errors"
	"github.com/senslabs/alpha/sens/httpclient"
	"github.com/senslabs/alpha/sens/logger"
	"github.com/senslabs/alpha/sens/types"
)

func EndpointMain(r *mux.Router) {
	r.HandleFunc("/api/endpoints/create", CreateEndpoint)
	r.HandleFunc("/api/endpoints/batch/create", BatchCreateEndpoint)
	
	r.HandleFunc("/api/endpoints/{id}/update", UpdateEndpoint)
	r.HandleFunc("/api/endpoints/{id}/get", GetEndpoint)
    
	r.HandleFunc("/api/endpoints/update", UpdateEndpointWhere)
	r.HandleFunc("/api/endpoints/find", FindEndpoint).Queries("limit", "{limit}")
}

func EndpointRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateEndpoint(w http.ResponseWriter, r *http.Request) {
	defer EndpointRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertEndpoint(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateEndpoint(w http.ResponseWriter, r *http.Request) {
	defer EndpointRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertEndpoint(data)
	w.WriteHeader(http.StatusOK)
}


func UpdateEndpoint(w http.ResponseWriter, r *http.Request) {
	defer EndpointRecovery(w)
	vars := mux.Vars(r)
	id := vars["id"]
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateEndpoint(id, data)
	w.WriteHeader(http.StatusOK)
}

func GetEndpoint(w http.ResponseWriter, r *http.Request) {
	defer EndpointRecovery(w)
	vars := mux.Vars(r)
	id := vars["id"]
	m := fn.SelectEndpoint(id)
	types.MarshalInto(m, w)
}


func UpdateEndpointWhere(w http.ResponseWriter, r *http.Request) {
	defer EndpointRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateEndpointWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindEndpoint(w http.ResponseWriter, r *http.Request) {
	defer EndpointRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindEndpoint(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindEndpoint: %#v", m)
	types.MarshalInto(m, w)
}
