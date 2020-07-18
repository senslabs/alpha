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

func BaselineMain(r *mux.Router) {
	r.HandleFunc("/api/baselines/create", CreateBaseline)
	r.HandleFunc("/api/baselines/batch/create", BatchCreateBaseline)
	
	r.HandleFunc("/api/baselines/{id}/update", UpdateBaseline)
	r.HandleFunc("/api/baselines/{id}/get", GetBaseline)
    
	r.HandleFunc("/api/baselines/update", UpdateBaselineWhere)
	r.HandleFunc("/api/baselines/find", FindBaseline).Queries("limit", "{limit}")
	r.HandleFunc("/api/baselines/delete", DeleteBaseline)
}

func BaselineRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateBaseline(w http.ResponseWriter, r *http.Request) {
	defer BaselineRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	id := fn.InsertBaseline(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateBaseline(w http.ResponseWriter, r *http.Request) {
	defer BaselineRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	fn.BatchInsertBaseline(data)
	w.WriteHeader(http.StatusOK)
}


func UpdateBaseline(w http.ResponseWriter, r *http.Request) {
	defer BaselineRecovery(w)
	vars := mux.Vars(r)
	id := vars["id"]
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateBaseline(id, data)
	w.WriteHeader(http.StatusOK)
}

func GetBaseline(w http.ResponseWriter, r *http.Request) {
	defer BaselineRecovery(w)
	vars := mux.Vars(r)
	id := vars["id"]
	m := fn.SelectBaseline(id)
	types.MarshalInto(m, w)
}


func UpdateBaselineWhere(w http.ResponseWriter, r *http.Request) {
	defer BaselineRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateBaselineWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindBaseline(w http.ResponseWriter, r *http.Request) {
	defer BaselineRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindBaseline(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindBaseline: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteBaseline(w http.ResponseWriter, r *http.Request) {
	defer BaselineRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteBaseline(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteBaseline: %d", n)
	types.MarshalInto(n, w)
}