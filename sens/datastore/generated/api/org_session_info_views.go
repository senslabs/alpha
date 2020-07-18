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

func OrgSessionInfoViewMain(r *mux.Router) {
	r.HandleFunc("/api/org-session-info-views/create", CreateOrgSessionInfoView)
	r.HandleFunc("/api/org-session-info-views/batch/create", BatchCreateOrgSessionInfoView)
	
	r.HandleFunc("/api/org-session-info-views/update", UpdateOrgSessionInfoViewWhere)
	r.HandleFunc("/api/org-session-info-views/find", FindOrgSessionInfoView).Queries("limit", "{limit}")
	r.HandleFunc("/api/org-session-info-views/delete", DeleteOrgSessionInfoView)
}

func OrgSessionInfoViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateOrgSessionInfoView(w http.ResponseWriter, r *http.Request) {
	defer OrgSessionInfoViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	id := fn.InsertOrgSessionInfoView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateOrgSessionInfoView(w http.ResponseWriter, r *http.Request) {
	defer OrgSessionInfoViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	fn.BatchInsertOrgSessionInfoView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateOrgSessionInfoViewWhere(w http.ResponseWriter, r *http.Request) {
	defer OrgSessionInfoViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateOrgSessionInfoViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindOrgSessionInfoView(w http.ResponseWriter, r *http.Request) {
	defer OrgSessionInfoViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindOrgSessionInfoView(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindOrgSessionInfoView: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteOrgSessionInfoView(w http.ResponseWriter, r *http.Request) {
	defer OrgSessionInfoViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteOrgSessionInfoView(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteOrgSessionInfoView: %d", n)
	types.MarshalInto(n, w)
}