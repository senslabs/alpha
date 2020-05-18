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

func OrgAlertViewMain(r *mux.Router) {
	r.HandleFunc("/api/org-alert-views/create", CreateOrgAlertView)
	r.HandleFunc("/api/org-alert-views/batch/create", BatchCreateOrgAlertView)
	
	r.HandleFunc("/api/org-alert-views/update", UpdateOrgAlertViewWhere)
	r.HandleFunc("/api/org-alert-views/find", FindOrgAlertView).Queries("limit", "{limit}")
	r.HandleFunc("/api/org-alert-views/delete", DeleteOrgAlertView)
}

func OrgAlertViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateOrgAlertView(w http.ResponseWriter, r *http.Request) {
	defer OrgAlertViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertOrgAlertView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateOrgAlertView(w http.ResponseWriter, r *http.Request) {
	defer OrgAlertViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertOrgAlertView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateOrgAlertViewWhere(w http.ResponseWriter, r *http.Request) {
	defer OrgAlertViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateOrgAlertViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindOrgAlertView(w http.ResponseWriter, r *http.Request) {
	defer OrgAlertViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindOrgAlertView(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindOrgAlertView: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteOrgAlertView(w http.ResponseWriter, r *http.Request) {
	defer OrgAlertViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteOrgAlertView(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteOrgAlertView: %d", n)
	types.MarshalInto(n, w)
}