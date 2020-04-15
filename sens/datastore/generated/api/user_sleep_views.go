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

func UserSleepViewMain(r *mux.Router) {
	r.HandleFunc("/api/user-sleep-views/create", CreateUserSleepView)
	r.HandleFunc("/api/user-sleep-views/batch/create", BatchCreateUserSleepView)
	
	r.HandleFunc("/api/user-sleep-views/update", UpdateUserSleepViewWhere)
	r.HandleFunc("/api/user-sleep-views/find", FindUserSleepView).Queries("limit", "{limit}")
}

func UserSleepViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateUserSleepView(w http.ResponseWriter, r *http.Request) {
	defer UserSleepViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertUserSleepView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateUserSleepView(w http.ResponseWriter, r *http.Request) {
	defer UserSleepViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertUserSleepView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateUserSleepViewWhere(w http.ResponseWriter, r *http.Request) {
	defer UserSleepViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateUserSleepViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindUserSleepView(w http.ResponseWriter, r *http.Request) {
	defer UserSleepViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindUserSleepView(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindUserSleepView: %#v", m)
	types.MarshalInto(m, w)
}
