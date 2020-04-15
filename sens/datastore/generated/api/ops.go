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

func OpMain(r *mux.Router) {
	r.HandleFunc("/api/ops/create", CreateOp)
	r.HandleFunc("/api/ops/batch/create", BatchCreateOp)
	
	r.HandleFunc("/api/ops/{id}/update", UpdateOp)
	r.HandleFunc("/api/ops/{id}/get", GetOp)
    
	r.HandleFunc("/api/ops/update", UpdateOpWhere)
	r.HandleFunc("/api/ops/find", FindOp).Queries("limit", "{limit}")
}

func OpRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateOp(w http.ResponseWriter, r *http.Request) {
	defer OpRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertOp(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateOp(w http.ResponseWriter, r *http.Request) {
	defer OpRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertOp(data)
	w.WriteHeader(http.StatusOK)
}


func UpdateOp(w http.ResponseWriter, r *http.Request) {
	defer OpRecovery(w)
	vars := mux.Vars(r)
	id := vars["id"]
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateOp(id, data)
	w.WriteHeader(http.StatusOK)
}

func GetOp(w http.ResponseWriter, r *http.Request) {
	defer OpRecovery(w)
	vars := mux.Vars(r)
	id := vars["id"]
	m := fn.SelectOp(id)
	types.MarshalInto(m, w)
}


func UpdateOpWhere(w http.ResponseWriter, r *http.Request) {
	defer OpRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateOpWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindOp(w http.ResponseWriter, r *http.Request) {
	defer OpRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindOp(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindOp: %#v", m)
	types.MarshalInto(m, w)
}
