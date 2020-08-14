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

func LatestSleepSummaryViewMain(r *mux.Router) {
	r.HandleFunc("/api/latest-sleep-summary-view/create", CreateLatestSleepSummaryView)
	r.HandleFunc("/api/latest-sleep-summary-view/batch/create", BatchCreateLatestSleepSummaryView)
	
	r.HandleFunc("/api/latest-sleep-summary-view/update", UpdateLatestSleepSummaryViewWhere)
	r.HandleFunc("/api/latest-sleep-summary-view/find", FindLatestSleepSummaryView).Queries("limit", "{limit}")
	r.HandleFunc("/api/latest-sleep-summary-view/delete", DeleteLatestSleepSummaryView)
}

func LatestSleepSummaryViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateLatestSleepSummaryView(w http.ResponseWriter, r *http.Request) {
	defer LatestSleepSummaryViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	id := fn.InsertLatestSleepSummaryView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateLatestSleepSummaryView(w http.ResponseWriter, r *http.Request) {
	defer LatestSleepSummaryViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	fn.BatchUpsertLatestSleepSummaryView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateLatestSleepSummaryViewWhere(w http.ResponseWriter, r *http.Request) {
	defer LatestSleepSummaryViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateLatestSleepSummaryViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindLatestSleepSummaryView(w http.ResponseWriter, r *http.Request) {
	defer LatestSleepSummaryViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindLatestSleepSummaryView(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindLatestSleepSummaryView: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteLatestSleepSummaryView(w http.ResponseWriter, r *http.Request) {
	defer LatestSleepSummaryViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteLatestSleepSummaryView(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteLatestSleepSummaryView: %d", n)
	types.MarshalInto(n, w)
}