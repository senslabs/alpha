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

func OpViewMain(r *mux.Router) {
	r.HandleFunc("/api/op-views/create", CreateOpView)
	r.HandleFunc("/api/op-views/batch/create", BatchCreateOpView)
	
	r.HandleFunc("/api/op-views/update", UpdateOpViewWhere)
	r.HandleFunc("/api/op-views/find", FindOpView).Queries("limit", "{limit}")
	r.HandleFunc("/api/op-views/delete", DeleteOpView)
}

func OpViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateOpView(w http.ResponseWriter, r *http.Request) {
	defer OpViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	id := fn.InsertOpView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateOpView(w http.ResponseWriter, r *http.Request) {
	defer OpViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	fn.BatchInsertOpView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateOpViewWhere(w http.ResponseWriter, r *http.Request) {
	defer OpViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateOpViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindOpView(w http.ResponseWriter, r *http.Request) {
	defer OpViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindOpView(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindOpView: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteOpView(w http.ResponseWriter, r *http.Request) {
	defer OpViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteOpView(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteOpView: %d", n)
	types.MarshalInto(n, w)
}