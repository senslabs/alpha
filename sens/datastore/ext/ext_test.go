package ext

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/senslabs/alpha/sens/datastore"
	"github.com/senslabs/alpha/sens/errors"
	"github.com/senslabs/alpha/sens/logger"
	"github.com/senslabs/alpha/sens/types"
)

func TestA(t *testing.T) {
	logger.InitLogger("TestA")
	userId := "6f129b1c-43a6-4771-86f6-1749bfe1a5af"
	db := datastore.GetConnection()
	rows, err := db.Query("SELECT * FROM user_dated_session_views WHERE user_id=$1", userId)
	errors.Pie(err)

	var ss []sessionViews
	for rows.Next() {
		s := sessionViews{}
		var ts []byte
		rows.Scan(&s.SessionId, &s.Date, &s.UserId, &ts)
		s.Properties = types.UnmarshalMap(ts)
		ss = append(ss, s)
	}

	i := 1
	query := []string{}
	var values []interface{}
	ph := "SELECT max(timestamp::timestamp::date) AS date, key, min(value), avg(value), max(value) FROM session_records sr WHERE key in ('HeartRate', 'BreathRate') AND value > 0 AND timestamp >= $%d AND timestamp <= $%d AND user_id = $%d GROUP BY key"
	for _, s := range ss {
		query = append(query, fmt.Sprintf(ph, i, i+1, i+2))
		values = append(values, s.Properties["SleepTime"], s.Properties["WakeupTime"], s.UserId)
		i = i + 3
	}

	var trends []Trend
	rows, err = db.Query(strings.Join(query, " UNION "), values...)
	for rows.Next() {
		t := Trend{}
		rows.Scan(&t.Date, &t.Key, &t.Min, &t.Avg, &t.Max)
		trends = append(trends, t)
		fmt.Println(trends)
	}
	types.MarshalInto(trends, os.Stdout)
}
