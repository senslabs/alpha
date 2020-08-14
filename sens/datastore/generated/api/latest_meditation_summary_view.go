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

func LatestMeditationSummaryViewMain(r *mux.Router) {
	r.HandleFunc("/api/latest-meditation-summary-view/create", CreateLatestMeditationSummaryView)
	r.HandleFunc("/api/latest-meditation-summary-view/batch/create", BatchCreateLatestMeditationSummaryView)
	
	r.HandleFunc("/api/latest-meditation-summary-view/update", UpdateLatestMeditationSummaryViewWhere)
	r.HandleFunc("/api/latest-meditation-summary-view/find", FindLatestMeditationSummaryView).Queries("limit", "{limit}")
	r.HandleFunc("/api/latest-meditation-summary-view/delete", DeleteLatestMeditationSummaryView)
}

func LatestMeditationSummaryViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateLatestMeditationSummaryView(w http.ResponseWriter, r *http.Request) {
	defer LatestMeditationSummaryViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	id := fn.InsertLatestMeditationSummaryView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateLatestMeditationSummaryView(w http.ResponseWriter, r *http.Request) {
	defer LatestMeditationSummaryViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	fn.BatchUpsertLatestMeditationSummaryView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateLatestMeditationSummaryViewWhere(w http.ResponseWriter, r *http.Request) {
	defer LatestMeditationSummaryViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateLatestMeditationSummaryViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindLatestMeditationSummaryView(w http.ResponseWriter, r *http.Request) {
	defer LatestMeditationSummaryViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindLatestMeditationSummaryView(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindLatestMeditationSummaryView: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteLatestMeditationSummaryView(w http.ResponseWriter, r *http.Request) {
	defer LatestMeditationSummaryViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteLatestMeditationSummaryView(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteLatestMeditationSummaryView: %d", n)
	types.MarshalInto(n, w)
}