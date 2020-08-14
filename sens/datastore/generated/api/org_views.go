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

func OrgViewMain(r *mux.Router) {
	r.HandleFunc("/api/org-views/create", CreateOrgView)
	r.HandleFunc("/api/org-views/batch/create", BatchCreateOrgView)
	
	r.HandleFunc("/api/org-views/update", UpdateOrgViewWhere)
	r.HandleFunc("/api/org-views/find", FindOrgView).Queries("limit", "{limit}")
	r.HandleFunc("/api/org-views/delete", DeleteOrgView)
}

func OrgViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateOrgView(w http.ResponseWriter, r *http.Request) {
	defer OrgViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	id := fn.InsertOrgView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateOrgView(w http.ResponseWriter, r *http.Request) {
	defer OrgViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	fn.BatchUpsertOrgView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateOrgViewWhere(w http.ResponseWriter, r *http.Request) {
	defer OrgViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateOrgViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindOrgView(w http.ResponseWriter, r *http.Request) {
	defer OrgViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindOrgView(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindOrgView: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteOrgView(w http.ResponseWriter, r *http.Request) {
	defer OrgViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteOrgView(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteOrgView: %d", n)
	types.MarshalInto(n, w)
}