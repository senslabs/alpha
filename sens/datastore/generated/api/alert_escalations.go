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

func AlertEscalationMain(r *mux.Router) {
	r.HandleFunc("/api/alert-escalations/create", CreateAlertEscalation)
	r.HandleFunc("/api/alert-escalations/batch/create", BatchCreateAlertEscalation)
	
	r.HandleFunc("/api/alert-escalations/update", UpdateAlertEscalationWhere)
	r.HandleFunc("/api/alert-escalations/find", FindAlertEscalation).Queries("limit", "{limit}")
}

func AlertEscalationRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateAlertEscalation(w http.ResponseWriter, r *http.Request) {
	defer AlertEscalationRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertAlertEscalation(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateAlertEscalation(w http.ResponseWriter, r *http.Request) {
	defer AlertEscalationRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertAlertEscalation(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateAlertEscalationWhere(w http.ResponseWriter, r *http.Request) {
	defer AlertEscalationRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateAlertEscalationWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindAlertEscalation(w http.ResponseWriter, r *http.Request) {
	defer AlertEscalationRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindAlertEscalation(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindAlertEscalation: %#v", m)
	types.MarshalInto(m, w)
}
