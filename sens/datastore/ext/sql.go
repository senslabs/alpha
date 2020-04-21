package ext

const (
	ACTIVITY_DASHBOARD_QUERY = `SELECT count(activity_type) as Count, activity_type as ActivityType FROM
	    org_activity_views
		WHERE timestamp > $1 and org_id = $2 GROUP BY (activity_type)`
)
