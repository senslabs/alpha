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

func UserAlertViewMain(r *mux.Router) {
	r.HandleFunc("/api/user-alert-views/create", CreateUserAlertView)
	r.HandleFunc("/api/user-alert-views/batch/create", BatchCreateUserAlertView)
	
	r.HandleFunc("/api/user-alert-views/update", UpdateUserAlertViewWhere)
	r.HandleFunc("/api/user-alert-views/find", FindUserAlertView).Queries("limit", "{limit}")
}

func UserAlertViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateUserAlertView(w http.ResponseWriter, r *http.Request) {
	defer UserAlertViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertUserAlertView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateUserAlertView(w http.ResponseWriter, r *http.Request) {
	defer UserAlertViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertUserAlertView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateUserAlertViewWhere(w http.ResponseWriter, r *http.Request) {
	defer UserAlertViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateUserAlertViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindUserAlertView(w http.ResponseWriter, r *http.Request) {
	defer UserAlertViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindUserAlertView(or, and, in, span, limit, column, order)
	types.MarshalInto(m, w)
}
