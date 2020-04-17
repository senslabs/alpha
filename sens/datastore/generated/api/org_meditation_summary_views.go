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

func OrgMeditationSummaryViewMain(r *mux.Router) {
	r.HandleFunc("/api/org-meditation-summary-views/create", CreateOrgMeditationSummaryView)
	r.HandleFunc("/api/org-meditation-summary-views/batch/create", BatchCreateOrgMeditationSummaryView)
	
	r.HandleFunc("/api/org-meditation-summary-views/update", UpdateOrgMeditationSummaryViewWhere)
	r.HandleFunc("/api/org-meditation-summary-views/find", FindOrgMeditationSummaryView).Queries("limit", "{limit}")
}

func OrgMeditationSummaryViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateOrgMeditationSummaryView(w http.ResponseWriter, r *http.Request) {
	defer OrgMeditationSummaryViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertOrgMeditationSummaryView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateOrgMeditationSummaryView(w http.ResponseWriter, r *http.Request) {
	defer OrgMeditationSummaryViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertOrgMeditationSummaryView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateOrgMeditationSummaryViewWhere(w http.ResponseWriter, r *http.Request) {
	defer OrgMeditationSummaryViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateOrgMeditationSummaryViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindOrgMeditationSummaryView(w http.ResponseWriter, r *http.Request) {
	defer OrgMeditationSummaryViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindOrgMeditationSummaryView(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindOrgMeditationSummaryView: %#v", m)
	types.MarshalInto(m, w)
}
