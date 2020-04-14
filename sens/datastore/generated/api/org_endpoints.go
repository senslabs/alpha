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

func OrgEndpointMain(r *mux.Router) {
	r.HandleFunc("/api/org-endpoints/create", CreateOrgEndpoint)
	r.HandleFunc("/api/org-endpoints/batch/create", BatchCreateOrgEndpoint)
	
	r.HandleFunc("/api/org-endpoints/update", UpdateOrgEndpointWhere)
	r.HandleFunc("/api/org-endpoints/find", FindOrgEndpoint).Queries("limit", "{limit}")
}

func OrgEndpointRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateOrgEndpoint(w http.ResponseWriter, r *http.Request) {
	defer OrgEndpointRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertOrgEndpoint(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateOrgEndpoint(w http.ResponseWriter, r *http.Request) {
	defer OrgEndpointRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertOrgEndpoint(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateOrgEndpointWhere(w http.ResponseWriter, r *http.Request) {
	defer OrgEndpointRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateOrgEndpointWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindOrgEndpoint(w http.ResponseWriter, r *http.Request) {
	defer OrgEndpointRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindOrgEndpoint(or, and, in, span, limit, column, order)
	types.MarshalInto(m, w)
}
