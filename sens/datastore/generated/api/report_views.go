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

func ReportViewMain(r *mux.Router) {
	r.HandleFunc("/api/report-views/create", CreateReportView)
	r.HandleFunc("/api/report-views/batch/create", BatchCreateReportView)
	
	r.HandleFunc("/api/report-views/update", UpdateReportViewWhere)
	r.HandleFunc("/api/report-views/find", FindReportView).Queries("limit", "{limit}")
}

func ReportViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateReportView(w http.ResponseWriter, r *http.Request) {
	defer ReportViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertReportView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateReportView(w http.ResponseWriter, r *http.Request) {
	defer ReportViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertReportView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateReportViewWhere(w http.ResponseWriter, r *http.Request) {
	defer ReportViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateReportViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindReportView(w http.ResponseWriter, r *http.Request) {
	defer ReportViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindReportView(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindReportView: %#v", m)
	types.MarshalInto(m, w)
}
