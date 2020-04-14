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

func UserSessionViewMain(r *mux.Router) {
	r.HandleFunc("/api/user-session-views/create", CreateUserSessionView)
	r.HandleFunc("/api/user-session-views/batch/create", BatchCreateUserSessionView)
	
	r.HandleFunc("/api/user-session-views/update", UpdateUserSessionViewWhere)
	r.HandleFunc("/api/user-session-views/find", FindUserSessionView).Queries("limit", "{limit}")
}

func UserSessionViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateUserSessionView(w http.ResponseWriter, r *http.Request) {
	defer UserSessionViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertUserSessionView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateUserSessionView(w http.ResponseWriter, r *http.Request) {
	defer UserSessionViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertUserSessionView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateUserSessionViewWhere(w http.ResponseWriter, r *http.Request) {
	defer UserSessionViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateUserSessionViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindUserSessionView(w http.ResponseWriter, r *http.Request) {
	defer UserSessionViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindUserSessionView(or, and, in, span, limit, column, order)
	types.MarshalInto(m, w)
}
