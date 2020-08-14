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

func RecorderViewMain(r *mux.Router) {
	r.HandleFunc("/api/recorder-views/create", CreateRecorderView)
	r.HandleFunc("/api/recorder-views/batch/create", BatchCreateRecorderView)
	
	r.HandleFunc("/api/recorder-views/update", UpdateRecorderViewWhere)
	r.HandleFunc("/api/recorder-views/find", FindRecorderView).Queries("limit", "{limit}")
	r.HandleFunc("/api/recorder-views/delete", DeleteRecorderView)
}

func RecorderViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateRecorderView(w http.ResponseWriter, r *http.Request) {
	defer RecorderViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	id := fn.InsertRecorderView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateRecorderView(w http.ResponseWriter, r *http.Request) {
	defer RecorderViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	fn.BatchUpsertRecorderView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateRecorderViewWhere(w http.ResponseWriter, r *http.Request) {
	defer RecorderViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateRecorderViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindRecorderView(w http.ResponseWriter, r *http.Request) {
	defer RecorderViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindRecorderView(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindRecorderView: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteRecorderView(w http.ResponseWriter, r *http.Request) {
	defer RecorderViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteRecorderView(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteRecorderView: %d", n)
	types.MarshalInto(n, w)
}