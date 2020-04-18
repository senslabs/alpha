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

func DevicePropertieMain(r *mux.Router) {
	r.HandleFunc("/api/device-properties/create", CreateDevicePropertie)
	r.HandleFunc("/api/device-properties/batch/create", BatchCreateDevicePropertie)
	
	r.HandleFunc("/api/device-properties/update", UpdateDevicePropertieWhere)
	r.HandleFunc("/api/device-properties/find", FindDevicePropertie).Queries("limit", "{limit}")
}

func DevicePropertieRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateDevicePropertie(w http.ResponseWriter, r *http.Request) {
	defer DevicePropertieRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertDevicePropertie(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateDevicePropertie(w http.ResponseWriter, r *http.Request) {
	defer DevicePropertieRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertDevicePropertie(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateDevicePropertieWhere(w http.ResponseWriter, r *http.Request) {
	defer DevicePropertieRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateDevicePropertieWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindDevicePropertie(w http.ResponseWriter, r *http.Request) {
	defer DevicePropertieRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindDevicePropertie(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindDevicePropertie: %#v", m)
	types.MarshalInto(m, w)
}
