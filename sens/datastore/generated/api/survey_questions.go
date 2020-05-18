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

func SurveyQuestionMain(r *mux.Router) {
	r.HandleFunc("/api/survey-questions/create", CreateSurveyQuestion)
	r.HandleFunc("/api/survey-questions/batch/create", BatchCreateSurveyQuestion)
	
	r.HandleFunc("/api/survey-questions/{id}/update", UpdateSurveyQuestion)
	r.HandleFunc("/api/survey-questions/{id}/get", GetSurveyQuestion)
    
	r.HandleFunc("/api/survey-questions/update", UpdateSurveyQuestionWhere)
	r.HandleFunc("/api/survey-questions/find", FindSurveyQuestion).Queries("limit", "{limit}")
	r.HandleFunc("/api/survey-questions/delete", DeleteSurveyQuestion)
}

func SurveyQuestionRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateSurveyQuestion(w http.ResponseWriter, r *http.Request) {
	defer SurveyQuestionRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.InsertSurveyQuestion(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateSurveyQuestion(w http.ResponseWriter, r *http.Request) {
	defer SurveyQuestionRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsertSurveyQuestion(data)
	w.WriteHeader(http.StatusOK)
}


func UpdateSurveyQuestion(w http.ResponseWriter, r *http.Request) {
	defer SurveyQuestionRecovery(w)
	vars := mux.Vars(r)
	id := vars["id"]
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateSurveyQuestion(id, data)
	w.WriteHeader(http.StatusOK)
}

func GetSurveyQuestion(w http.ResponseWriter, r *http.Request) {
	defer SurveyQuestionRecovery(w)
	vars := mux.Vars(r)
	id := vars["id"]
	m := fn.SelectSurveyQuestion(id)
	types.MarshalInto(m, w)
}


func UpdateSurveyQuestionWhere(w http.ResponseWriter, r *http.Request) {
	defer SurveyQuestionRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.UpdateSurveyQuestionWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindSurveyQuestion(w http.ResponseWriter, r *http.Request) {
	defer SurveyQuestionRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindSurveyQuestion(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindSurveyQuestion: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteSurveyQuestion(w http.ResponseWriter, r *http.Request) {
	defer SurveyQuestionRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteSurveyQuestion(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteSurveyQuestion: %d", n)
	types.MarshalInto(n, w)
}