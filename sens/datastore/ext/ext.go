package ext

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"github.com/senslabs/alpha/sens/datastore"
	"github.com/senslabs/alpha/sens/errors"
	"github.com/senslabs/alpha/sens/httpclient"
	"github.com/senslabs/alpha/sens/logger"
	"github.com/senslabs/alpha/sens/types"
)

func ExtMain(r *mux.Router) {
	s := r.PathPrefix("/api/ext").Subrouter()
	s.HandleFunc("/activities/get", GetOrgActivites).Queries("days", "{days:[0-9]+}")
}

// type Activity struct {
// 	Count        int    `db:"count"`
// 	ActivityType string `db:"activity_type"`
// }

func GetOrgActivites(w http.ResponseWriter, r *http.Request) {
	in := r.URL.Query().Get("in")
	days := r.URL.Query().Get("days")
	userIds := strings.Split(in, "^")
	if len(userIds) < 2 {
		logger.Error("Too less number of arguments")
		httpclient.WriteError(w, http.StatusInternalServerError, errors.New(errors.GO_ERROR, "Too less number of arguments"))
	} else if duration, err := strconv.Atoi(days); err != nil {
		logger.Error("Too less number of arguments")
		httpclient.WriteError(w, http.StatusInternalServerError, errors.New(errors.GO_ERROR, "Too less number of arguments"))
	} else {
		when := time.Now().Add(-time.Duration(duration*24) * time.Hour).Unix()
		db := datastore.GetConnection()
		stmt, err := db.Prepare(ACTIVITY_DASHBOARD_QUERY)
		errors.Pie(err)
		rows, err := stmt.Query(when, pq.Array(userIds[1:]))
		errors.Pie(err)
		result := datastore.RowsToMapReflect(rows)
		types.MarshalInto(result, w)
	}
}
