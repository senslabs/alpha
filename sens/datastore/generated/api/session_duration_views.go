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

func SessionDurationViewMain(r *mux.Router) {
	r.HandleFunc("/api/session-duration-views/create", CreateSessionDurationView)
	r.HandleFunc("/api/session-duration-views/batch/create", BatchCreateSessionDurationView)
	
	r.HandleFunc("/api/session-duration-views/update", UpdateSessionDurationViewWhere)
	r.HandleFunc("/api/session-duration-views/find", FindSessionDurationView).Queries("limit", "{limit}")
	r.HandleFunc("/api/session-duration-views/delete", DeleteSessionDurationView)
}

func SessionDurationViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateSessionDurationView(w http.ResponseWriter, r *http.Request) {
	defer SessionDurationViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	id := fn.InsertSessionDurationView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateSessionDurationView(w http.ResponseWriter, r *http.Request) {
	defer SessionDurationViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	fn.BatchInsertSessionDurationView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateSessionDurationViewWhere(w http.ResponseWriter, r *http.Request) {
	defer SessionDurationViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateSessionDurationViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindSessionDurationView(w http.ResponseWriter, r *http.Request) {
	defer SessionDurationViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindSessionDurationView(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindSessionDurationView: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteSessionDurationView(w http.ResponseWriter, r *http.Request) {
	defer SessionDurationViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteSessionDurationView(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteSessionDurationView: %d", n)
	types.MarshalInto(n, w)
}