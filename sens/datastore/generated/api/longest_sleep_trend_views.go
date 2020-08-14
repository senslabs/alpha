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

func LongestSleepTrendViewMain(r *mux.Router) {
	r.HandleFunc("/api/longest-sleep-trend-views/create", CreateLongestSleepTrendView)
	r.HandleFunc("/api/longest-sleep-trend-views/batch/create", BatchCreateLongestSleepTrendView)
	
	r.HandleFunc("/api/longest-sleep-trend-views/update", UpdateLongestSleepTrendViewWhere)
	r.HandleFunc("/api/longest-sleep-trend-views/find", FindLongestSleepTrendView).Queries("limit", "{limit}")
	r.HandleFunc("/api/longest-sleep-trend-views/delete", DeleteLongestSleepTrendView)
}

func LongestSleepTrendViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateLongestSleepTrendView(w http.ResponseWriter, r *http.Request) {
	defer LongestSleepTrendViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	id := fn.InsertLongestSleepTrendView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateLongestSleepTrendView(w http.ResponseWriter, r *http.Request) {
	defer LongestSleepTrendViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	fn.BatchUpsertLongestSleepTrendView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateLongestSleepTrendViewWhere(w http.ResponseWriter, r *http.Request) {
	defer LongestSleepTrendViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateLongestSleepTrendViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindLongestSleepTrendView(w http.ResponseWriter, r *http.Request) {
	defer LongestSleepTrendViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindLongestSleepTrendView(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindLongestSleepTrendView: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteLongestSleepTrendView(w http.ResponseWriter, r *http.Request) {
	defer LongestSleepTrendViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteLongestSleepTrendView(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteLongestSleepTrendView: %d", n)
	types.MarshalInto(n, w)
}