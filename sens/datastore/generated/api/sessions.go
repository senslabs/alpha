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
	
	r.HandleFunc("/api/sessions/update", UpdateSessionWhere)
	r.HandleFunc("/api/sessions/find", FindSession).Queries("limit", "{limit}")
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
	id := fn.InsertSession(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateSession(w http.ResponseWriter, r *http.Request) {
	defer SessionRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertSession(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateSessionWhere(w http.ResponseWriter, r *http.Request) {
	defer SessionRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
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
	types.MarshalInto(m, w)
}
