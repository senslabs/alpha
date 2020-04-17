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

func OrgQuarterUsageViewMain(r *mux.Router) {
	r.HandleFunc("/api/org-quarter-usage-views/create", CreateOrgQuarterUsageView)
	r.HandleFunc("/api/org-quarter-usage-views/batch/create", BatchCreateOrgQuarterUsageView)
	
	r.HandleFunc("/api/org-quarter-usage-views/update", UpdateOrgQuarterUsageViewWhere)
	r.HandleFunc("/api/org-quarter-usage-views/find", FindOrgQuarterUsageView).Queries("limit", "{limit}")
}

func OrgQuarterUsageViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateOrgQuarterUsageView(w http.ResponseWriter, r *http.Request) {
	defer OrgQuarterUsageViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertOrgQuarterUsageView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateOrgQuarterUsageView(w http.ResponseWriter, r *http.Request) {
	defer OrgQuarterUsageViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertOrgQuarterUsageView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateOrgQuarterUsageViewWhere(w http.ResponseWriter, r *http.Request) {
	defer OrgQuarterUsageViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateOrgQuarterUsageViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindOrgQuarterUsageView(w http.ResponseWriter, r *http.Request) {
	defer OrgQuarterUsageViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindOrgQuarterUsageView(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindOrgQuarterUsageView: %#v", m)
	types.MarshalInto(m, w)
}
