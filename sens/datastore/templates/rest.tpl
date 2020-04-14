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

func {{.Model}}Main(r *mux.Router) {
	r.HandleFunc("/api/{{.Path}}/create", Create{{.Model}})
	r.HandleFunc("/api/{{.Path}}/batch/create", BatchCreate{{.Model}})
	{{if .HasId}}
	r.HandleFunc("/api/{{.Path}}/{id}/update", Update{{.Model}})
	r.HandleFunc("/api/{{.Path}}/{id}/get", Get{{.Model}})
    {{end}}
	r.HandleFunc("/api/{{.Path}}/update", Update{{.Model}}Where)
	r.HandleFunc("/api/{{.Path}}/find", Find{{.Model}}).Queries("limit", "{limit}")
}

func {{.Model}}Recovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func Create{{.Model}}(w http.ResponseWriter, r *http.Request) {
	defer {{.Model}}Recovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	id := fn.Insert{{.Model}}(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreate{{.Model}}(w http.ResponseWriter, r *http.Request) {
	defer {{.Model}}Recovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.BatchInsert{{.Model}}(data)
	w.WriteHeader(http.StatusOK)
}

{{if .HasId}}
func Update{{.Model}}(w http.ResponseWriter, r *http.Request) {
	defer {{.Model}}Recovery(w)
	vars := mux.Vars(r)
	id := vars["id"]
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.Update{{.Model}}(id, data)
	w.WriteHeader(http.StatusOK)
}

func Get{{.Model}}(w http.ResponseWriter, r *http.Request) {
	defer {{.Model}}Recovery(w)
	vars := mux.Vars(r)
	id := vars["id"]
	m := fn.Select{{.Model}}(id)
	types.MarshalInto(m, w)
}
{{end}}

func Update{{.Model}}Where(w http.ResponseWriter, r *http.Request) {
	defer {{.Model}}Recovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	fn.Update{{.Model}}Where(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func Find{{.Model}}(w http.ResponseWriter, r *http.Request) {
	defer {{.Model}}Recovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.Find{{.Model}}(or, and, in, span, limit, column, order)
	types.MarshalInto(m, w)
}
