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

func SleepViewMain(r *mux.Router) {
	r.HandleFunc("/api/sleep-views/create", CreateSleepView)
	r.HandleFunc("/api/sleep-views/batch/create", BatchCreateSleepView)
	
	r.HandleFunc("/api/sleep-views/update", UpdateSleepViewWhere)
	r.HandleFunc("/api/sleep-views/find", FindSleepView).Queries("limit", "{limit}")
}

func SleepViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateSleepView(w http.ResponseWriter, r *http.Request) {
	defer SleepViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertSleepView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateSleepView(w http.ResponseWriter, r *http.Request) {
	defer SleepViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertSleepView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateSleepViewWhere(w http.ResponseWriter, r *http.Request) {
	defer SleepViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateSleepViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindSleepView(w http.ResponseWriter, r *http.Request) {
	defer SleepViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindSleepView(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindSleepView: %#v", m)
	types.MarshalInto(m, w)
}
