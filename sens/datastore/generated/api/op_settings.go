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

func OpSettingMain(r *mux.Router) {
	r.HandleFunc("/api/op-settings/create", CreateOpSetting)
	r.HandleFunc("/api/op-settings/batch/create", BatchCreateOpSetting)
	
	r.HandleFunc("/api/op-settings/{id}/update", UpdateOpSetting)
	r.HandleFunc("/api/op-settings/{id}/get", GetOpSetting)
    
	r.HandleFunc("/api/op-settings/update", UpdateOpSettingWhere)
	r.HandleFunc("/api/op-settings/find", FindOpSetting).Queries("limit", "{limit}")
	r.HandleFunc("/api/op-settings/delete", DeleteOpSetting)
}

func OpSettingRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateOpSetting(w http.ResponseWriter, r *http.Request) {
	defer OpSettingRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertOpSetting(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateOpSetting(w http.ResponseWriter, r *http.Request) {
	defer OpSettingRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertOpSetting(data)
	w.WriteHeader(http.StatusOK)
}


func UpdateOpSetting(w http.ResponseWriter, r *http.Request) {
	defer OpSettingRecovery(w)
	vars := mux.Vars(r)
	id := vars["id"]
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateOpSetting(id, data)
	w.WriteHeader(http.StatusOK)
}

func GetOpSetting(w http.ResponseWriter, r *http.Request) {
	defer OpSettingRecovery(w)
	vars := mux.Vars(r)
	id := vars["id"]
	m := fn.SelectOpSetting(id)
	types.MarshalInto(m, w)
}


func UpdateOpSettingWhere(w http.ResponseWriter, r *http.Request) {
	defer OpSettingRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateOpSettingWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindOpSetting(w http.ResponseWriter, r *http.Request) {
	defer OpSettingRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindOpSetting(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindOpSetting: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteOpSetting(w http.ResponseWriter, r *http.Request) {
	defer OpSettingRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteOpSetting(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteOpSetting: %d", n)
	types.MarshalInto(n, w)
}