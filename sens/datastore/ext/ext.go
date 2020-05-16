package ext

import (
	"fmt"
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
	s.HandleFunc("/users/{id}/trends", GetUserTrends).Queries("From", "{From}", "To", "{To}")
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

type sessionViews struct {
	SessionId  string
	Date       string
	UserId     string
	Timestamps map[string]interface{}
}

type Trend struct {
	Date string
	Key  string
	Min  float64
	Avg  float64
	Max  float64
}

func GetUserTrends(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["id"]
	from := r.URL.Query().Get("From")
	to := r.URL.Query().Get("To")
	db := datastore.GetConnection()

	rows, err := db.Query("SELECT * FROM user_dated_session_views WHERE user_id=$1 AND date >= $2 AND date <= $3", userId, from, to)
	errors.Pie(err)

	var ss []sessionViews
	for rows.Next() {
		s := sessionViews{}
		var ts []byte
		rows.Scan(&s.SessionId, &s.Date, &s.UserId, &ts)
		s.Timestamps = types.UnmarshalMap(ts)
		ss = append(ss, s)
	}

	i := 1
	query := []string{}
	var values []interface{}
	ph := `SELECT max(timestamp::timestamp::date) AS date, key, min(value), avg(value), max(value) FROM session_records sr WHERE key in ('HeartRate', 'BreathRate') AND value > 0 AND timestamp >= $%d AND timestamp <= $%d AND user_id = $%d GROUP BY key`
	for _, s := range ss {
		query = append(query, fmt.Sprintf(ph, i, i+1, i+2))
		values = append(values, s.Timestamps["SleepTime"], s.Timestamps["WakeupTime"], s.UserId)
		i = i + 3
	}

	trends := []Trend{}
	rows, err = db.Query(strings.Join(query, " UNION "), values...)
	for rows.Next() {
		t := Trend{}
		rows.Scan(&t.Date, &t.Key, &t.Min, &t.Avg, &t.Max)
		trends = append(trends, t)
	}
	types.MarshalInto(trends, w)
}
