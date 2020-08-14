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

func AuthViewMain(r *mux.Router) {
	r.HandleFunc("/api/auth-views/create", CreateAuthView)
	r.HandleFunc("/api/auth-views/batch/create", BatchCreateAuthView)
	
	r.HandleFunc("/api/auth-views/update", UpdateAuthViewWhere)
	r.HandleFunc("/api/auth-views/find", FindAuthView).Queries("limit", "{limit}")
	r.HandleFunc("/api/auth-views/delete", DeleteAuthView)
}

func AuthViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateAuthView(w http.ResponseWriter, r *http.Request) {
	defer AuthViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	id := fn.InsertAuthView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateAuthView(w http.ResponseWriter, r *http.Request) {
	defer AuthViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	fn.BatchUpsertAuthView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateAuthViewWhere(w http.ResponseWriter, r *http.Request) {
	defer AuthViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateAuthViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindAuthView(w http.ResponseWriter, r *http.Request) {
	defer AuthViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindAuthView(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindAuthView: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteAuthView(w http.ResponseWriter, r *http.Request) {
	defer AuthViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteAuthView(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteAuthView: %d", n)
	types.MarshalInto(n, w)
}