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

func OrgSettingMain(r *mux.Router) {
	r.HandleFunc("/api/org-settings/create", CreateOrgSetting)
	r.HandleFunc("/api/org-settings/batch/create", BatchCreateOrgSetting)
	
	r.HandleFunc("/api/org-settings/{id}/update", UpdateOrgSetting)
	r.HandleFunc("/api/org-settings/{id}/get", GetOrgSetting)
    
	r.HandleFunc("/api/org-settings/update", UpdateOrgSettingWhere)
	r.HandleFunc("/api/org-settings/find", FindOrgSetting).Queries("limit", "{limit}")
	r.HandleFunc("/api/org-settings/delete", DeleteOrgSetting)
}

func OrgSettingRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateOrgSetting(w http.ResponseWriter, r *http.Request) {
	defer OrgSettingRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertOrgSetting(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateOrgSetting(w http.ResponseWriter, r *http.Request) {
	defer OrgSettingRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertOrgSetting(data)
	w.WriteHeader(http.StatusOK)
}


func UpdateOrgSetting(w http.ResponseWriter, r *http.Request) {
	defer OrgSettingRecovery(w)
	vars := mux.Vars(r)
	id := vars["id"]
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateOrgSetting(id, data)
	w.WriteHeader(http.StatusOK)
}

func GetOrgSetting(w http.ResponseWriter, r *http.Request) {
	defer OrgSettingRecovery(w)
	vars := mux.Vars(r)
	id := vars["id"]
	m := fn.SelectOrgSetting(id)
	types.MarshalInto(m, w)
}


func UpdateOrgSettingWhere(w http.ResponseWriter, r *http.Request) {
	defer OrgSettingRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateOrgSettingWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindOrgSetting(w http.ResponseWriter, r *http.Request) {
	defer OrgSettingRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindOrgSetting(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindOrgSetting: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteOrgSetting(w http.ResponseWriter, r *http.Request) {
	defer OrgSettingRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteOrgSetting(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteOrgSetting: %d", n)
	types.MarshalInto(n, w)
}