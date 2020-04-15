package ext

const (
	ACTIVITY_DASHBOARD_QUERY = `SELECT count(activity_type) as Count, activity_type as ActivityType FROM
		user_session_views
		WHERE timestamp > $1 and user_id = ANY($2) GROUP BY (activity_type)`
)
