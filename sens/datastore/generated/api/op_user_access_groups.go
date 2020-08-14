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

func OpUserAccessGroupMain(r *mux.Router) {
	r.HandleFunc("/api/op-user-access-groups/create", CreateOpUserAccessGroup)
	r.HandleFunc("/api/op-user-access-groups/batch/create", BatchCreateOpUserAccessGroup)
	
	r.HandleFunc("/api/op-user-access-groups/update", UpdateOpUserAccessGroupWhere)
	r.HandleFunc("/api/op-user-access-groups/find", FindOpUserAccessGroup).Queries("limit", "{limit}")
	r.HandleFunc("/api/op-user-access-groups/delete", DeleteOpUserAccessGroup)
}

func OpUserAccessGroupRecovery(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err := r.(error)
		logger.Error(err)
		httpclient.WriteInternalServerError(w, err)
	}
}

func CreateOpUserAccessGroup(w http.ResponseWriter, r *http.Request) {
	defer OpUserAccessGroupRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	id := fn.InsertOpUserAccessGroup(data)
	errors.Pie(err)
	fmt.Fprint(w, id)
}

func BatchCreateOpUserAccessGroup(w http.ResponseWriter, r *http.Request) {
	defer OpUserAccessGroupRecovery(w)
	data, err := ioutil.ReadAll(r.Body)
	errors.Pie(err)
	defer r.Body.Close()
	fn.BatchUpsertOpUserAccessGroup(data)
	w.WriteHeader(http.StatusOK)
}



func UpdateOpUserAccessGroupWhere(w http.ResponseWriter, r *http.Request) {
	defer OpUserAccessGroupRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	errors.Pie(err)
	fn.UpdateOpUserAccessGroupWhere(or, and, in, span, data)
	w.WriteHeader(http.StatusOK)
}

func FindOpUserAccessGroup(w http.ResponseWriter, r *http.Request) {
	defer OpUserAccessGroupRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")
	limit := values.Get("limit")
	column := values.Get("column")
	order := values.Get("order")

	m := fn.FindOpUserAccessGroup(or, and, in, span, limit, column, order)
	logger.Debugf("RESPONSE of FindOpUserAccessGroup: %#v", m)
	types.MarshalInto(m, w)
}

func DeleteOpUserAccessGroup(w http.ResponseWriter, r *http.Request) {
	defer OpUserAccessGroupRecovery(w)
	values := r.URL.Query()
	span := values["span"]
	or := values["or"]
	and := values["and"]
	in := values.Get("in")

	n := fn.DeleteOpUserAccessGroup(or, and, in, span)
	logger.Debugf("RESPONSE of DeleteOpUserAccessGroup: %d", n)
	types.MarshalInto(n, w)
}