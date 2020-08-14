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

func OrgMain(r *mux.Router) {
	r.HandleFunc("/api/orgs/create", CreateOrg)
	r.HandleFunc("/api/orgs/batch/create", BatchCreateOrg)
	
	r.HandleFunc("/api/orgs/{id}/update", UpdateOrg)
	r.HandleFunc("/api/orgs/{id}/get", GetOrg)
    
	r.HandleFunc("/api/orgs/update", UpdateOrgWhere)
	r.HandleFunc("/api/orgs/find", FindOrg).Queries("limit", "{limit}")
	r.HandleFunc("/api/orgs/delete", DeleteOrg)
}

func OrgRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateOrg(w http.ResponseWriter, r *http.Request) {
	defer OrgRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	id := fn.InsertOrg(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateOrg(w http.ResponseWriter, r *http.Request) {
	defer OrgRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	fn.BatchUpsertOrg(data)
	w.WriteHeader(http.StatusOK)
}


func UpdateOrg(w http.ResponseWriter, r *http.Request) {
	defer OrgRecovery(w)
	vars := mux.Vars(r)
	id := vars["id"]
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateOrg(id, data)
	w.WriteHeader(http.StatusOK)
}

func GetOrg(w http.ResponseWriter, r *http.Request) {
	defer OrgRecovery(w)
	vars := mux.Vars(r)
	id := vars["id"]
	m := fn.SelectOrg(id)
	types.MarshalInto(m, w)
}


func UpdateOrgWhere(w http.ResponseWriter, r *http.Request) {
	defer OrgRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateOrgWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindOrg(w http.ResponseWriter, r *http.Request) {
	defer OrgRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindOrg(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindOrg: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteOrg(w http.ResponseWriter, r *http.Request) {
	defer OrgRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteOrg(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteOrg: %d", n)
	types.MarshalInto(n, w)
}