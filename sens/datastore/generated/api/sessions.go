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

func SessionMain(r *mux.Router) {
	r.HandleFunc("/api/sessions/create", CreateSession)
	r.HandleFunc("/api/sessions/batch/create", BatchCreateSession)
	
	r.HandleFunc("/api/sessions/{id}/update", UpdateSession)
	r.HandleFunc("/api/sessions/{id}/get", GetSession)
    
	r.HandleFunc("/api/sessions/update", UpdateSessionWhere)
	r.HandleFunc("/api/sessions/find", FindSession).Queries("limit", "{limit}")
	r.HandleFunc("/api/sessions/delete", DeleteSession)
}

func SessionRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateSession(w http.ResponseWriter, r *http.Request) {
	defer SessionRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	id := fn.InsertSession(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateSession(w http.ResponseWriter, r *http.Request) {
	defer SessionRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	fn.BatchUpsertSession(data)
	w.WriteHeader(http.StatusOK)
}


func UpdateSession(w http.ResponseWriter, r *http.Request) {
	defer SessionRecovery(w)
	vars := mux.Vars(r)
	id := vars["id"]
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateSession(id, data)
	w.WriteHeader(http.StatusOK)
}

func GetSession(w http.ResponseWriter, r *http.Request) {
	defer SessionRecovery(w)
	vars := mux.Vars(r)
	id := vars["id"]
	m := fn.SelectSession(id)
	types.MarshalInto(m, w)
}


func UpdateSessionWhere(w http.ResponseWriter, r *http.Request) {
	defer SessionRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateSessionWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindSession(w http.ResponseWriter, r *http.Request) {
	defer SessionRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindSession(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindSession: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteSession(w http.ResponseWriter, r *http.Request) {
	defer SessionRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteSession(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteSession: %d", n)
	types.MarshalInto(n, w)
}