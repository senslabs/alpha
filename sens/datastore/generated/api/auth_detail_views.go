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

func AuthDetailViewMain(r *mux.Router) {
	r.HandleFunc("/api/auth-detail-views/create", CreateAuthDetailView)
	r.HandleFunc("/api/auth-detail-views/batch/create", BatchCreateAuthDetailView)
	
	r.HandleFunc("/api/auth-detail-views/update", UpdateAuthDetailViewWhere)
	r.HandleFunc("/api/auth-detail-views/find", FindAuthDetailView).Queries("limit", "{limit}")
}

func AuthDetailViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateAuthDetailView(w http.ResponseWriter, r *http.Request) {
	defer AuthDetailViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertAuthDetailView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateAuthDetailView(w http.ResponseWriter, r *http.Request) {
	defer AuthDetailViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertAuthDetailView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateAuthDetailViewWhere(w http.ResponseWriter, r *http.Request) {
	defer AuthDetailViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateAuthDetailViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindAuthDetailView(w http.ResponseWriter, r *http.Request) {
	defer AuthDetailViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindAuthDetailView(or, and, in, span, limit, column, order)
	types.MarshalInto(m, w)
}
