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

func OrgSessionDetailViewMain(r *mux.Router) {
	r.HandleFunc("/api/org-session-detail-views/create", CreateOrgSessionDetailView)
	r.HandleFunc("/api/org-session-detail-views/batch/create", BatchCreateOrgSessionDetailView)
	
	r.HandleFunc("/api/org-session-detail-views/update", UpdateOrgSessionDetailViewWhere)
	r.HandleFunc("/api/org-session-detail-views/find", FindOrgSessionDetailView).Queries("limit", "{limit}")
	r.HandleFunc("/api/org-session-detail-views/delete", DeleteOrgSessionDetailView)
}

func OrgSessionDetailViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateOrgSessionDetailView(w http.ResponseWriter, r *http.Request) {
	defer OrgSessionDetailViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertOrgSessionDetailView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateOrgSessionDetailView(w http.ResponseWriter, r *http.Request) {
	defer OrgSessionDetailViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertOrgSessionDetailView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateOrgSessionDetailViewWhere(w http.ResponseWriter, r *http.Request) {
	defer OrgSessionDetailViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateOrgSessionDetailViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindOrgSessionDetailView(w http.ResponseWriter, r *http.Request) {
	defer OrgSessionDetailViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindOrgSessionDetailView(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindOrgSessionDetailView: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteOrgSessionDetailView(w http.ResponseWriter, r *http.Request) {
	defer OrgSessionDetailViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteOrgSessionDetailView(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteOrgSessionDetailView: %d", n)
	types.MarshalInto(n, w)
}