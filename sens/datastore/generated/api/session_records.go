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

func SessionRecordMain(r *mux.Router) {
	r.HandleFunc("/api/session-records/create", CreateSessionRecord)
	r.HandleFunc("/api/session-records/batch/create", BatchCreateSessionRecord)
	
	r.HandleFunc("/api/session-records/update", UpdateSessionRecordWhere)
	r.HandleFunc("/api/session-records/find", FindSessionRecord).Queries("limit", "{limit}")
}

func SessionRecordRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateSessionRecord(w http.ResponseWriter, r *http.Request) {
	defer SessionRecordRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertSessionRecord(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateSessionRecord(w http.ResponseWriter, r *http.Request) {
	defer SessionRecordRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertSessionRecord(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateSessionRecordWhere(w http.ResponseWriter, r *http.Request) {
	defer SessionRecordRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateSessionRecordWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindSessionRecord(w http.ResponseWriter, r *http.Request) {
	defer SessionRecordRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindSessionRecord(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindSessionRecord: %#v", m)
	types.MarshalInto(m, w)
}
