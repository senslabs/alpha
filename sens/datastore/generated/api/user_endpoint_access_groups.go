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

func UserEndpointAccessGroupMain(r *mux.Router) {
	r.HandleFunc("/api/user-endpoint-access-groups/create", CreateUserEndpointAccessGroup)
	r.HandleFunc("/api/user-endpoint-access-groups/batch/create", BatchCreateUserEndpointAccessGroup)
	
	r.HandleFunc("/api/user-endpoint-access-groups/update", UpdateUserEndpointAccessGroupWhere)
	r.HandleFunc("/api/user-endpoint-access-groups/find", FindUserEndpointAccessGroup).Queries("limit", "{limit}")
}

func UserEndpointAccessGroupRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateUserEndpointAccessGroup(w http.ResponseWriter, r *http.Request) {
	defer UserEndpointAccessGroupRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertUserEndpointAccessGroup(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateUserEndpointAccessGroup(w http.ResponseWriter, r *http.Request) {
	defer UserEndpointAccessGroupRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertUserEndpointAccessGroup(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateUserEndpointAccessGroupWhere(w http.ResponseWriter, r *http.Request) {
	defer UserEndpointAccessGroupRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateUserEndpointAccessGroupWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindUserEndpointAccessGroup(w http.ResponseWriter, r *http.Request) {
	defer UserEndpointAccessGroupRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindUserEndpointAccessGroup(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindUserEndpointAccessGroup: %#v", m)
	types.MarshalInto(m, w)
}
