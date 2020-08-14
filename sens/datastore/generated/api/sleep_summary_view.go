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

func SleepSummaryViewMain(r *mux.Router) {
	r.HandleFunc("/api/sleep-summary-view/create", CreateSleepSummaryView)
	r.HandleFunc("/api/sleep-summary-view/batch/create", BatchCreateSleepSummaryView)
	
	r.HandleFunc("/api/sleep-summary-view/update", UpdateSleepSummaryViewWhere)
	r.HandleFunc("/api/sleep-summary-view/find", FindSleepSummaryView).Queries("limit", "{limit}")
	r.HandleFunc("/api/sleep-summary-view/delete", DeleteSleepSummaryView)
}

func SleepSummaryViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateSleepSummaryView(w http.ResponseWriter, r *http.Request) {
	defer SleepSummaryViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	id := fn.InsertSleepSummaryView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateSleepSummaryView(w http.ResponseWriter, r *http.Request) {
	defer SleepSummaryViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	fn.BatchUpsertSleepSummaryView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateSleepSummaryViewWhere(w http.ResponseWriter, r *http.Request) {
	defer SleepSummaryViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateSleepSummaryViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindSleepSummaryView(w http.ResponseWriter, r *http.Request) {
	defer SleepSummaryViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindSleepSummaryView(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindSleepSummaryView: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteSleepSummaryView(w http.ResponseWriter, r *http.Request) {
	defer SleepSummaryViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteSleepSummaryView(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteSleepSummaryView: %d", n)
	types.MarshalInto(n, w)
}