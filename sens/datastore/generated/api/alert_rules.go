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

func AlertRuleMain(r *mux.Router) {
	r.HandleFunc("/api/alert-rules/create", CreateAlertRule)
	r.HandleFunc("/api/alert-rules/batch/create", BatchCreateAlertRule)
	
	r.HandleFunc("/api/alert-rules/{id}/update", UpdateAlertRule)
	r.HandleFunc("/api/alert-rules/{id}/get", GetAlertRule)
    
	r.HandleFunc("/api/alert-rules/update", UpdateAlertRuleWhere)
	r.HandleFunc("/api/alert-rules/find", FindAlertRule).Queries("limit", "{limit}")
	r.HandleFunc("/api/alert-rules/delete", DeleteAlertRule)
}

func AlertRuleRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateAlertRule(w http.ResponseWriter, r *http.Request) {
	defer AlertRuleRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	id := fn.InsertAlertRule(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateAlertRule(w http.ResponseWriter, r *http.Request) {
	defer AlertRuleRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	fn.BatchUpsertAlertRule(data)
	w.WriteHeader(http.StatusOK)
}


func UpdateAlertRule(w http.ResponseWriter, r *http.Request) {
	defer AlertRuleRecovery(w)
	vars := mux.Vars(r)
	id := vars["id"]
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateAlertRule(id, data)
	w.WriteHeader(http.StatusOK)
}

func GetAlertRule(w http.ResponseWriter, r *http.Request) {
	defer AlertRuleRecovery(w)
	vars := mux.Vars(r)
	id := vars["id"]
	m := fn.SelectAlertRule(id)
	types.MarshalInto(m, w)
}


func UpdateAlertRuleWhere(w http.ResponseWriter, r *http.Request) {
	defer AlertRuleRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateAlertRuleWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindAlertRule(w http.ResponseWriter, r *http.Request) {
	defer AlertRuleRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindAlertRule(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindAlertRule: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteAlertRule(w http.ResponseWriter, r *http.Request) {
	defer AlertRuleRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteAlertRule(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteAlertRule: %d", n)
	types.MarshalInto(n, w)
}