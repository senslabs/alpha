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

func RecorderMain(r *mux.Router) {
	r.HandleFunc("/api/recorders/create", CreateRecorder)
	r.HandleFunc("/api/recorders/batch/create", BatchCreateRecorder)
	
	r.HandleFunc("/api/recorders/update", UpdateRecorderWhere)
	r.HandleFunc("/api/recorders/find", FindRecorder).Queries("limit", "{limit}")
	r.HandleFunc("/api/recorders/delete", DeleteRecorder)
}

func RecorderRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateRecorder(w http.ResponseWriter, r *http.Request) {
	defer RecorderRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertRecorder(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateRecorder(w http.ResponseWriter, r *http.Request) {
	defer RecorderRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertRecorder(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateRecorderWhere(w http.ResponseWriter, r *http.Request) {
	defer RecorderRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateRecorderWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindRecorder(w http.ResponseWriter, r *http.Request) {
	defer RecorderRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindRecorder(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindRecorder: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteRecorder(w http.ResponseWriter, r *http.Request) {
	defer RecorderRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteRecorder(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteRecorder: %d", n)
	types.MarshalInto(n, w)
}