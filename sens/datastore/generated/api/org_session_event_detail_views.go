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

func OrgSessionEventDetailViewMain(r *mux.Router) {
	r.HandleFunc("/api/org-session-event-detail-views/create", CreateOrgSessionEventDetailView)
	r.HandleFunc("/api/org-session-event-detail-views/batch/create", BatchCreateOrgSessionEventDetailView)
	
	r.HandleFunc("/api/org-session-event-detail-views/update", UpdateOrgSessionEventDetailViewWhere)
	r.HandleFunc("/api/org-session-event-detail-views/find", FindOrgSessionEventDetailView).Queries("limit", "{limit}")
	r.HandleFunc("/api/org-session-event-detail-views/delete", DeleteOrgSessionEventDetailView)
}

func OrgSessionEventDetailViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateOrgSessionEventDetailView(w http.ResponseWriter, r *http.Request) {
	defer OrgSessionEventDetailViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	id := fn.InsertOrgSessionEventDetailView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateOrgSessionEventDetailView(w http.ResponseWriter, r *http.Request) {
	defer OrgSessionEventDetailViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	fn.BatchInsertOrgSessionEventDetailView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateOrgSessionEventDetailViewWhere(w http.ResponseWriter, r *http.Request) {
	defer OrgSessionEventDetailViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateOrgSessionEventDetailViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindOrgSessionEventDetailView(w http.ResponseWriter, r *http.Request) {
	defer OrgSessionEventDetailViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindOrgSessionEventDetailView(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindOrgSessionEventDetailView: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteOrgSessionEventDetailView(w http.ResponseWriter, r *http.Request) {
	defer OrgSessionEventDetailViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteOrgSessionEventDetailView(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteOrgSessionEventDetailView: %d", n)
	types.MarshalInto(n, w)
}