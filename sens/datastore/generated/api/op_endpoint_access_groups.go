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

func OpEndpointAccessGroupMain(r *mux.Router) {
	r.HandleFunc("/api/op-endpoint-access-groups/create", CreateOpEndpointAccessGroup)
	r.HandleFunc("/api/op-endpoint-access-groups/batch/create", BatchCreateOpEndpointAccessGroup)
	
	r.HandleFunc("/api/op-endpoint-access-groups/update", UpdateOpEndpointAccessGroupWhere)
	r.HandleFunc("/api/op-endpoint-access-groups/find", FindOpEndpointAccessGroup).Queries("limit", "{limit}")
	r.HandleFunc("/api/op-endpoint-access-groups/delete", DeleteOpEndpointAccessGroup)
}

func OpEndpointAccessGroupRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateOpEndpointAccessGroup(w http.ResponseWriter, r *http.Request) {
	defer OpEndpointAccessGroupRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	id := fn.InsertOpEndpointAccessGroup(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateOpEndpointAccessGroup(w http.ResponseWriter, r *http.Request) {
	defer OpEndpointAccessGroupRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	fn.BatchUpsertOpEndpointAccessGroup(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateOpEndpointAccessGroupWhere(w http.ResponseWriter, r *http.Request) {
	defer OpEndpointAccessGroupRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateOpEndpointAccessGroupWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindOpEndpointAccessGroup(w http.ResponseWriter, r *http.Request) {
	defer OpEndpointAccessGroupRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindOpEndpointAccessGroup(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindOpEndpointAccessGroup: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteOpEndpointAccessGroup(w http.ResponseWriter, r *http.Request) {
	defer OpEndpointAccessGroupRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteOpEndpointAccessGroup(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteOpEndpointAccessGroup: %d", n)
	types.MarshalInto(n, w)
}