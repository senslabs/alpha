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

func AuthMain(r *mux.Router) {
	r.HandleFunc("/api/auths/create", CreateAuth)
	r.HandleFunc("/api/auths/batch/create", BatchCreateAuth)
	
	r.HandleFunc("/api/auths/{id}/update", UpdateAuth)
	r.HandleFunc("/api/auths/{id}/get", GetAuth)
    
	r.HandleFunc("/api/auths/update", UpdateAuthWhere)
	r.HandleFunc("/api/auths/find", FindAuth).Queries("limit", "{limit}")
}

func AuthRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateAuth(w http.ResponseWriter, r *http.Request) {
	defer AuthRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertAuth(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateAuth(w http.ResponseWriter, r *http.Request) {
	defer AuthRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertAuth(data)
	w.WriteHeader(http.StatusOK)
}


func UpdateAuth(w http.ResponseWriter, r *http.Request) {
	defer AuthRecovery(w)
	vars := mux.Vars(r)
	id := vars["id"]
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateAuth(id, data)
	w.WriteHeader(http.StatusOK)
}

func GetAuth(w http.ResponseWriter, r *http.Request) {
	defer AuthRecovery(w)
	vars := mux.Vars(r)
	id := vars["id"]
	m := fn.SelectAuth(id)
	types.MarshalInto(m, w)
}


func UpdateAuthWhere(w http.ResponseWriter, r *http.Request) {
	defer AuthRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateAuthWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindAuth(w http.ResponseWriter, r *http.Request) {
	defer AuthRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindAuth(or, and, in, span, limit, column, order)
	types.MarshalInto(m, w)
}
