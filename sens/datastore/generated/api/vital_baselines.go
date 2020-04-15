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

func VitalBaselineMain(r *mux.Router) {
	r.HandleFunc("/api/vital-baselines/create", CreateVitalBaseline)
	r.HandleFunc("/api/vital-baselines/batch/create", BatchCreateVitalBaseline)
	
	r.HandleFunc("/api/vital-baselines/update", UpdateVitalBaselineWhere)
	r.HandleFunc("/api/vital-baselines/find", FindVitalBaseline).Queries("limit", "{limit}")
}

func VitalBaselineRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateVitalBaseline(w http.ResponseWriter, r *http.Request) {
	defer VitalBaselineRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertVitalBaseline(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateVitalBaseline(w http.ResponseWriter, r *http.Request) {
	defer VitalBaselineRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertVitalBaseline(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateVitalBaselineWhere(w http.ResponseWriter, r *http.Request) {
	defer VitalBaselineRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateVitalBaselineWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindVitalBaseline(w http.ResponseWriter, r *http.Request) {
	defer VitalBaselineRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindVitalBaseline(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindVitalBaseline: %#v", m)
	types.MarshalInto(m, w)
}
