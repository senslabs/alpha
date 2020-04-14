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

func UserMeditationViewMain(r *mux.Router) {
	r.HandleFunc("/api/user-meditation-views/create", CreateUserMeditationView)
	r.HandleFunc("/api/user-meditation-views/batch/create", BatchCreateUserMeditationView)
	
	r.HandleFunc("/api/user-meditation-views/update", UpdateUserMeditationViewWhere)
	r.HandleFunc("/api/user-meditation-views/find", FindUserMeditationView).Queries("limit", "{limit}")
}

func UserMeditationViewRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateUserMeditationView(w http.ResponseWriter, r *http.Request) {
	defer UserMeditationViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertUserMeditationView(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateUserMeditationView(w http.ResponseWriter, r *http.Request) {
	defer UserMeditationViewRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertUserMeditationView(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateUserMeditationViewWhere(w http.ResponseWriter, r *http.Request) {
	defer UserMeditationViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateUserMeditationViewWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindUserMeditationView(w http.ResponseWriter, r *http.Request) {
	defer UserMeditationViewRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindUserMeditationView(or, and, in, span, limit, column, order)
	types.MarshalInto(m, w)
}
