CREATE VIEW org_detail_views as
SELECT
	a.id as auth_id,
	a.email,
	a.mobile,
	a.social,
	a.first_name,
	a.last_name,
    --ORGS
	o.id as org_id,
	o.name as org_name
FROM
	auths a
JOIN orgs o on
	a.id = o.auth_id;

CREATE VIEW op_detail_views as
SELECT
	a.id as auth_id,
	a.email,
	a.mobile,
	a.social,
	a.first_name,
	a.last_name,
    --OPS
	o.id as op_id
FROM
	auths a
JOIN ops o on
	a.id = o.auth_id;

CREATE VIEW user_detail_views as
SELECT
	a.id as auth_id,
	a.email,
	a.mobile,
	a.social,
	a.first_name,
	a.last_name,
    --USERS
	u.id as user_id
FROM
	auths a
JOIN "users" u on
	a.id = u.auth_id;