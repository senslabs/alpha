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

func OpPropertieMain(r *mux.Router) {
	r.HandleFunc("/api/op-properties/create", CreateOpPropertie)
	r.HandleFunc("/api/op-properties/batch/create", BatchCreateOpPropertie)
	
	r.HandleFunc("/api/op-properties/update", UpdateOpPropertieWhere)
	r.HandleFunc("/api/op-properties/find", FindOpPropertie).Queries("limit", "{limit}")
	r.HandleFunc("/api/op-properties/delete", DeleteOpPropertie)
}

func OpPropertieRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateOpPropertie(w http.ResponseWriter, r *http.Request) {
	defer OpPropertieRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	id := fn.InsertOpPropertie(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateOpPropertie(w http.ResponseWriter, r *http.Request) {
	defer OpPropertieRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	fn.BatchUpsertOpPropertie(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateOpPropertieWhere(w http.ResponseWriter, r *http.Request) {
	defer OpPropertieRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateOpPropertieWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindOpPropertie(w http.ResponseWriter, r *http.Request) {
	defer OpPropertieRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindOpPropertie(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindOpPropertie: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteOpPropertie(w http.ResponseWriter, r *http.Request) {
	defer OpPropertieRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteOpPropertie(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteOpPropertie: %d", n)
	types.MarshalInto(n, w)
}