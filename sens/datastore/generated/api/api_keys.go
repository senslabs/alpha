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

func ApiKeyMain(r *mux.Router) {
	r.HandleFunc("/api/api-keys/create", CreateApiKey)
	r.HandleFunc("/api/api-keys/batch/create", BatchCreateApiKey)
	
	r.HandleFunc("/api/api-keys/{id}/update", UpdateApiKey)
	r.HandleFunc("/api/api-keys/{id}/get", GetApiKey)
    
	r.HandleFunc("/api/api-keys/update", UpdateApiKeyWhere)
	r.HandleFunc("/api/api-keys/find", FindApiKey).Queries("limit", "{limit}")
}

func ApiKeyRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateApiKey(w http.ResponseWriter, r *http.Request) {
	defer ApiKeyRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertApiKey(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateApiKey(w http.ResponseWriter, r *http.Request) {
	defer ApiKeyRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertApiKey(data)
	w.WriteHeader(http.StatusOK)
}


func UpdateApiKey(w http.ResponseWriter, r *http.Request) {
	defer ApiKeyRecovery(w)
	vars := mux.Vars(r)
	id := vars["id"]
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateApiKey(id, data)
	w.WriteHeader(http.StatusOK)
}

func GetApiKey(w http.ResponseWriter, r *http.Request) {
	defer ApiKeyRecovery(w)
	vars := mux.Vars(r)
	id := vars["id"]
	m := fn.SelectApiKey(id)
	types.MarshalInto(m, w)
}


func UpdateApiKeyWhere(w http.ResponseWriter, r *http.Request) {
	defer ApiKeyRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateApiKeyWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindApiKey(w http.ResponseWriter, r *http.Request) {
	defer ApiKeyRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindApiKey(or, and, in, span, limit, column, order)
	types.MarshalInto(m, w)
}
