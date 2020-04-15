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

func OpDetailViewMain(r *mux.Router) {
	r.HandleFunc("/api/op-detail-views/create", CreateOpDetailView)
	r.HandleFunc("/api/op-detail-views/batch/create", BatchCreateOpDetailView)
	
	r.HandleFunc("/api/op-detail-views/update", UpdateOpDetailViewWhere)
	r.HandleFunc("/api/op-detail-views/find", FindOpDetailView).Queries("limit", "{limit}")
}

func OpDetailViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateOpDetailView(w http.ResponseWriter, r *http.Request) {
	defer OpDetailViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertOpDetailView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateOpDetailView(w http.ResponseWriter, r *http.Request) {
	defer OpDetailViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertOpDetailView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateOpDetailViewWhere(w http.ResponseWriter, r *http.Request) {
	defer OpDetailViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateOpDetailViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindOpDetailView(w http.ResponseWriter, r *http.Request) {
	defer OpDetailViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindOpDetailView(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindOpDetailView: %#v", m)
	types.MarshalInto(m, w)
}
