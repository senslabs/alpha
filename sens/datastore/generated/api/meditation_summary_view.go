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

func MeditationSummaryViewMain(r *mux.Router) {
	r.HandleFunc("/api/meditation-summary-view/create", CreateMeditationSummaryView)
	r.HandleFunc("/api/meditation-summary-view/batch/create", BatchCreateMeditationSummaryView)
	
	r.HandleFunc("/api/meditation-summary-view/update", UpdateMeditationSummaryViewWhere)
	r.HandleFunc("/api/meditation-summary-view/find", FindMeditationSummaryView).Queries("limit", "{limit}")
	r.HandleFunc("/api/meditation-summary-view/delete", DeleteMeditationSummaryView)
}

func MeditationSummaryViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateMeditationSummaryView(w http.ResponseWriter, r *http.Request) {
	defer MeditationSummaryViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertMeditationSummaryView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateMeditationSummaryView(w http.ResponseWriter, r *http.Request) {
	defer MeditationSummaryViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertMeditationSummaryView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateMeditationSummaryViewWhere(w http.ResponseWriter, r *http.Request) {
	defer MeditationSummaryViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateMeditationSummaryViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindMeditationSummaryView(w http.ResponseWriter, r *http.Request) {
	defer MeditationSummaryViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindMeditationSummaryView(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindMeditationSummaryView: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteMeditationSummaryView(w http.ResponseWriter, r *http.Request) {
	defer MeditationSummaryViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteMeditationSummaryView(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteMeditationSummaryView: %d", n)
	types.MarshalInto(n, w)
}