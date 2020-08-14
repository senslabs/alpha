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

func DeviceMain(r *mux.Router) {
	r.HandleFunc("/api/devices/create", CreateDevice)
	r.HandleFunc("/api/devices/batch/create", BatchCreateDevice)
	
	r.HandleFunc("/api/devices/update", UpdateDeviceWhere)
	r.HandleFunc("/api/devices/find", FindDevice).Queries("limit", "{limit}")
	r.HandleFunc("/api/devices/delete", DeleteDevice)
}

func DeviceRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateDevice(w http.ResponseWriter, r *http.Request) {
	defer DeviceRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	id := fn.InsertDevice(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateDevice(w http.ResponseWriter, r *http.Request) {
	defer DeviceRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	fn.BatchUpsertDevice(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateDeviceWhere(w http.ResponseWriter, r *http.Request) {
	defer DeviceRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateDeviceWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindDevice(w http.ResponseWriter, r *http.Request) {
	defer DeviceRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindDevice(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindDevice: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteDevice(w http.ResponseWriter, r *http.Request) {
	defer DeviceRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteDevice(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteDevice: %d", n)
	types.MarshalInto(n, w)
}