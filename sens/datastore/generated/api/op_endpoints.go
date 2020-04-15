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

func OpEndpointMain(r *mux.Router) {
	r.HandleFunc("/api/op-endpoints/create", CreateOpEndpoint)
	r.HandleFunc("/api/op-endpoints/batch/create", BatchCreateOpEndpoint)
	
	r.HandleFunc("/api/op-endpoints/update", UpdateOpEndpointWhere)
	r.HandleFunc("/api/op-endpoints/find", FindOpEndpoint).Queries("limit", "{limit}")
}

func OpEndpointRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateOpEndpoint(w http.ResponseWriter, r *http.Request) {
	defer OpEndpointRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertOpEndpoint(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateOpEndpoint(w http.ResponseWriter, r *http.Request) {
	defer OpEndpointRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertOpEndpoint(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateOpEndpointWhere(w http.ResponseWriter, r *http.Request) {
	defer OpEndpointRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateOpEndpointWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindOpEndpoint(w http.ResponseWriter, r *http.Request) {
	defer OpEndpointRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindOpEndpoint(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindOpEndpoint: %#v", m)
	types.MarshalInto(m, w)
}
