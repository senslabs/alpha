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

func OrgMeditationViewMain(r *mux.Router) {
	r.HandleFunc("/api/org-meditation-views/create", CreateOrgMeditationView)
	r.HandleFunc("/api/org-meditation-views/batch/create", BatchCreateOrgMeditationView)
	
	r.HandleFunc("/api/org-meditation-views/update", UpdateOrgMeditationViewWhere)
	r.HandleFunc("/api/org-meditation-views/find", FindOrgMeditationView).Queries("limit", "{limit}")
}

func OrgMeditationViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateOrgMeditationView(w http.ResponseWriter, r *http.Request) {
	defer OrgMeditationViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertOrgMeditationView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateOrgMeditationView(w http.ResponseWriter, r *http.Request) {
	defer OrgMeditationViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertOrgMeditationView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateOrgMeditationViewWhere(w http.ResponseWriter, r *http.Request) {
	defer OrgMeditationViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateOrgMeditationViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindOrgMeditationView(w http.ResponseWriter, r *http.Request) {
	defer OrgMeditationViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindOrgMeditationView(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindOrgMeditationView: %#v", m)
	types.MarshalInto(m, w)
}
