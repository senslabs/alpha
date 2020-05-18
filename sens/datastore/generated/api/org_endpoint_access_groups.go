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

func OrgEndpointAccessGroupMain(r *mux.Router) {
	r.HandleFunc("/api/org-endpoint-access-groups/create", CreateOrgEndpointAccessGroup)
	r.HandleFunc("/api/org-endpoint-access-groups/batch/create", BatchCreateOrgEndpointAccessGroup)
	
	r.HandleFunc("/api/org-endpoint-access-groups/update", UpdateOrgEndpointAccessGroupWhere)
	r.HandleFunc("/api/org-endpoint-access-groups/find", FindOrgEndpointAccessGroup).Queries("limit", "{limit}")
	r.HandleFunc("/api/org-endpoint-access-groups/delete", DeleteOrgEndpointAccessGroup)
}

func OrgEndpointAccessGroupRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateOrgEndpointAccessGroup(w http.ResponseWriter, r *http.Request) {
	defer OrgEndpointAccessGroupRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertOrgEndpointAccessGroup(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateOrgEndpointAccessGroup(w http.ResponseWriter, r *http.Request) {
	defer OrgEndpointAccessGroupRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertOrgEndpointAccessGroup(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateOrgEndpointAccessGroupWhere(w http.ResponseWriter, r *http.Request) {
	defer OrgEndpointAccessGroupRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateOrgEndpointAccessGroupWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindOrgEndpointAccessGroup(w http.ResponseWriter, r *http.Request) {
	defer OrgEndpointAccessGroupRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindOrgEndpointAccessGroup(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindOrgEndpointAccessGroup: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteOrgEndpointAccessGroup(w http.ResponseWriter, r *http.Request) {
	defer OrgEndpointAccessGroupRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteOrgEndpointAccessGroup(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteOrgEndpointAccessGroup: %d", n)
	types.MarshalInto(n, w)
}