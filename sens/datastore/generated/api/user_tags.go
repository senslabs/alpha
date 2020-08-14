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

func UserTagMain(r *mux.Router) {
	r.HandleFunc("/api/user-tags/create", CreateUserTag)
	r.HandleFunc("/api/user-tags/batch/create", BatchCreateUserTag)
	
	r.HandleFunc("/api/user-tags/{id}/update", UpdateUserTag)
	r.HandleFunc("/api/user-tags/{id}/get", GetUserTag)
    
	r.HandleFunc("/api/user-tags/update", UpdateUserTagWhere)
	r.HandleFunc("/api/user-tags/find", FindUserTag).Queries("limit", "{limit}")
	r.HandleFunc("/api/user-tags/delete", DeleteUserTag)
}

func UserTagRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateUserTag(w http.ResponseWriter, r *http.Request) {
	defer UserTagRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	id := fn.InsertUserTag(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateUserTag(w http.ResponseWriter, r *http.Request) {
	defer UserTagRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	fn.BatchUpsertUserTag(data)
	w.WriteHeader(http.StatusOK)
}


func UpdateUserTag(w http.ResponseWriter, r *http.Request) {
	defer UserTagRecovery(w)
	vars := mux.Vars(r)
	id := vars["id"]
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateUserTag(id, data)
	w.WriteHeader(http.StatusOK)
}

func GetUserTag(w http.ResponseWriter, r *http.Request) {
	defer UserTagRecovery(w)
	vars := mux.Vars(r)
	id := vars["id"]
	m := fn.SelectUserTag(id)
	types.MarshalInto(m, w)
}


func UpdateUserTagWhere(w http.ResponseWriter, r *http.Request) {
	defer UserTagRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateUserTagWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindUserTag(w http.ResponseWriter, r *http.Request) {
	defer UserTagRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindUserTag(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindUserTag: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteUserTag(w http.ResponseWriter, r *http.Request) {
	defer UserTagRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteUserTag(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteUserTag: %d", n)
	types.MarshalInto(n, w)
}