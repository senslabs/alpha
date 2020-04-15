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

func UserPropertieMain(r *mux.Router) {
	r.HandleFunc("/api/user-properties/create", CreateUserPropertie)
	r.HandleFunc("/api/user-properties/batch/create", BatchCreateUserPropertie)
	
	r.HandleFunc("/api/user-properties/update", UpdateUserPropertieWhere)
	r.HandleFunc("/api/user-properties/find", FindUserPropertie).Queries("limit", "{limit}")
}

func UserPropertieRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateUserPropertie(w http.ResponseWriter, r *http.Request) {
	defer UserPropertieRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertUserPropertie(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateUserPropertie(w http.ResponseWriter, r *http.Request) {
	defer UserPropertieRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertUserPropertie(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateUserPropertieWhere(w http.ResponseWriter, r *http.Request) {
	defer UserPropertieRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateUserPropertieWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindUserPropertie(w http.ResponseWriter, r *http.Request) {
	defer UserPropertieRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindUserPropertie(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindUserPropertie: %#v", m)
	types.MarshalInto(m, w)
}
