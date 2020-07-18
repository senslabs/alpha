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

func OrgSessionRecordViewMain(r *mux.Router) {
	r.HandleFunc("/api/org-session-record-views/create", CreateOrgSessionRecordView)
	r.HandleFunc("/api/org-session-record-views/batch/create", BatchCreateOrgSessionRecordView)
	
	r.HandleFunc("/api/org-session-record-views/update", UpdateOrgSessionRecordViewWhere)
	r.HandleFunc("/api/org-session-record-views/find", FindOrgSessionRecordView).Queries("limit", "{limit}")
	r.HandleFunc("/api/org-session-record-views/delete", DeleteOrgSessionRecordView)
}

func OrgSessionRecordViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateOrgSessionRecordView(w http.ResponseWriter, r *http.Request) {
	defer OrgSessionRecordViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	id := fn.InsertOrgSessionRecordView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateOrgSessionRecordView(w http.ResponseWriter, r *http.Request) {
	defer OrgSessionRecordViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	fn.BatchInsertOrgSessionRecordView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateOrgSessionRecordViewWhere(w http.ResponseWriter, r *http.Request) {
	defer OrgSessionRecordViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateOrgSessionRecordViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindOrgSessionRecordView(w http.ResponseWriter, r *http.Request) {
	defer OrgSessionRecordViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindOrgSessionRecordView(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindOrgSessionRecordView: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteOrgSessionRecordView(w http.ResponseWriter, r *http.Request) {
	defer OrgSessionRecordViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteOrgSessionRecordView(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteOrgSessionRecordView: %d", n)
	types.MarshalInto(n, w)
}