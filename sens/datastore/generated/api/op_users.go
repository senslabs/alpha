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

func OpUserMain(r *mux.Router) {
	r.HandleFunc("/api/op-users/create", CreateOpUser)
	r.HandleFunc("/api/op-users/batch/create", BatchCreateOpUser)
	
	r.HandleFunc("/api/op-users/update", UpdateOpUserWhere)
	r.HandleFunc("/api/op-users/find", FindOpUser).Queries("limit", "{limit}")
	r.HandleFunc("/api/op-users/delete", DeleteOpUser)
}

func OpUserRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateOpUser(w http.ResponseWriter, r *http.Request) {
	defer OpUserRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	id := fn.InsertOpUser(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateOpUser(w http.ResponseWriter, r *http.Request) {
	defer OpUserRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	fn.BatchInsertOpUser(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateOpUserWhere(w http.ResponseWriter, r *http.Request) {
	defer OpUserRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateOpUserWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindOpUser(w http.ResponseWriter, r *http.Request) {
	defer OpUserRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindOpUser(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindOpUser: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteOpUser(w http.ResponseWriter, r *http.Request) {
	defer OpUserRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteOpUser(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteOpUser: %d", n)
	types.MarshalInto(n, w)
}