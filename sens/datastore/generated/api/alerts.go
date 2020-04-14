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

func AlertMain(r *mux.Router) {
	r.HandleFunc("/api/alerts/create", CreateAlert)
	r.HandleFunc("/api/alerts/batch/create", BatchCreateAlert)
	
	r.HandleFunc("/api/alerts/{id}/update", UpdateAlert)
	r.HandleFunc("/api/alerts/{id}/get", GetAlert)
    
	r.HandleFunc("/api/alerts/update", UpdateAlertWhere)
	r.HandleFunc("/api/alerts/find", FindAlert).Queries("limit", "{limit}")
}

func AlertRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateAlert(w http.ResponseWriter, r *http.Request) {
	defer AlertRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertAlert(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateAlert(w http.ResponseWriter, r *http.Request) {
	defer AlertRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertAlert(data)
	w.WriteHeader(http.StatusOK)
}


func UpdateAlert(w http.ResponseWriter, r *http.Request) {
	defer AlertRecovery(w)
	vars := mux.Vars(r)
	id := vars["id"]
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateAlert(id, data)
	w.WriteHeader(http.StatusOK)
}

func GetAlert(w http.ResponseWriter, r *http.Request) {
	defer AlertRecovery(w)
	vars := mux.Vars(r)
	id := vars["id"]
	m := fn.SelectAlert(id)
	types.MarshalInto(m, w)
}


func UpdateAlertWhere(w http.ResponseWriter, r *http.Request) {
	defer AlertRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateAlertWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindAlert(w http.ResponseWriter, r *http.Request) {
	defer AlertRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindAlert(or, and, in, span, limit, column, order)
	types.MarshalInto(m, w)
}
