package ext

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"github.com/senslabs/alpha/sens/datastore"
	"github.com/senslabs/alpha/sens/datastore/generated/models"
	"github.com/senslabs/alpha/sens/errors"
	"github.com/senslabs/alpha/sens/logger"
	"github.com/senslabs/alpha/sens/types"
)

func ExtMain(r *mux.Router) {
	s := r.PathPrefix("/api/ext").Subrouter()
	s.HandleFunc("/activities/get", GetOrgActivites).Queries("days", "{days:[0-9]+}")
	s.HandleFunc("/users/{id}/delete", DeleteUserAuth)
	s.HandleFunc("/records/avg", GetAvgValues)
	s.HandleFunc("/users/{id}/trends", GetUserTrends).Queries("From", "{From}", "To", "{To}")
	s.HandleFunc("/users/sessions/summary", GetOrgSleepView).Queries("UserId", "{UserId}")
	s.HandleFunc("/session-records/get", GetSessionRecords).Queries("UserId", "{UserId}", "From", "{From}", "To", "{To}", "Key", "{Key}")
}

func GetOrgActivites(w http.ResponseWriter, r *http.Request) {
	and := r.URL.Query().Get("and")
	tokens := strings.Split(and, "^")
	if len(tokens) == 2 {
		orgId := tokens[1]
		// days := r.URL.Query().Get("days")
		// duration, err := strconv.Atoi(days)
		// errors.Pie(err)
		// when := time.Now().Add(-time.Duration(duration*24) * time.Hour).Unix()
		daily := time.Now().Truncate(24 * time.Hour).Unix()
		weekly := time.Now().Truncate(24 * time.Hour).Add(-6 * 24 * time.Hour).Unix()
		db := datastore.GetConnection()
		stmt, err := db.Prepare(ACTIVITY_DASHBOARD_QUERY)
		errors.Pie(err)
		rows, err := stmt.Query(daily, orgId, weekly, orgId)
		logger.Debug(daily, ",", weekly)
		errors.Pie(err)

		var result []map[string]interface{}
		var days interface{}
		var count interface{}
		var activityType interface{}
		for rows.Next() {
			m := map[string]interface{}{}
			rows.Scan(&days, &count, &activityType)
			m["Days"] = days
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
	o := r.URL.Query().Get("orgId")
	k := r.URL.Query().Get("key")
	ll := r.URL.Query().Get("lowerLimit")
	ul := r.URL.Query().Get("upperLimit")

	duration, err := strconv.ParseInt(m, 10, 64)
	errors.Pie(err)
	when := time.Now().Add(-time.Minute * time.Duration(duration)).Unix()
	db := datastore.GetConnection()
	stmt, err := db.Prepare(AVG_RECORD_VALUE_QUERY)
	errors.Pie(err)
	rows, err := stmt.Query(o, k, when, ll, ul)
	errors.Pie(err)

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
	Properties map[string]interface{}
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

	sm := map[string][]sessionViews{}
	for rows.Next() {
		s := sessionViews{}
		var props []byte
		rows.Scan(&s.SessionId, &s.Date, &s.UserId, &props)
		s.Properties = types.UnmarshalMap(props)
		if s.Properties["SleepTime"] == nil || s.Properties["WakeupTime"] == nil {
			continue
		}
		sm[s.Date] = append(sm[s.Date], s)
	}

	i := 1
	query := []string{}
	var values []interface{}
	ph := `SELECT $%d::date, key, min(value), avg(value), max(value) FROM session_records sr WHERE key in ('HeartRate', 'BreathRate', 'Stress') AND value > 0 AND timestamp >= $%d AND timestamp <= $%d AND user_id = $%d GROUP BY key`
	for d, ss := range sm {
		sz := len(ss)
		sort.Slice(ss, func(l int, r int) bool {
			left, err := strconv.ParseInt(ss[l].Properties["SleepTime"].(string), 10, 64)
			errors.Pie(err)
			right, err := strconv.ParseInt(ss[l].Properties["SleepTime"].(string), 10, 64)
			errors.Pie(err)
			return left < right
		})
		query = append(query, fmt.Sprintf(ph, i, i+1, i+2, i+3))
		values = append(values, d, ss[0].Properties["SleepTime"], ss[sz-1].Properties["WakeupTime"], ss[0].UserId)
		i = i + 4
	}

	logger.Debug(strings.Join(query, " UNION "))
	logger.Debug(values...)

	trends := []Trend{}
	logger.Error(strings.Join(query, " UNION "))
	rows, err = db.Query(strings.Join(query, " UNION "), values...)
	errors.Pie(err)
	for rows.Next() {
		t := Trend{}
		rows.Scan(&t.Date, &t.Key, &t.Min, &t.Avg, &t.Max)
		trends = append(trends, t)
	}
	types.MarshalInto(trends, w)
}

func GetOrgSleepView(w http.ResponseWriter, r *http.Request) {
	db := datastore.GetConnection()
	stmt, err := db.Prepare(ORG_SESSION_QUERY)
	errors.Pie(err)
	userIds := r.URL.Query()["UserId"]
	rows, err := stmt.Query(pq.Array(userIds))
	errors.Pie(err)
	result := datastore.RowsToMap(rows, models.GetSessionViewReverseFieldMap(), models.GetSessionViewTypeMap())
	types.MarshalInto(result, w)
}

func GetSessionRecords(w http.ResponseWriter, r *http.Request) {
	db := datastore.GetConnection()
	stmt, err := db.Prepare(SESSION_RECORD_QUERY)
	errors.Pie(err)
	userId := r.URL.Query().Get("UserId")
	from := r.URL.Query().Get("From")
	to := r.URL.Query().Get("To")
	key := r.URL.Query()["Key"]

	fromTime, err := strconv.ParseInt(from, 10, 64)
	errors.Pie(err)
	toTime, err := strconv.ParseInt(to, 10, 64)
	errors.Pie(err)
	rows, err := stmt.Query(userId, fromTime, toTime, pq.Array(key))
	errors.Pie(err)

	result := map[string]map[string]interface{}{}
	for rows.Next() {
		var key string
		var timestamps json.RawMessage
		var values json.RawMessage
		rows.Scan(&key, &timestamps, &values)
		result[key] = map[string]interface{}{
			"Timestamps": timestamps,
			"Values":     values,
		}
	}
	types.MarshalInto(result, w)
}
