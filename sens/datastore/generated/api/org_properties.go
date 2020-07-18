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

func OrgPropertieMain(r *mux.Router) {
	r.HandleFunc("/api/org-properties/create", CreateOrgPropertie)
	r.HandleFunc("/api/org-properties/batch/create", BatchCreateOrgPropertie)
	
	r.HandleFunc("/api/org-properties/update", UpdateOrgPropertieWhere)
	r.HandleFunc("/api/org-properties/find", FindOrgPropertie).Queries("limit", "{limit}")
	r.HandleFunc("/api/org-properties/delete", DeleteOrgPropertie)
}

func OrgPropertieRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateOrgPropertie(w http.ResponseWriter, r *http.Request) {
	defer OrgPropertieRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	id := fn.InsertOrgPropertie(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateOrgPropertie(w http.ResponseWriter, r *http.Request) {
	defer OrgPropertieRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	fn.BatchInsertOrgPropertie(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateOrgPropertieWhere(w http.ResponseWriter, r *http.Request) {
	defer OrgPropertieRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateOrgPropertieWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindOrgPropertie(w http.ResponseWriter, r *http.Request) {
	defer OrgPropertieRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindOrgPropertie(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindOrgPropertie: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteOrgPropertie(w http.ResponseWriter, r *http.Request) {
	defer OrgPropertieRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteOrgPropertie(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteOrgPropertie: %d", n)
	types.MarshalInto(n, w)
}