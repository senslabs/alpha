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

func UserSettingViewMain(r *mux.Router) {
	r.HandleFunc("/api/user-setting-views/create", CreateUserSettingView)
	r.HandleFunc("/api/user-setting-views/batch/create", BatchCreateUserSettingView)
	
	r.HandleFunc("/api/user-setting-views/update", UpdateUserSettingViewWhere)
	r.HandleFunc("/api/user-setting-views/find", FindUserSettingView).Queries("limit", "{limit}")
}

func UserSettingViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateUserSettingView(w http.ResponseWriter, r *http.Request) {
	defer UserSettingViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertUserSettingView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateUserSettingView(w http.ResponseWriter, r *http.Request) {
	defer UserSettingViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertUserSettingView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateUserSettingViewWhere(w http.ResponseWriter, r *http.Request) {
	defer UserSettingViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateUserSettingViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindUserSettingView(w http.ResponseWriter, r *http.Request) {
	defer UserSettingViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindUserSettingView(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindUserSettingView: %#v", m)
	types.MarshalInto(m, w)
}
