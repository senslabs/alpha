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

func UserEndpointMain(r *mux.Router) {
	r.HandleFunc("/api/user-endpoints/create", CreateUserEndpoint)
	r.HandleFunc("/api/user-endpoints/batch/create", BatchCreateUserEndpoint)
	
	r.HandleFunc("/api/user-endpoints/update", UpdateUserEndpointWhere)
	r.HandleFunc("/api/user-endpoints/find", FindUserEndpoint).Queries("limit", "{limit}")
	r.HandleFunc("/api/user-endpoints/delete", DeleteUserEndpoint)
}

func UserEndpointRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateUserEndpoint(w http.ResponseWriter, r *http.Request) {
	defer UserEndpointRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertUserEndpoint(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateUserEndpoint(w http.ResponseWriter, r *http.Request) {
	defer UserEndpointRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertUserEndpoint(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateUserEndpointWhere(w http.ResponseWriter, r *http.Request) {
	defer UserEndpointRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateUserEndpointWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindUserEndpoint(w http.ResponseWriter, r *http.Request) {
	defer UserEndpointRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindUserEndpoint(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindUserEndpoint: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteUserEndpoint(w http.ResponseWriter, r *http.Request) {
	defer UserEndpointRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteUserEndpoint(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteUserEndpoint: %d", n)
	types.MarshalInto(n, w)
}