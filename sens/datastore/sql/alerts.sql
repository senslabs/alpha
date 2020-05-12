CREATE TABLE "alert_rules" (
  "alert_rule_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "org_id" uuid,
  "alert_name" string UNIQUE,
  "key" string UNIQUE,
  "duration" int,
  "enabled" boolean DEFAULT true,
  "created_at" int DEFAULT (now()::int),
  "updated_at" int,
  "upper_limit" float,
  "lower_limit" float,
  "valid_from" string,
  "valid_till" string,
  "snoozed_at" int,
  "snoozed_for" int,
  "default_snooze" int
);

CREATE TABLE "alert_escalations" (
  "alert_escalation_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "alert_rule_id" uuid,
  "medium" string,
  "medium_value" string,
  "created_at" int DEFAULT (now()::int),
  "timeout" int
);

CREATE TABLE "alerts" (
  "alert_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "user_id" uuid,
  "alert_rule_id" uuid NOT NULL,
  "created_at" int DEFAULT (now()::int),
  "updated_at" int DEFAULT (now()::int),
  "alert_name" string,
  "valid" boolean DEFAULT true,
  "status" string,
  "remarks" string,
  "triggered_level" int
);

ALTER TABLE "alert_escalations" ADD FOREIGN KEY ("alert_rule_id") REFERENCES "alert_rules" ("alert_rule_id");
