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

func OrgSessionViewMain(r *mux.Router) {
	r.HandleFunc("/api/org-session-views/create", CreateOrgSessionView)
	r.HandleFunc("/api/org-session-views/batch/create", BatchCreateOrgSessionView)
	
	r.HandleFunc("/api/org-session-views/update", UpdateOrgSessionViewWhere)
	r.HandleFunc("/api/org-session-views/find", FindOrgSessionView).Queries("limit", "{limit}")
}

func OrgSessionViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateOrgSessionView(w http.ResponseWriter, r *http.Request) {
	defer OrgSessionViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertOrgSessionView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateOrgSessionView(w http.ResponseWriter, r *http.Request) {
	defer OrgSessionViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertOrgSessionView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateOrgSessionViewWhere(w http.ResponseWriter, r *http.Request) {
	defer OrgSessionViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateOrgSessionViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindOrgSessionView(w http.ResponseWriter, r *http.Request) {
	defer OrgSessionViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindOrgSessionView(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindOrgSessionView: %#v", m)
	types.MarshalInto(m, w)
}
