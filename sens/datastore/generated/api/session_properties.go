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

func SessionPropertieMain(r *mux.Router) {
	r.HandleFunc("/api/session-properties/create", CreateSessionPropertie)
	r.HandleFunc("/api/session-properties/batch/create", BatchCreateSessionPropertie)
	
	r.HandleFunc("/api/session-properties/update", UpdateSessionPropertieWhere)
	r.HandleFunc("/api/session-properties/find", FindSessionPropertie).Queries("limit", "{limit}")
	r.HandleFunc("/api/session-properties/delete", DeleteSessionPropertie)
}

func SessionPropertieRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateSessionPropertie(w http.ResponseWriter, r *http.Request) {
	defer SessionPropertieRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	id := fn.InsertSessionPropertie(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateSessionPropertie(w http.ResponseWriter, r *http.Request) {
	defer SessionPropertieRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	fn.BatchInsertSessionPropertie(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateSessionPropertieWhere(w http.ResponseWriter, r *http.Request) {
	defer SessionPropertieRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateSessionPropertieWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindSessionPropertie(w http.ResponseWriter, r *http.Request) {
	defer SessionPropertieRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindSessionPropertie(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindSessionPropertie: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteSessionPropertie(w http.ResponseWriter, r *http.Request) {
	defer SessionPropertieRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteSessionPropertie(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteSessionPropertie: %d", n)
	types.MarshalInto(n, w)
}