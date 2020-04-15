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
	
	r.HandleFunc("/api/session-settings/update", UpdateSessionSettingWhere)
	r.HandleFunc("/api/session-settings/find", FindSessionSetting).Queries("limit", "{limit}")
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
