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

func CollectorPropertieMain(r *mux.Router) {
	r.HandleFunc("/api/collector-properties/create", CreateCollectorPropertie)
	r.HandleFunc("/api/collector-properties/batch/create", BatchCreateCollectorPropertie)
	
	r.HandleFunc("/api/collector-properties/update", UpdateCollectorPropertieWhere)
	r.HandleFunc("/api/collector-properties/find", FindCollectorPropertie).Queries("limit", "{limit}")
	r.HandleFunc("/api/collector-properties/delete", DeleteCollectorPropertie)
}

func CollectorPropertieRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateCollectorPropertie(w http.ResponseWriter, r *http.Request) {
	defer CollectorPropertieRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertCollectorPropertie(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateCollectorPropertie(w http.ResponseWriter, r *http.Request) {
	defer CollectorPropertieRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertCollectorPropertie(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateCollectorPropertieWhere(w http.ResponseWriter, r *http.Request) {
	defer CollectorPropertieRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateCollectorPropertieWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindCollectorPropertie(w http.ResponseWriter, r *http.Request) {
	defer CollectorPropertieRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindCollectorPropertie(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindCollectorPropertie: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteCollectorPropertie(w http.ResponseWriter, r *http.Request) {
	defer CollectorPropertieRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteCollectorPropertie(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteCollectorPropertie: %d", n)
	types.MarshalInto(n, w)
}