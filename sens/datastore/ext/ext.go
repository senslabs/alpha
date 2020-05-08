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
	s.HandleFunc("/users/{id}/delete", DeleteUserAuth)
	s.HandleFunc("/records/avg", GetAvgValues)
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

func GetAvgValues(w http.ResponseWriter, r *http.Request) {
	m := r.URL.Query().Get("minutes")
	u := r.URL.Query().Get("userId")
	k := r.URL.Query().Get("key")
	ll := r.URL.Query().Get("lowerLimit")
	ul := r.URL.Query().Get("upperLimit")

	duration, err := strconv.ParseInt(m, 10, 64)
	errors.Pie(err)
	when := time.Now().Add(-time.Minute * time.Duration(duration)).Unix()
	db := datastore.GetConnection()
	stmt, err := db.Prepare(AVG_RECORD_VALUE_QUERY)
	errors.Pie(err)
	rows, err := stmt.Query(u, k, when, ll, ul)

	result := []map[string]interface{}{}
	var userId interface{}
	var key interface{}
	var avg interface{}
	for rows.Next() {
		m := map[string]interface{}{}
		rows.Scan(&userId, &key, &avg)
		m["UserId"] = userId
		m["Key"] = key
		m["Avg"] = avg
		result = append(result, m)
	}
	types.MarshalInto(result, w)
}

func DeleteUserAuth(w http.ResponseWriter, r *http.Request) {
	db := datastore.GetConnection()
	stmt, err := db.Prepare(ACTIVITY_DASHBOARD_QUERY)
	errors.Pie(err)
	userId := mux.Vars(r)["id"]
	_, err = stmt.Exec(userId)
	errors.Pie(err)
	w.WriteHeader(http.StatusOK)
}
