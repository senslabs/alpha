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

func SleepSummarieMain(r *mux.Router) {
	r.HandleFunc("/api/sleep-summaries/create", CreateSleepSummarie)
	r.HandleFunc("/api/sleep-summaries/batch/create", BatchCreateSleepSummarie)
	
	r.HandleFunc("/api/sleep-summaries/update", UpdateSleepSummarieWhere)
	r.HandleFunc("/api/sleep-summaries/find", FindSleepSummarie).Queries("limit", "{limit}")
}

func SleepSummarieRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateSleepSummarie(w http.ResponseWriter, r *http.Request) {
	defer SleepSummarieRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertSleepSummarie(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateSleepSummarie(w http.ResponseWriter, r *http.Request) {
	defer SleepSummarieRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertSleepSummarie(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateSleepSummarieWhere(w http.ResponseWriter, r *http.Request) {
	defer SleepSummarieRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateSleepSummarieWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindSleepSummarie(w http.ResponseWriter, r *http.Request) {
	defer SleepSummarieRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindSleepSummarie(or, and, in, span, limit, column, order)
	types.MarshalInto(m, w)
}
