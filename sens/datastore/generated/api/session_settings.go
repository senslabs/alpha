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

func SessionSettingMain(r *mux.Router) {
	r.HandleFunc("/api/session-settings/create", CreateSessionSetting)
	r.HandleFunc("/api/session-settings/batch/create", BatchCreateSessionSetting)
	
	r.HandleFunc("/api/session-settings/{id}/update", UpdateSessionSetting)
	r.HandleFunc("/api/session-settings/{id}/get", GetSessionSetting)
    
	r.HandleFunc("/api/session-settings/update", UpdateSessionSettingWhere)
	r.HandleFunc("/api/session-settings/find", FindSessionSetting).Queries("limit", "{limit}")
	r.HandleFunc("/api/session-settings/delete", DeleteSessionSetting)
}

func SessionSettingRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateSessionSetting(w http.ResponseWriter, r *http.Request) {
	defer SessionSettingRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertSessionSetting(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateSessionSetting(w http.ResponseWriter, r *http.Request) {
	defer SessionSettingRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertSessionSetting(data)
	w.WriteHeader(http.StatusOK)
}


func UpdateSessionSetting(w http.ResponseWriter, r *http.Request) {
	defer SessionSettingRecovery(w)
	vars := mux.Vars(r)
	id := vars["id"]
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateSessionSetting(id, data)
	w.WriteHeader(http.StatusOK)
}

func GetSessionSetting(w http.ResponseWriter, r *http.Request) {
	defer SessionSettingRecovery(w)
	vars := mux.Vars(r)
	id := vars["id"]
	m := fn.SelectSessionSetting(id)
	types.MarshalInto(m, w)
}


func UpdateSessionSettingWhere(w http.ResponseWriter, r *http.Request) {
	defer SessionSettingRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateSessionSettingWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindSessionSetting(w http.ResponseWriter, r *http.Request) {
	defer SessionSettingRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindSessionSetting(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindSessionSetting: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteSessionSetting(w http.ResponseWriter, r *http.Request) {
	defer SessionSettingRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteSessionSetting(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteSessionSetting: %d", n)
	types.MarshalInto(n, w)
}