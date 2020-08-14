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

func OrgActivityViewMain(r *mux.Router) {
	r.HandleFunc("/api/org-activity-views/create", CreateOrgActivityView)
	r.HandleFunc("/api/org-activity-views/batch/create", BatchCreateOrgActivityView)
	
	r.HandleFunc("/api/org-activity-views/update", UpdateOrgActivityViewWhere)
	r.HandleFunc("/api/org-activity-views/find", FindOrgActivityView).Queries("limit", "{limit}")
	r.HandleFunc("/api/org-activity-views/delete", DeleteOrgActivityView)
}

func OrgActivityViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateOrgActivityView(w http.ResponseWriter, r *http.Request) {
	defer OrgActivityViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	id := fn.InsertOrgActivityView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateOrgActivityView(w http.ResponseWriter, r *http.Request) {
	defer OrgActivityViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	fn.BatchUpsertOrgActivityView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateOrgActivityViewWhere(w http.ResponseWriter, r *http.Request) {
	defer OrgActivityViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateOrgActivityViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindOrgActivityView(w http.ResponseWriter, r *http.Request) {
	defer OrgActivityViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindOrgActivityView(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindOrgActivityView: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteOrgActivityView(w http.ResponseWriter, r *http.Request) {
	defer OrgActivityViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteOrgActivityView(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteOrgActivityView: %d", n)
	types.MarshalInto(n, w)
}