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

func UserDatedSessionViewMain(r *mux.Router) {
	r.HandleFunc("/api/user-dated-session-views/create", CreateUserDatedSessionView)
	r.HandleFunc("/api/user-dated-session-views/batch/create", BatchCreateUserDatedSessionView)
	
	r.HandleFunc("/api/user-dated-session-views/update", UpdateUserDatedSessionViewWhere)
	r.HandleFunc("/api/user-dated-session-views/find", FindUserDatedSessionView).Queries("limit", "{limit}")
}

func UserDatedSessionViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateUserDatedSessionView(w http.ResponseWriter, r *http.Request) {
	defer UserDatedSessionViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertUserDatedSessionView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateUserDatedSessionView(w http.ResponseWriter, r *http.Request) {
	defer UserDatedSessionViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertUserDatedSessionView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateUserDatedSessionViewWhere(w http.ResponseWriter, r *http.Request) {
	defer UserDatedSessionViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateUserDatedSessionViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindUserDatedSessionView(w http.ResponseWriter, r *http.Request) {
	defer UserDatedSessionViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindUserDatedSessionView(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindUserDatedSessionView: %#v", m)
	types.MarshalInto(m, w)
}
