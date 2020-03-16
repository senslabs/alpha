CREATE VIEW org_details as
SELECT
	a.id as auth_id,
	a.email,
	a.mobile,
	a.social,
	a.first_name,
	a.last_name,
    --ORGS
	o.id as org_id
FROM
	auths a
JOIN org_auths oa on
	oa.auth_id = a.id
JOIN "orgs" o on
	o.id = oa.org_id;

CREATE VIEW op_details as
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
JOIN op_auths oa on
	oa.auth_id = a.id
JOIN "ops" o on
	o.id = oa.op_id;

CREATE VIEW user_details as
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
JOIN user_auths ua on
	ua.auth_id = a.id
JOIN "users" u on
	u.id = ua.user_id;