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

func UserSettingMain(r *mux.Router) {
	r.HandleFunc("/api/user-settings/create", CreateUserSetting)
	r.HandleFunc("/api/user-settings/batch/create", BatchCreateUserSetting)
	
	r.HandleFunc("/api/user-settings/{id}/update", UpdateUserSetting)
	r.HandleFunc("/api/user-settings/{id}/get", GetUserSetting)
    
	r.HandleFunc("/api/user-settings/update", UpdateUserSettingWhere)
	r.HandleFunc("/api/user-settings/find", FindUserSetting).Queries("limit", "{limit}")
	r.HandleFunc("/api/user-settings/delete", DeleteUserSetting)
}

func UserSettingRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateUserSetting(w http.ResponseWriter, r *http.Request) {
	defer UserSettingRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertUserSetting(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateUserSetting(w http.ResponseWriter, r *http.Request) {
	defer UserSettingRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertUserSetting(data)
	w.WriteHeader(http.StatusOK)
}


func UpdateUserSetting(w http.ResponseWriter, r *http.Request) {
	defer UserSettingRecovery(w)
	vars := mux.Vars(r)
	id := vars["id"]
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateUserSetting(id, data)
	w.WriteHeader(http.StatusOK)
}

func GetUserSetting(w http.ResponseWriter, r *http.Request) {
	defer UserSettingRecovery(w)
	vars := mux.Vars(r)
	id := vars["id"]
	m := fn.SelectUserSetting(id)
	types.MarshalInto(m, w)
}


func UpdateUserSettingWhere(w http.ResponseWriter, r *http.Request) {
	defer UserSettingRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateUserSettingWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindUserSetting(w http.ResponseWriter, r *http.Request) {
	defer UserSettingRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindUserSetting(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindUserSetting: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteUserSetting(w http.ResponseWriter, r *http.Request) {
	defer UserSettingRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteUserSetting(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteUserSetting: %d", n)
	types.MarshalInto(n, w)
}