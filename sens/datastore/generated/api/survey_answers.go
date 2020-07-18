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

func SurveyAnswerMain(r *mux.Router) {
	r.HandleFunc("/api/survey-answers/create", CreateSurveyAnswer)
	r.HandleFunc("/api/survey-answers/batch/create", BatchCreateSurveyAnswer)
	
	r.HandleFunc("/api/survey-answers/{id}/update", UpdateSurveyAnswer)
	r.HandleFunc("/api/survey-answers/{id}/get", GetSurveyAnswer)
    
	r.HandleFunc("/api/survey-answers/update", UpdateSurveyAnswerWhere)
	r.HandleFunc("/api/survey-answers/find", FindSurveyAnswer).Queries("limit", "{limit}")
	r.HandleFunc("/api/survey-answers/delete", DeleteSurveyAnswer)
}

func SurveyAnswerRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateSurveyAnswer(w http.ResponseWriter, r *http.Request) {
	defer SurveyAnswerRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	id := fn.InsertSurveyAnswer(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateSurveyAnswer(w http.ResponseWriter, r *http.Request) {
	defer SurveyAnswerRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	fn.BatchInsertSurveyAnswer(data)
	w.WriteHeader(http.StatusOK)
}


func UpdateSurveyAnswer(w http.ResponseWriter, r *http.Request) {
	defer SurveyAnswerRecovery(w)
	vars := mux.Vars(r)
	id := vars["id"]
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateSurveyAnswer(id, data)
	w.WriteHeader(http.StatusOK)
}

func GetSurveyAnswer(w http.ResponseWriter, r *http.Request) {
	defer SurveyAnswerRecovery(w)
	vars := mux.Vars(r)
	id := vars["id"]
	m := fn.SelectSurveyAnswer(id)
	types.MarshalInto(m, w)
}


func UpdateSurveyAnswerWhere(w http.ResponseWriter, r *http.Request) {
	defer SurveyAnswerRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateSurveyAnswerWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindSurveyAnswer(w http.ResponseWriter, r *http.Request) {
	defer SurveyAnswerRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindSurveyAnswer(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindSurveyAnswer: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteSurveyAnswer(w http.ResponseWriter, r *http.Request) {
	defer SurveyAnswerRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteSurveyAnswer(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteSurveyAnswer: %d", n)
	types.MarshalInto(n, w)
}