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

func ReportMain(r *mux.Router) {
	r.HandleFunc("/api/reports/create", CreateReport)
	r.HandleFunc("/api/reports/batch/create", BatchCreateReport)
	
	r.HandleFunc("/api/reports/{id}/update", UpdateReport)
	r.HandleFunc("/api/reports/{id}/get", GetReport)
    
	r.HandleFunc("/api/reports/update", UpdateReportWhere)
	r.HandleFunc("/api/reports/find", FindReport).Queries("limit", "{limit}")
	r.HandleFunc("/api/reports/delete", DeleteReport)
}

func ReportRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateReport(w http.ResponseWriter, r *http.Request) {
	defer ReportRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	id := fn.InsertReport(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateReport(w http.ResponseWriter, r *http.Request) {
	defer ReportRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	fn.BatchUpsertReport(data)
	w.WriteHeader(http.StatusOK)
}


func UpdateReport(w http.ResponseWriter, r *http.Request) {
	defer ReportRecovery(w)
	vars := mux.Vars(r)
	id := vars["id"]
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateReport(id, data)
	w.WriteHeader(http.StatusOK)
}

func GetReport(w http.ResponseWriter, r *http.Request) {
	defer ReportRecovery(w)
	vars := mux.Vars(r)
	id := vars["id"]
	m := fn.SelectReport(id)
	types.MarshalInto(m, w)
}


func UpdateReportWhere(w http.ResponseWriter, r *http.Request) {
	defer ReportRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateReportWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindReport(w http.ResponseWriter, r *http.Request) {
	defer ReportRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindReport(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindReport: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteReport(w http.ResponseWriter, r *http.Request) {
	defer ReportRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteReport(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteReport: %d", n)
	types.MarshalInto(n, w)
}