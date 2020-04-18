CREATE TABLE "auths" (
  "auth_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "email" text UNIQUE,
  "mobile" text UNIQUE NOT NULL,
  "social" text UNIQUE,
  "first_name" text NOT NULL,
  "last_name" text NOT NULL,
  "created_at" int DEFAULT (now()::int),
  "updated_at" int,
  "is_sens" bool
);

CREATE TABLE "orgs" (
  "org_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "auth_id" uuid,
  "org_name" text UNIQUE NOT NULL,
  "created_at" int DEFAULT (now()::int),
  "updated_at" int
);

CREATE TABLE "ops" (
  "op_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "auth_id" uuid,
  "org_id" uuid,
  "created_at" int DEFAULT (now()::int),
  "updated_at" int,
  "status" text
);

CREATE TABLE "users" (
  "user_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "auth_id" uuid,
  "org_id" uuid,
  "access_group" text DEFAULT 'DEFAULT',
  "created_at" int DEFAULT (now()::int),
  "updated_at" int,
  "age" int,
  "status" text
);

CREATE TABLE "survey_questions" (
  "survey_question_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "question" text
);

CREATE TABLE "survey_answers" (
  "survey_answer_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "survey_question_id" uuid,
  "answer" text
);

CREATE TABLE "user_settings" (
  "user_setting_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "user_id" uuid,
  "created_at" int DEFAULT (now()::int),
  "key" text NOT NULL,
  "value" text NOT NULL
);

CREATE TABLE "user_properties" (
  "user_id" uuid,
  "key" text NOT NULL,
  "value" text NOT NULL,
  PRIMARY KEY ("user_id", "key")
);

CREATE TABLE "api_keys" (
  "api_key_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "org_id" uuid,
  "key_name" text UNIQUE NOT NULL,
  "description" text,
  "key" text UNIQUE NOT NULL
);

CREATE TABLE "op_user_access_groups" (
  "op_id" uuid,
  "access_group" text DEFAULT 'DEFAULT',
  PRIMARY KEY ("op_id", "access_group")
);

CREATE TABLE "op_users" (
  "op_id" uuid,
  "user_id" uuid,
  "access" boolean DEFAULT false,
  PRIMARY KEY ("op_id", "user_id")
);

CREATE TABLE "endpoints" (
  "endpoint_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "access_group" text DEFAULT 'DEFAULT',
  "path" text,
  "secure" boolean DEFAULT true
);

CREATE TABLE "org_endpoint_access_groups" (
  "org_id" uuid,
  "access_group" text DEFAULT 'DEFAULT',
  PRIMARY KEY ("org_id", "access_group")
);

CREATE TABLE "org_endpoints" (
  "org_id" uuid,
  "endpoint_id" uuid,
  "access" boolean DEFAULT false,
  PRIMARY KEY ("org_id", "endpoint_id")
);

CREATE TABLE "op_endpoint_access_groups" (
  "op_id" uuid,
  "access_group" text DEFAULT 'DEFAULT',
  PRIMARY KEY ("op_id", "access_group")
);

CREATE TABLE "op_endpoints" (
  "op_id" uuid,
  "endpoint_id" uuid,
  "access" boolean DEFAULT false,
  PRIMARY KEY ("op_id", "endpoint_id")
);

CREATE TABLE "user_endpoint_access_groups" (
  "user_id" uuid,
  "endpoint_category" text DEFAULT 'DEFAULT',
  PRIMARY KEY ("user_id", "endpoint_category")
);

CREATE TABLE "user_endpoints" (
  "user_id" uuid,
  "endpoint_id" uuid,
  "access" boolean DEFAULT false,
  PRIMARY KEY ("user_id", "endpoint_id")
);

CREATE TABLE "devices" (
  "device_id" uuid,
  "created_at" int DEFAULT (now()::int),
  "device_name" text,
  "org_id" uuid,
  "user_id" uuid,
  "status" text,
  PRIMARY KEY ("device_id", "created_at")
);

CREATE TABLE "device_activities" (
  "device_id" uuid,
  "activity_type" text,
  "active_at" int DEFAULT (now()::int),
  PRIMARY KEY ("device_id", "activity_type")
);

