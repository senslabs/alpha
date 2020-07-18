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

func OrgSessionEventViewMain(r *mux.Router) {
	r.HandleFunc("/api/org-session-event-views/create", CreateOrgSessionEventView)
	r.HandleFunc("/api/org-session-event-views/batch/create", BatchCreateOrgSessionEventView)
	
	r.HandleFunc("/api/org-session-event-views/update", UpdateOrgSessionEventViewWhere)
	r.HandleFunc("/api/org-session-event-views/find", FindOrgSessionEventView).Queries("limit", "{limit}")
	r.HandleFunc("/api/org-session-event-views/delete", DeleteOrgSessionEventView)
}

func OrgSessionEventViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateOrgSessionEventView(w http.ResponseWriter, r *http.Request) {
	defer OrgSessionEventViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	id := fn.InsertOrgSessionEventView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateOrgSessionEventView(w http.ResponseWriter, r *http.Request) {
	defer OrgSessionEventViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	fn.BatchInsertOrgSessionEventView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateOrgSessionEventViewWhere(w http.ResponseWriter, r *http.Request) {
	defer OrgSessionEventViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateOrgSessionEventViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindOrgSessionEventView(w http.ResponseWriter, r *http.Request) {
	defer OrgSessionEventViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindOrgSessionEventView(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindOrgSessionEventView: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteOrgSessionEventView(w http.ResponseWriter, r *http.Request) {
	defer OrgSessionEventViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteOrgSessionEventView(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteOrgSessionEventView: %d", n)
	types.MarshalInto(n, w)
}