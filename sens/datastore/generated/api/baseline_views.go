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

func BaselineViewMain(r *mux.Router) {
	r.HandleFunc("/api/baseline-views/create", CreateBaselineView)
	r.HandleFunc("/api/baseline-views/batch/create", BatchCreateBaselineView)
	
	r.HandleFunc("/api/baseline-views/update", UpdateBaselineViewWhere)
	r.HandleFunc("/api/baseline-views/find", FindBaselineView).Queries("limit", "{limit}")
	r.HandleFunc("/api/baseline-views/delete", DeleteBaselineView)
}

func BaselineViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateBaselineView(w http.ResponseWriter, r *http.Request) {
	defer BaselineViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertBaselineView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateBaselineView(w http.ResponseWriter, r *http.Request) {
	defer BaselineViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertBaselineView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateBaselineViewWhere(w http.ResponseWriter, r *http.Request) {
	defer BaselineViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateBaselineViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindBaselineView(w http.ResponseWriter, r *http.Request) {
	defer BaselineViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindBaselineView(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindBaselineView: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteBaselineView(w http.ResponseWriter, r *http.Request) {
	defer BaselineViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteBaselineView(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteBaselineView: %d", n)
	types.MarshalInto(n, w)
}