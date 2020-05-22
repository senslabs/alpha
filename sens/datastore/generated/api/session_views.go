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

func SessionViewMain(r *mux.Router) {
	r.HandleFunc("/api/session-views/create", CreateSessionView)
	r.HandleFunc("/api/session-views/batch/create", BatchCreateSessionView)
	
	r.HandleFunc("/api/session-views/update", UpdateSessionViewWhere)
	r.HandleFunc("/api/session-views/find", FindSessionView).Queries("limit", "{limit}")
	r.HandleFunc("/api/session-views/delete", DeleteSessionView)
}

func SessionViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateSessionView(w http.ResponseWriter, r *http.Request) {
	defer SessionViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertSessionView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateSessionView(w http.ResponseWriter, r *http.Request) {
	defer SessionViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertSessionView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateSessionViewWhere(w http.ResponseWriter, r *http.Request) {
	defer SessionViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateSessionViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindSessionView(w http.ResponseWriter, r *http.Request) {
	defer SessionViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindSessionView(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindSessionView: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteSessionView(w http.ResponseWriter, r *http.Request) {
	defer SessionViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteSessionView(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteSessionView: %d", n)
	types.MarshalInto(n, w)
}