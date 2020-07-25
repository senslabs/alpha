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

func AlertViewMain(r *mux.Router) {
	r.HandleFunc("/api/alert-views/create", CreateAlertView)
	r.HandleFunc("/api/alert-views/batch/create", BatchCreateAlertView)
	
	r.HandleFunc("/api/alert-views/update", UpdateAlertViewWhere)
	r.HandleFunc("/api/alert-views/find", FindAlertView).Queries("limit", "{limit}")
	r.HandleFunc("/api/alert-views/delete", DeleteAlertView)
}

func AlertViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateAlertView(w http.ResponseWriter, r *http.Request) {
	defer AlertViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	id := fn.InsertAlertView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateAlertView(w http.ResponseWriter, r *http.Request) {
	defer AlertViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	fn.BatchInsertAlertView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateAlertViewWhere(w http.ResponseWriter, r *http.Request) {
	defer AlertViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateAlertViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindAlertView(w http.ResponseWriter, r *http.Request) {
	defer AlertViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindAlertView(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindAlertView: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteAlertView(w http.ResponseWriter, r *http.Request) {
	defer AlertViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteAlertView(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteAlertView: %d", n)
	types.MarshalInto(n, w)
}