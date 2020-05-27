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

func UserBaselineViewMain(r *mux.Router) {
	r.HandleFunc("/api/user-baseline-views/create", CreateUserBaselineView)
	r.HandleFunc("/api/user-baseline-views/batch/create", BatchCreateUserBaselineView)
	
	r.HandleFunc("/api/user-baseline-views/update", UpdateUserBaselineViewWhere)
	r.HandleFunc("/api/user-baseline-views/find", FindUserBaselineView).Queries("limit", "{limit}")
	r.HandleFunc("/api/user-baseline-views/delete", DeleteUserBaselineView)
}

func UserBaselineViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateUserBaselineView(w http.ResponseWriter, r *http.Request) {
	defer UserBaselineViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertUserBaselineView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateUserBaselineView(w http.ResponseWriter, r *http.Request) {
	defer UserBaselineViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertUserBaselineView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateUserBaselineViewWhere(w http.ResponseWriter, r *http.Request) {
	defer UserBaselineViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateUserBaselineViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindUserBaselineView(w http.ResponseWriter, r *http.Request) {
	defer UserBaselineViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindUserBaselineView(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindUserBaselineView: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteUserBaselineView(w http.ResponseWriter, r *http.Request) {
	defer UserBaselineViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteUserBaselineView(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteUserBaselineView: %d", n)
	types.MarshalInto(n, w)
}