CREATE VIEW org_detail_views as
SELECT
	a.id as auth_id,
	a.email,
	a.mobile,
	a.social,
	a.first_name,
	a.last_name,
    --ORGS
	o.id as id,
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
	o.id as id
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
    -- USERS
	u.id as id,
    -- DEVICES
    d.device_id,
	d.status,
    d.name,
    d.created_at
FROM
	auths a
    -- WHERE d.user_id=u.id
JOIN "users" u on
	a.id = u.auth_id
LEFT JOIN devices d on
    d.user_id = u.id;


CREATE VIEW device_views AS
SELECT
  device_id, name, org_id, user_id, created_at, status
FROM (
    SELECT
      DISTINCT ON(device_id) device_id, name, org_id, user_id, created_at, status
    from devices
    ORDER BY
      device_id,
      created_at desc
  ) t
ORDER BY
  created_at DESC;