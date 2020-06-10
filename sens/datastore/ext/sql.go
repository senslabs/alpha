package ext

const (
	ACTIVITY_DASHBOARD_QUERY = `SELECT 1 as Days, count(activity_type) AS Count, activity_type AS ActivityType FROM org_activity_views WHERE timestamp >= $1 and org_id = $2 GROUP BY (activity_type)
	UNION
	SELECT 7 as Days, count(activity_type) AS Count, activity_type AS ActivityType FROM org_activity_views WHERE timestamp >= $3 and org_id = $4 GROUP BY (activity_type)`

	DELETE_USER_QUERY = `DELETE FROM auths WHERE auth_id = (select auth_id from auths where user_id=$1)`

	AVG_RECORD_VALUE_QUERY = `SELECT UserId, Key, Avg FROM
	(SELECT user_id::text AS UserId, key AS Key, avg(value) AS Avg FROM session_records WHERE user_id = ANY(SELECT user_id FROM users WHERE org_id=$1) AND key = $2 AND timestamp > $3 GROUP BY user_id, key) t
	WHERE Avg <= $4 OR Avg >= $5`

	ORG_SESSION_QUERY = `select t.user_id, t.session_id, t.wakeup_time, t.records, t.properties from
	(
		select distinct on (user_id) user_id, session_id, wakeup_time, records, properties
		from session_views where user_id = ANY($1)
		order by user_id, wakeup_time desc
	) t order by wakeup_time`

	SESSION_RECORD_QUERY = `select key, json_agg(timestamp) as timestamps, json_agg(value) as values from session_records sr where user_id = $1 and timestamp >= $2 and timestamp <= $3 and key = ANY($4) group by key`
)
