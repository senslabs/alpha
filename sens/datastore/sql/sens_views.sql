CREATE VIEW auth_views AS
SELECT
	a.auth_id,
	a.email,
	a.mobile,
	a.social,
	a.first_name,
	a.last_name,
  a.is_sens
FROM
	auths a;

CREATE VIEW org_views AS
SELECT
	a.auth_id,
	a.email,
	a.mobile,
	a.social,
	a.first_name,
	a.last_name,
    --ORGS
	o.org_id,
	o.org_name
FROM
	auths a
JOIN orgs o on
	a.auth_id = o.auth_id;

CREATE VIEW op_views AS
SELECT
	a.auth_id,
	a.email,
	a.mobile,
	a.social,
	a.first_name,
	a.last_name,
    --OPS
	o.op_id,
	o.org_id
FROM
	auths a
JOIN ops o on
	a.auth_id = o.auth_id;

CREATE VIEW user_views AS
SELECT
	a.auth_id,
	a.email,
	a.mobile,
	a.social,
	a.first_name,
	a.last_name,
	u.user_id,
	u.org_id
FROM
	auths a
    -- WHERE d.user_id=u.id
JOIN "users" u on
	a.auth_id = u.auth_id;
-- LEFT JOIN devices d on
-- d.user_id = u.id;

CREATE VIEW device_views AS
SELECT
  device_id,
  device_name,
  org_id,
  user_id,
  created_at,
  status
