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

func OrgSleepSummaryViewMain(r *mux.Router) {
	r.HandleFunc("/api/org-sleep-summary-views/create", CreateOrgSleepSummaryView)
	r.HandleFunc("/api/org-sleep-summary-views/batch/create", BatchCreateOrgSleepSummaryView)
	
	r.HandleFunc("/api/org-sleep-summary-views/update", UpdateOrgSleepSummaryViewWhere)
	r.HandleFunc("/api/org-sleep-summary-views/find", FindOrgSleepSummaryView).Queries("limit", "{limit}")
}

func OrgSleepSummaryViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateOrgSleepSummaryView(w http.ResponseWriter, r *http.Request) {
	defer OrgSleepSummaryViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertOrgSleepSummaryView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateOrgSleepSummaryView(w http.ResponseWriter, r *http.Request) {
	defer OrgSleepSummaryViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertOrgSleepSummaryView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateOrgSleepSummaryViewWhere(w http.ResponseWriter, r *http.Request) {
	defer OrgSleepSummaryViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateOrgSleepSummaryViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindOrgSleepSummaryView(w http.ResponseWriter, r *http.Request) {
	defer OrgSleepSummaryViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindOrgSleepSummaryView(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindOrgSleepSummaryView: %#v", m)
	types.MarshalInto(m, w)
}
