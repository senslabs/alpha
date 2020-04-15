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

func UserSummaryViewMain(r *mux.Router) {
	r.HandleFunc("/api/user-summary-views/create", CreateUserSummaryView)
	r.HandleFunc("/api/user-summary-views/batch/create", BatchCreateUserSummaryView)
	
	r.HandleFunc("/api/user-summary-views/update", UpdateUserSummaryViewWhere)
	r.HandleFunc("/api/user-summary-views/find", FindUserSummaryView).Queries("limit", "{limit}")
}

func UserSummaryViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateUserSummaryView(w http.ResponseWriter, r *http.Request) {
	defer UserSummaryViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertUserSummaryView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateUserSummaryView(w http.ResponseWriter, r *http.Request) {
	defer UserSummaryViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertUserSummaryView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateUserSummaryViewWhere(w http.ResponseWriter, r *http.Request) {
	defer UserSummaryViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateUserSummaryViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindUserSummaryView(w http.ResponseWriter, r *http.Request) {
	defer UserSummaryViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindUserSummaryView(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindUserSummaryView: %#v", m)
	types.MarshalInto(m, w)
}
