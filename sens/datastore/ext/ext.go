package ext

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/senslabs/alpha/sens/datastore"
	"github.com/senslabs/alpha/sens/errors"
	"github.com/senslabs/alpha/sens/types"
)

func ExtMain(r *mux.Router) {
	s := r.PathPrefix("/api/ext").Subrouter()
	s.HandleFunc("/activities/get", GetOrgActivites).Queries("days", "{days:[0-9]+}")
}

func GetOrgActivites(w http.ResponseWriter, r *http.Request) {
	and := r.URL.Query().Get("and")
	tokens := strings.Split(and, "^")
	if len(tokens) == 2 {
		orgId := tokens[1]
		days := r.URL.Query().Get("days")
		duration, err := strconv.Atoi(days)
		errors.Pie(err)
		when := time.Now().Add(-time.Duration(duration*24) * time.Hour).Unix()
		db := datastore.GetConnection()
		stmt, err := db.Prepare(ACTIVITY_DASHBOARD_QUERY)
		errors.Pie(err)
		rows, err := stmt.Query(when, orgId)
		errors.Pie(err)

		var result []map[string]interface{}
		var count interface{}
		var activityType interface{}
		for rows.Next() {
			m := map[string]interface{}{}
			rows.Scan(&count, &activityType)
			m["Count"] = count
			m["ActivityType"] = activityType
			result = append(result, m)
		}
		types.MarshalInto(result, w)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
