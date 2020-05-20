package ext

const (
	ACTIVITY_DASHBOARD_QUERY = `SELECT count(activity_type) AS Count, activity_type AS ActivityType FROM
	    org_activity_views
		WHERE timestamp > $1 and org_id = $2 GROUP BY (activity_type)`

	DELETE_USER_QUERY = `DELETE FROM auths WHERE auth_id = (select auth_id from auths where user_id=$1)`

	AVG_RECORD_VALUE_QUERY = `SELECT UserId, Key, Avg FROM
	(SELECT user_id::text AS UserId, key AS Key, avg(value) AS Avg FROM session_records WHERE user_id = ANY(SELECT user_id FROM users WHERE org_id=$1) AND key = $2 AND timestamp > $3 GROUP BY user_id, key) t
	WHERE Avg <= $4 OR Avg >= $5`
)
