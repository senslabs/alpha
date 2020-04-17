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

func OrgSleepViewMain(r *mux.Router) {
	r.HandleFunc("/api/org-sleep-views/create", CreateOrgSleepView)
	r.HandleFunc("/api/org-sleep-views/batch/create", BatchCreateOrgSleepView)
	
	r.HandleFunc("/api/org-sleep-views/update", UpdateOrgSleepViewWhere)
	r.HandleFunc("/api/org-sleep-views/find", FindOrgSleepView).Queries("limit", "{limit}")
}

func OrgSleepViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateOrgSleepView(w http.ResponseWriter, r *http.Request) {
	defer OrgSleepViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertOrgSleepView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateOrgSleepView(w http.ResponseWriter, r *http.Request) {
	defer OrgSleepViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertOrgSleepView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateOrgSleepViewWhere(w http.ResponseWriter, r *http.Request) {
	defer OrgSleepViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateOrgSleepViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindOrgSleepView(w http.ResponseWriter, r *http.Request) {
	defer OrgSleepViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindOrgSleepView(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindOrgSleepView: %#v", m)
	types.MarshalInto(m, w)
}
