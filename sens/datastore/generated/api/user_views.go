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

func UserViewMain(r *mux.Router) {
	r.HandleFunc("/api/user-views/create", CreateUserView)
	r.HandleFunc("/api/user-views/batch/create", BatchCreateUserView)
	
	r.HandleFunc("/api/user-views/update", UpdateUserViewWhere)
	r.HandleFunc("/api/user-views/find", FindUserView).Queries("limit", "{limit}")
	r.HandleFunc("/api/user-views/delete", DeleteUserView)
}

func UserViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateUserView(w http.ResponseWriter, r *http.Request) {
	defer UserViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	id := fn.InsertUserView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateUserView(w http.ResponseWriter, r *http.Request) {
	defer UserViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	fn.BatchUpsertUserView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateUserViewWhere(w http.ResponseWriter, r *http.Request) {
	defer UserViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateUserViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindUserView(w http.ResponseWriter, r *http.Request) {
	defer UserViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindUserView(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindUserView: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteUserView(w http.ResponseWriter, r *http.Request) {
	defer UserViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteUserView(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteUserView: %d", n)
	types.MarshalInto(n, w)
}