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

func CollectorMain(r *mux.Router) {
	r.HandleFunc("/api/collectors/create", CreateCollector)
	r.HandleFunc("/api/collectors/batch/create", BatchCreateCollector)
	
	r.HandleFunc("/api/collectors/update", UpdateCollectorWhere)
	r.HandleFunc("/api/collectors/find", FindCollector).Queries("limit", "{limit}")
	r.HandleFunc("/api/collectors/delete", DeleteCollector)
}

func CollectorRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateCollector(w http.ResponseWriter, r *http.Request) {
	defer CollectorRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertCollector(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateCollector(w http.ResponseWriter, r *http.Request) {
	defer CollectorRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertCollector(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateCollectorWhere(w http.ResponseWriter, r *http.Request) {
	defer CollectorRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateCollectorWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindCollector(w http.ResponseWriter, r *http.Request) {
	defer CollectorRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindCollector(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindCollector: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteCollector(w http.ResponseWriter, r *http.Request) {
	defer CollectorRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteCollector(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteCollector: %d", n)
	types.MarshalInto(n, w)
}