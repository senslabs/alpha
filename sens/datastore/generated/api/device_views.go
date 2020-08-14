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

func DeviceViewMain(r *mux.Router) {
	r.HandleFunc("/api/device-views/create", CreateDeviceView)
	r.HandleFunc("/api/device-views/batch/create", BatchCreateDeviceView)
	
	r.HandleFunc("/api/device-views/update", UpdateDeviceViewWhere)
	r.HandleFunc("/api/device-views/find", FindDeviceView).Queries("limit", "{limit}")
	r.HandleFunc("/api/device-views/delete", DeleteDeviceView)
}

func DeviceViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateDeviceView(w http.ResponseWriter, r *http.Request) {
	defer DeviceViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	id := fn.InsertDeviceView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateDeviceView(w http.ResponseWriter, r *http.Request) {
	defer DeviceViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	fn.BatchUpsertDeviceView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateDeviceViewWhere(w http.ResponseWriter, r *http.Request) {
	defer DeviceViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateDeviceViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindDeviceView(w http.ResponseWriter, r *http.Request) {
	defer DeviceViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindDeviceView(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindDeviceView: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteDeviceView(w http.ResponseWriter, r *http.Request) {
	defer DeviceViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteDeviceView(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteDeviceView: %d", n)
	types.MarshalInto(n, w)
}