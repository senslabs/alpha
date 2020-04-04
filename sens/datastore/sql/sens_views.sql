CREATE VIEW org_detail_views as
SELECT
	a.id as auth_id,
	a.email,
	a.mobile,
	a.social,
	a.first_name,
	a.last_name,
    --ORGS
	o.id,
	o.org_name
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
	o.id,
	o.org_id
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
	u.id,
	u.org_id
    -- DEVICES
    -- d.device_id,
	-- d.status,
    -- d.name,
    -- d.created_at
FROM
	auths a
    -- WHERE d.user_id=u.id
JOIN "users" u on
	a.id = u.auth_id;
-- LEFT JOIN devices d on
-- d.user_id = u.id;

CREATE VIEW device_views AS
SELECT
  device_id,
  name,
  org_id,
  user_id,
  created_at,
  status
FROM (
    SELECT
      DISTINCT ON(device_id) device_id,
      name,
      org_id,
      user_id,
      created_at,
      status
    FROM devices
    ORDER BY
      device_id,
      created_at desc
  ) t
ORDER BY
  created_at DESC;

-- The user_session_views is meant to view the activity type and timestamp of a user
CREATE VIEW user_session_views AS 
SELECT type, ended_at AS timestamp, user_id FROM sessions WHERE type = 'Sleep'
UNION
SELECT type, ended_at AS timestamp, user_id FROM sessions WHERE type = 'Meditation'
UNION
SELECT 'Alert', created_at AS timestamp, user_id FROM alerts
UNION
SELECT 'Device', da.active_at AS timestamp, dv.user_id FROM device_activities da JOIN device_views dv ON dv.device_id = da.device_id

CREATE VIEW user_alert_views AS
SELECT
  a.user_id,
  u.org_id,
  au.first_name,
  au.last_name,
  a.created_at,
  a.alert_name,
  a.status,
  a.remarks
FROM alerts a
JOIN users u ON u.id = a.user_id
JOIN auths au ON au.id = u.auth_id;


CREATE VIEW sleep_views AS
SELECT id, user_id, name, type, started_at, ended_at
FROM
(
    SELECT
    DISTINCT ON (user_id) user_id,
    id,
    name,
    type,
    started_at,
    ended_at
    FROM sessions WHERE type = 'Sleep'
    ORDER BY
    user_id,
    ended_at DESC
) t
ORDER BY ended_at DESC;

CREATE VIEW meditation_views AS
SELECT id, user_id, name, type, started_at, ended_at
FROM
(
    SELECT
    DISTINCT ON (user_id) user_id,
    id,
    name,
    type,
    started_at,
    ended_at
    FROM sessions WHERE type = 'Meditation'
    ORDER BY
    user_id,
    ended_at DESC
) t
ORDER BY ended_at DESC;

CREATE VIEW sleep_summaries AS
SELECT
  sv.user_id,
  (sv.ended_at - sv.started_at) AS duration,
  json_build_object(sp.name, sp.value) AS properties,
  sp.session_id
FROM session_properties sp
JOIN sleep_views sv ON sp.session_id = sv.id
WHERE
  sp.name IN (
    'HeartRate',
    'BreathRate',
    'Lastsyncedat',
    'Stress',
    'Score'
  )

CREATE VIEW meditation_summaries AS
SELECT
  mv.user_id,
  (mv.ended_at - mv.started_at) AS duration,
  json_build_object(sp.name, sp.value) AS properties,
  sp.session_id
FROM session_properties sp
JOIN meditation_views mv ON sp.session_id = mv.id
WHERE
  sp.name IN (
    'HeartRate',
    'BreathRate',
    'Lastsyncedat',
    'Stress',
    'Score'
  )

CREATE VIEW user_summary_views AS
SELECT
  count(type),
  type,
  user_id
FROM user_session_views
GROUP BY(type, user_id)

CREATE VIEW user_sleep_views AS
SELECT
  s.user_id,
  (s.ended_at - s.started_at) AS duration,
  json_build_object(sp.name, sp.value) AS properties,
  sp.session_id
FROM session_properties sp
JOIN sessions s ON sp.session_id = s.id
WHERE
  s.type = 'Sleep' AND
  sp.name IN (
    'HeartRate',
    'BreathRate',
    'Lastsyncedat',
    'Stress',
    'Score'
  )

CREATE VIEW user_meditation_views AS
SELECT
  s.user_id,
  (s.ended_at - s.started_at) AS duration,
  json_build_object(sp.name, sp.value) AS properties,
  sp.session_id
FROM session_properties sp
JOIN sessions s ON sp.session_id = s.id
WHERE
  s.type = 'Meditation' AND
  sp.name IN (
    'HeartRate',
    'BreathRate',
    'Lastsyncedat',
    'Stress',
    'Score'
  )
  