CREATE TABLE "device_properties" (
  "device_id" uuid NOT NULL,
  "key" text NOT NULL,
  "value" text,
  PRIMARY KEY ("device_id", "key")
);

CREATE TABLE "alerts" (
  "alert_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "user_id" uuid,
  "created_at" int,
  "alert_name" text,
  "status" text,
  "remarks" text
);

CREATE TABLE "sessions" (
  "session_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "user_id" uuid,
  "session_name" text,
  "session_type" text,
  "started_at" int,
  "ended_at" int
);

CREATE TABLE "session_settings" (
  "session_setting_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "user_id" uuid,
  "key" text,
  "created_at" int,
  "session_type" text,
  "value" text
);

CREATE TABLE "vital_baselines" (
  "vital_baseline_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "user_id" uuid,
  "key" text NOT NULL,
  "created_at" int DEFAULT (now()::int),
  "lower_limit" int NOT NULL,
  "upper_limit" int NOT NULL
);

CREATE TABLE "session_events" (
  "user_id" uuid,
  "key" text,
  "started_at" int NOT NULL,
  "ended_at" int,
  "properties" jsonb,
  PRIMARY KEY ("user_id", "key", "started_at")
);

CREATE TABLE "session_records" (
  "user_id" uuid,
  "key" text,
  "timestamp" int,
  "value" float,
  "properties" jsonb,
  PRIMARY KEY ("user_id", "key", "timestamp")
);

CREATE TABLE "session_properties" (
  "session_id" uuid,
  "key" text,
  "value" text,
  PRIMARY KEY ("session_id", "key")
);

ALTER TABLE "orgs" ADD FOREIGN KEY ("auth_id") REFERENCES "auths" ("auth_id");

ALTER TABLE "ops" ADD FOREIGN KEY ("auth_id") REFERENCES "auths" ("auth_id");

ALTER TABLE "ops" ADD FOREIGN KEY ("org_id") REFERENCES "orgs" ("org_id");

ALTER TABLE "users" ADD FOREIGN KEY ("auth_id") REFERENCES "auths" ("auth_id");

ALTER TABLE "users" ADD FOREIGN KEY ("org_id") REFERENCES "orgs" ("org_id");

ALTER TABLE "survey_answers" ADD FOREIGN KEY ("survey_question_id") REFERENCES "survey_questions" ("survey_question_id");

ALTER TABLE "user_settings" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "user_properties" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "api_keys" ADD FOREIGN KEY ("org_id") REFERENCES "orgs" ("org_id");

ALTER TABLE "op_user_access_groups" ADD FOREIGN KEY ("op_id") REFERENCES "ops" ("op_id");

ALTER TABLE "op_users" ADD FOREIGN KEY ("op_id") REFERENCES "ops" ("op_id");

ALTER TABLE "op_users" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "org_endpoint_access_groups" ADD FOREIGN KEY ("org_id") REFERENCES "orgs" ("org_id");

ALTER TABLE "org_endpoints" ADD FOREIGN KEY ("org_id") REFERENCES "orgs" ("org_id");

ALTER TABLE "org_endpoints" ADD FOREIGN KEY ("endpoint_id") REFERENCES "endpoints" ("endpoint_id");

ALTER TABLE "op_endpoint_access_groups" ADD FOREIGN KEY ("op_id") REFERENCES "ops" ("op_id");

ALTER TABLE "op_endpoints" ADD FOREIGN KEY ("op_id") REFERENCES "ops" ("op_id");

ALTER TABLE "op_endpoints" ADD FOREIGN KEY ("endpoint_id") REFERENCES "endpoints" ("endpoint_id");

ALTER TABLE "user_endpoint_access_groups" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "user_endpoints" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "user_endpoints" ADD FOREIGN KEY ("endpoint_id") REFERENCES "endpoints" ("endpoint_id");

ALTER TABLE "devices" ADD FOREIGN KEY ("org_id") REFERENCES "orgs" ("org_id");

ALTER TABLE "devices" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "alerts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "sessions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "session_settings" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "vital_baselines" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "session_events" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "session_records" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "session_properties" ADD FOREIGN KEY ("session_id") REFERENCES "sessions" ("session_id");

CREATE INDEX ON "sessions" ("ended_at");

CREATE INDEX ON "sessions" ("user_id");
