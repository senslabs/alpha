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

func OrgDetailViewMain(r *mux.Router) {
	r.HandleFunc("/api/org-detail-views/create", CreateOrgDetailView)
	r.HandleFunc("/api/org-detail-views/batch/create", BatchCreateOrgDetailView)
	
	r.HandleFunc("/api/org-detail-views/update", UpdateOrgDetailViewWhere)
	r.HandleFunc("/api/org-detail-views/find", FindOrgDetailView).Queries("limit", "{limit}")
}

func OrgDetailViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateOrgDetailView(w http.ResponseWriter, r *http.Request) {
	defer OrgDetailViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertOrgDetailView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateOrgDetailView(w http.ResponseWriter, r *http.Request) {
	defer OrgDetailViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertOrgDetailView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateOrgDetailViewWhere(w http.ResponseWriter, r *http.Request) {
	defer OrgDetailViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateOrgDetailViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindOrgDetailView(w http.ResponseWriter, r *http.Request) {
	defer OrgDetailViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindOrgDetailView(or, and, in, span, limit, column, order)
	types.MarshalInto(m, w)
}