FROM (
    SELECT
      DISTINCT ON(device_id) device_id,
      device_name,
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

CREATE VIEW org_alert_views AS
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
JOIN users u ON u.user_id = a.user_id
JOIN auths au ON au.auth_id = u.auth_id;

-- ORG SESSIONS

CREATE VIEW org_session_views AS
SELECT
  s.session_id,
  s.user_id,
  u.org_id,
  s.session_name,
  s.session_type,
  s.started_at,
  s.ended_at
FROM sessions s
JOIN users u ON u.user_id = s.user_id;

CREATE VIEW org_session_info_views AS
SELECT
	osv.user_id,
	osv.org_id,
	osv.session_id,
	osv.session_type,
	osv.session_name,
	osv.started_at,
	osv.ended_at,
	json_object(array_agg(key), array_agg(value)) AS properties
FROM
	org_session_views osv
JOIN session_properties sp ON
	sp.session_id = osv.session_id
GROUP BY osv.user_id,
	osv.org_id,
	osv.session_id,
	osv.session_type,
	osv.session_name,
	osv.started_at,
	osv.ended_at;

CREATE VIEW org_sleep_views AS
SELECT t.session_id, t.user_id, t.org_id, t.session_name, t.session_type, t.started_at, t.ended_at, t.properties
FROM
(
    SELECT
    DISTINCT ON (user_id) user_id,
    session_id,
    user_id,
    org_id,
    session_name,
    session_type,
    started_at,
    ended_at,
    properties
    FROM org_session_info_views WHERE session_type = 'Sleep'
    ORDER BY
    user_id,
    ended_at DESC
) t
ORDER BY ended_at DESC;

CREATE VIEW org_meditation_views AS
SELECT t.session_id, t.user_id, t.org_id, t.session_name, t.session_type, t.started_at, t.ended_at, t.properties
FROM
(
    SELECT
    DISTINCT ON (user_id) user_id,
    session_id,
    user_id,
    org_id,
    session_name,
    session_type,
    started_at,
    ended_at,
    properties
    FROM org_session_info_views WHERE session_type = 'Meditation'
    ORDER BY
    user_id,
    ended_at DESC
) t
ORDER BY ended_at DESC;

-- CREATE VIEW org_sleep_summary_views AS
-- SELECT
--   sv.user_id, sv.org_id,
--   (sv.ended_at - sv.started_at) AS duration,
--   json_build_object(sp.key, sp.value) AS properties,
--   sp.session_id
-- FROM session_properties sp
-- JOIN org_sleep_views sv ON sp.session_id = sv.session_id
-- WHERE
--   sp.key IN (
--     'HeartRate',
--     'BreathRate',
--     'LastSyncedAt',
--     'Stress',
--     'Score'
--   );

-- CREATE VIEW org_meditation_summary_views AS
-- SELECT
--   mv.user_id, mv.org_id,
--   (mv.ended_at - mv.started_at) AS duration,
--   json_build_object(sp.key, sp.value) AS properties,
--   sp.session_id
-- FROM session_properties sp
-- JOIN org_meditation_views mv ON sp.session_id = mv.session_id
-- WHERE
--   sp.key IN (
--     'HeartRate',
--     'BreathRate',
--     'LastSyncedAt',
--     'Stress',
--     'Score'
--   );

-- ACTIVITIES
CREATE VIEW org_activity_views AS 
SELECT t.activity_type, t.timestamp, t.user_id, u.org_id FROM
(
	SELECT session_type AS activity_type, ended_at AS timestamp, user_id FROM sessions WHERE session_type = 'Sleep'
	UNION
	SELECT session_type AS activity_type, ended_at AS timestamp, user_id FROM sessions WHERE session_type = 'Meditation'
	UNION
	SELECT 'Alert' AS activity_type, a.created_at AS timestamp, a.user_id FROM alerts a JOIN users u on u.user_id  = a.user_id
	UNION
	SELECT 'Device' AS activity_type, da.active_at AS timestamp, dv.user_id FROM device_activities da JOIN device_views dv ON dv.device_id = da.device_id
) t JOIN users u ON t.user_id = u.user_id;

CREATE VIEW org_activity_summary_views AS
SELECT
  count(activity_type),
  activity_type,
  user_id,
  org_id
FROM org_activity_views
GROUP BY(activity_type, user_id, org_id); 

CREATE VIEW org_quarter_usage_views AS
SELECT
  count(activity_type),
  activity_type,
  org_id,
  timestamp::timestamp::date AS date
FROM org_activity_views
WHERE
  timestamp::timestamp > (CURRENT_DATE - INTERVAL '90 days')
GROUP BY
  date,
  activity_type,
  org_id;

-- USER SESSIONS
CREATE VIEW org_session_record_views AS
SELECT
  sr.user_id,
  u.org_id,
  s.session_id,
  s.session_type,
  s.started_at,
  s.ended_at,
  sr.key,
  sr.timestamp,
  sr.value,
  sr.properties
FROM session_records sr
JOIN sessions s ON s.user_id = sr.user_id
JOIN users u ON u.user_id  = s.user_id 
WHERE sr.timestamp >= s.started_at AND sr.timestamp <= s.ended_at;

CREATE VIEW org_session_detail_views AS
SELECT
  user_id,
  org_id,
  session_id,
  session_type,
  started_at,
  ended_at,
  key,
  json_agg(timestamp) as timestamps,
  json_agg(value) as values,
  min(value),
  max(value),
  avg(value)
FROM org_session_record_views
WHERE key IN ('HeartRate', 'BreathRate', 'Stage', 'Recovery', 'Stress')
GROUP BY
  user_id,
  org_id,
  session_id,
  session_type,
  started_at,
  ended_at,
  key;


CREATE VIEW org_session_event_views AS
SELECT
  se.user_id,
  u.org_id,
  s.session_id,
  s.session_type,
  se.key,
  se.started_at as event_started_at,
  se.ended_at as event_ended_at,
  se.properties
FROM session_events se
JOIN sessions s ON s.user_id = se.user_id
JOIN users u ON u.user_id  = s.user_id 
WHERE se.started_at >= s.started_at AND se.started_at <= s.ended_at;

CREATE VIEW org_session_event_detail_views AS
SELECT
  user_id,
  org_id,
  session_id,
  session_type,
  json_agg(event_started_at) AS event_started_at,
  json_agg(event_ended_at) AS event_ended_at,
  key
FROM org_session_event_views
GROUP BY
  user_id,
  org_id,
  session_id,
  session_type,
  key;