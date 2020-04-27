package ext

const (
	ACTIVITY_DASHBOARD_QUERY = `SELECT count(activity_type) as Count, activity_type as ActivityType FROM
	    org_activity_views
		WHERE timestamp > $1 and org_id = $2 GROUP BY (activity_type)`
	DELETE_USER_QUERY = `DELETE FROM auths WHERE auth_id = (select auth_id from auths where user_id=$1)`
)
