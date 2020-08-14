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

func NoteMain(r *mux.Router) {
	r.HandleFunc("/api/notes/create", CreateNote)
	r.HandleFunc("/api/notes/batch/create", BatchCreateNote)
	
	r.HandleFunc("/api/notes/update", UpdateNoteWhere)
	r.HandleFunc("/api/notes/find", FindNote).Queries("limit", "{limit}")
	r.HandleFunc("/api/notes/delete", DeleteNote)
}

func NoteRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
	defer NoteRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	id := fn.InsertNote(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateNote(w http.ResponseWriter, r *http.Request) {
	defer NoteRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	fn.BatchUpsertNote(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateNoteWhere(w http.ResponseWriter, r *http.Request) {
	defer NoteRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateNoteWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindNote(w http.ResponseWriter, r *http.Request) {
	defer NoteRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindNote(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindNote: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	defer NoteRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteNote(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteNote: %d", n)
	types.MarshalInto(n, w)
}