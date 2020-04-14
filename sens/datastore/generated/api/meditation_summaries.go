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

func MeditationSummarieMain(r *mux.Router) {
	r.HandleFunc("/api/meditation-summaries/create", CreateMeditationSummarie)
	r.HandleFunc("/api/meditation-summaries/batch/create", BatchCreateMeditationSummarie)
	
	r.HandleFunc("/api/meditation-summaries/update", UpdateMeditationSummarieWhere)
	r.HandleFunc("/api/meditation-summaries/find", FindMeditationSummarie).Queries("limit", "{limit}")
}

func MeditationSummarieRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateMeditationSummarie(w http.ResponseWriter, r *http.Request) {
	defer MeditationSummarieRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertMeditationSummarie(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateMeditationSummarie(w http.ResponseWriter, r *http.Request) {
	defer MeditationSummarieRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertMeditationSummarie(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateMeditationSummarieWhere(w http.ResponseWriter, r *http.Request) {
	defer MeditationSummarieRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateMeditationSummarieWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindMeditationSummarie(w http.ResponseWriter, r *http.Request) {
	defer MeditationSummarieRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindMeditationSummarie(or, and, in, span, limit, column, order)
	types.MarshalInto(m, w)
}
