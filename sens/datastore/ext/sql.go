package ext

import (
	"github.com/jmoiron/sqlx"
	"github.com/senslabs/alpha/sens/datastore"
)

const (
	ACTIVITY_DASHBOARD_QUERY = `SELECT count(type), type FROM (
		SELECT type, ended_at as timestamp FROM sessions WHERE type = 'Sleep' and user_id IN (:user_ids)
		UNION
		SELECT type, ended_at as timestamp FROM sessions WHERE type = 'Sleep' and user_id IN (:user_ids)
		UNION
		SELECT 'Alert', created_at as timestamp FROM alerts WHERE user_id IN (:user_ids)
		UNION
		SELECT 'Device', active_at as timestamp FROM device_activities WHERE device_id IN (SELECT device_id FROM device_views WHERE user_id IN (:user_ids) and status = 'PAIRED')
  ) WHERE timestamp > :when
  GROUP BY (type)`
)

func GetPrepared(name string) (*sqlx.NamedStmt, error) {
	db := datastore.GetConnection()
	return db.PrepareNamed(name)
}
