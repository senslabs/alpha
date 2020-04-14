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

func UserDetailViewMain(r *mux.Router) {
	r.HandleFunc("/api/user-detail-views/create", CreateUserDetailView)
	r.HandleFunc("/api/user-detail-views/batch/create", BatchCreateUserDetailView)
	
	r.HandleFunc("/api/user-detail-views/update", UpdateUserDetailViewWhere)
	r.HandleFunc("/api/user-detail-views/find", FindUserDetailView).Queries("limit", "{limit}")
}

func UserDetailViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateUserDetailView(w http.ResponseWriter, r *http.Request) {
	defer UserDetailViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertUserDetailView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateUserDetailView(w http.ResponseWriter, r *http.Request) {
	defer UserDetailViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertUserDetailView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateUserDetailViewWhere(w http.ResponseWriter, r *http.Request) {
	defer UserDetailViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateUserDetailViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindUserDetailView(w http.ResponseWriter, r *http.Request) {
	defer UserDetailViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindUserDetailView(or, and, in, span, limit, column, order)
	types.MarshalInto(m, w)
}
