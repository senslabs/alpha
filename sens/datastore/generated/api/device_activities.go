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

func DeviceActivitieMain(r *mux.Router) {
	r.HandleFunc("/api/device-activities/create", CreateDeviceActivitie)
	r.HandleFunc("/api/device-activities/batch/create", BatchCreateDeviceActivitie)
	
	r.HandleFunc("/api/device-activities/update", UpdateDeviceActivitieWhere)
	r.HandleFunc("/api/device-activities/find", FindDeviceActivitie).Queries("limit", "{limit}")
	r.HandleFunc("/api/device-activities/delete", DeleteDeviceActivitie)
}

func DeviceActivitieRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateDeviceActivitie(w http.ResponseWriter, r *http.Request) {
	defer DeviceActivitieRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertDeviceActivitie(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateDeviceActivitie(w http.ResponseWriter, r *http.Request) {
	defer DeviceActivitieRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertDeviceActivitie(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateDeviceActivitieWhere(w http.ResponseWriter, r *http.Request) {
	defer DeviceActivitieRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateDeviceActivitieWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindDeviceActivitie(w http.ResponseWriter, r *http.Request) {
	defer DeviceActivitieRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindDeviceActivitie(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindDeviceActivitie: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteDeviceActivitie(w http.ResponseWriter, r *http.Request) {
	defer DeviceActivitieRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteDeviceActivitie(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteDeviceActivitie: %d", n)
	types.MarshalInto(n, w)
}