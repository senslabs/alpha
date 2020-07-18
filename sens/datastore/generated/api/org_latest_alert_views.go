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

func OrgLatestAlertViewMain(r *mux.Router) {
	r.HandleFunc("/api/org-latest-alert-views/create", CreateOrgLatestAlertView)
	r.HandleFunc("/api/org-latest-alert-views/batch/create", BatchCreateOrgLatestAlertView)
	
	r.HandleFunc("/api/org-latest-alert-views/update", UpdateOrgLatestAlertViewWhere)
	r.HandleFunc("/api/org-latest-alert-views/find", FindOrgLatestAlertView).Queries("limit", "{limit}")
	r.HandleFunc("/api/org-latest-alert-views/delete", DeleteOrgLatestAlertView)
}

func OrgLatestAlertViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateOrgLatestAlertView(w http.ResponseWriter, r *http.Request) {
	defer OrgLatestAlertViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	id := fn.InsertOrgLatestAlertView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateOrgLatestAlertView(w http.ResponseWriter, r *http.Request) {
	defer OrgLatestAlertViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	fn.BatchInsertOrgLatestAlertView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateOrgLatestAlertViewWhere(w http.ResponseWriter, r *http.Request) {
	defer OrgLatestAlertViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateOrgLatestAlertViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindOrgLatestAlertView(w http.ResponseWriter, r *http.Request) {
	defer OrgLatestAlertViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindOrgLatestAlertView(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindOrgLatestAlertView: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteOrgLatestAlertView(w http.ResponseWriter, r *http.Request) {
	defer OrgLatestAlertViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteOrgLatestAlertView(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteOrgLatestAlertView: %d", n)
	types.MarshalInto(n, w)
}