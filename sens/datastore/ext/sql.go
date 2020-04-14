package ext

const (
	ACTIVITY_DASHBOARD_QUERY = `SELECT count(activity_type), activity_type FROM
		user_session_views
		WHERE timestamp > $1 and user_id = ANY($2) GROUP BY (activity_type)`
)
