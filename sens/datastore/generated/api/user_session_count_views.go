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

func UserSessionCountViewMain(r *mux.Router) {
	r.HandleFunc("/api/user-session-count-views/create", CreateUserSessionCountView)
	r.HandleFunc("/api/user-session-count-views/batch/create", BatchCreateUserSessionCountView)
	
	r.HandleFunc("/api/user-session-count-views/update", UpdateUserSessionCountViewWhere)
	r.HandleFunc("/api/user-session-count-views/find", FindUserSessionCountView).Queries("limit", "{limit}")
	r.HandleFunc("/api/user-session-count-views/delete", DeleteUserSessionCountView)
}

func UserSessionCountViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateUserSessionCountView(w http.ResponseWriter, r *http.Request) {
	defer UserSessionCountViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	id := fn.InsertUserSessionCountView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateUserSessionCountView(w http.ResponseWriter, r *http.Request) {
	defer UserSessionCountViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	fn.BatchInsertUserSessionCountView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateUserSessionCountViewWhere(w http.ResponseWriter, r *http.Request) {
	defer UserSessionCountViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateUserSessionCountViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindUserSessionCountView(w http.ResponseWriter, r *http.Request) {
	defer UserSessionCountViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindUserSessionCountView(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindUserSessionCountView: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteUserSessionCountView(w http.ResponseWriter, r *http.Request) {
	defer UserSessionCountViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteUserSessionCountView(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteUserSessionCountView: %d", n)
	types.MarshalInto(n, w)
}