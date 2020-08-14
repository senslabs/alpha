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

func OrgActivitySummaryViewMain(r *mux.Router) {
	r.HandleFunc("/api/org-activity-summary-views/create", CreateOrgActivitySummaryView)
	r.HandleFunc("/api/org-activity-summary-views/batch/create", BatchCreateOrgActivitySummaryView)
	
	r.HandleFunc("/api/org-activity-summary-views/update", UpdateOrgActivitySummaryViewWhere)
	r.HandleFunc("/api/org-activity-summary-views/find", FindOrgActivitySummaryView).Queries("limit", "{limit}")
	r.HandleFunc("/api/org-activity-summary-views/delete", DeleteOrgActivitySummaryView)
}

func OrgActivitySummaryViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateOrgActivitySummaryView(w http.ResponseWriter, r *http.Request) {
	defer OrgActivitySummaryViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	id := fn.InsertOrgActivitySummaryView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateOrgActivitySummaryView(w http.ResponseWriter, r *http.Request) {
	defer OrgActivitySummaryViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	fn.BatchUpsertOrgActivitySummaryView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateOrgActivitySummaryViewWhere(w http.ResponseWriter, r *http.Request) {
	defer OrgActivitySummaryViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateOrgActivitySummaryViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindOrgActivitySummaryView(w http.ResponseWriter, r *http.Request) {
	defer OrgActivitySummaryViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindOrgActivitySummaryView(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindOrgActivitySummaryView: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteOrgActivitySummaryView(w http.ResponseWriter, r *http.Request) {
	defer OrgActivitySummaryViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteOrgActivitySummaryView(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteOrgActivitySummaryView: %d", n)
	types.MarshalInto(n, w)
}