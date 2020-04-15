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

func UserMain(r *mux.Router) {
	r.HandleFunc("/api/users/create", CreateUser)
	r.HandleFunc("/api/users/batch/create", BatchCreateUser)
	
	r.HandleFunc("/api/users/update", UpdateUserWhere)
	r.HandleFunc("/api/users/find", FindUser).Queries("limit", "{limit}")
}

func UserRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	defer UserRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertUser(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateUser(w http.ResponseWriter, r *http.Request) {
	defer UserRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertUser(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateUserWhere(w http.ResponseWriter, r *http.Request) {
	defer UserRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateUserWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	defer UserRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindUser(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindUser: %#v", m)
	types.MarshalInto(m, w)
}
