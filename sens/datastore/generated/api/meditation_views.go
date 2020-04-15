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

func MeditationViewMain(r *mux.Router) {
	r.HandleFunc("/api/meditation-views/create", CreateMeditationView)
	r.HandleFunc("/api/meditation-views/batch/create", BatchCreateMeditationView)
	
	r.HandleFunc("/api/meditation-views/update", UpdateMeditationViewWhere)
	r.HandleFunc("/api/meditation-views/find", FindMeditationView).Queries("limit", "{limit}")
}

func MeditationViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateMeditationView(w http.ResponseWriter, r *http.Request) {
	defer MeditationViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertMeditationView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateMeditationView(w http.ResponseWriter, r *http.Request) {
	defer MeditationViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertMeditationView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateMeditationViewWhere(w http.ResponseWriter, r *http.Request) {
	defer MeditationViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateMeditationViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindMeditationView(w http.ResponseWriter, r *http.Request) {
	defer MeditationViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindMeditationView(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindMeditationView: %#v", m)
	types.MarshalInto(m, w)
}
