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

func SessionEventMain(r *mux.Router) {
	r.HandleFunc("/api/session-events/create", CreateSessionEvent)
	r.HandleFunc("/api/session-events/batch/create", BatchCreateSessionEvent)
	
	r.HandleFunc("/api/session-events/update", UpdateSessionEventWhere)
	r.HandleFunc("/api/session-events/find", FindSessionEvent).Queries("limit", "{limit}")
	r.HandleFunc("/api/session-events/delete", DeleteSessionEvent)
}

func SessionEventRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateSessionEvent(w http.ResponseWriter, r *http.Request) {
	defer SessionEventRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	id := fn.InsertSessionEvent(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateSessionEvent(w http.ResponseWriter, r *http.Request) {
	defer SessionEventRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	fn.BatchInsertSessionEvent(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateSessionEventWhere(w http.ResponseWriter, r *http.Request) {
	defer SessionEventRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateSessionEventWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindSessionEvent(w http.ResponseWriter, r *http.Request) {
	defer SessionEventRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindSessionEvent(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindSessionEvent: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteSessionEvent(w http.ResponseWriter, r *http.Request) {
	defer SessionEventRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteSessionEvent(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteSessionEvent: %d", n)
	types.MarshalInto(n, w)
}