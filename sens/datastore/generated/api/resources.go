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

func ResourceMain(r *mux.Router) {
	r.HandleFunc("/api/resources/create", CreateResource)
	r.HandleFunc("/api/resources/batch/create", BatchCreateResource)
	
	r.HandleFunc("/api/resources/{id}/update", UpdateResource)
	r.HandleFunc("/api/resources/{id}/get", GetResource)
    
	r.HandleFunc("/api/resources/update", UpdateResourceWhere)
	r.HandleFunc("/api/resources/find", FindResource).Queries("limit", "{limit}")
	r.HandleFunc("/api/resources/delete", DeleteResource)
}

func ResourceRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateResource(w http.ResponseWriter, r *http.Request) {
	defer ResourceRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	id := fn.InsertResource(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateResource(w http.ResponseWriter, r *http.Request) {
	defer ResourceRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	fn.BatchUpsertResource(data)
	w.WriteHeader(http.StatusOK)
}


func UpdateResource(w http.ResponseWriter, r *http.Request) {
	defer ResourceRecovery(w)
	vars := mux.Vars(r)
	id := vars["id"]
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateResource(id, data)
	w.WriteHeader(http.StatusOK)
}

func GetResource(w http.ResponseWriter, r *http.Request) {
	defer ResourceRecovery(w)
	vars := mux.Vars(r)
	id := vars["id"]
	m := fn.SelectResource(id)
	types.MarshalInto(m, w)
}


func UpdateResourceWhere(w http.ResponseWriter, r *http.Request) {
	defer ResourceRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateResourceWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindResource(w http.ResponseWriter, r *http.Request) {
	defer ResourceRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindResource(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindResource: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteResource(w http.ResponseWriter, r *http.Request) {
	defer ResourceRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteResource(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteResource: %d", n)
	types.MarshalInto(n, w)
}