CREATE TABLE "auths" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "email" text UNIQUE,
  "mobile" text UNIQUE NOT NULL,
  "social" text UNIQUE,
  "first_name" text NOT NULL,
  "last_name" text NOT NULL,
  "created_at" int DEFAULT (now()::int),
  "updated_at" int
);

CREATE TABLE "orgs" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "auth_id" uuid,
  "name" text UNIQUE NOT NULL,
  "created_at" int DEFAULT (now()::int),
  "updated_at" int
);

CREATE TABLE "ops" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "auth_id" uuid,
  "created_at" int DEFAULT (now()::int),
  "updated_at" int,
  "status" text
);

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "auth_id" uuid,
  "created_at" int DEFAULT (now()::int),
  "updated_at" int,
  "status" text
);

CREATE TABLE "endpoints" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "category" text DEFAULT 'ALL',
  "path" text,
  "secure" boolean DEFAULT true
);

CREATE TABLE "devices" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "device_id" uuid,
  "name" text,
  "org_id" uuid,
  "user_id" uuid,
  "created_at" int DEFAULT (now()::int),
  "status" text,
  "properties" jsonb
);

CREATE TABLE "org_ops" (
  "org_id" uuid,
  "op_id" uuid,
  PRIMARY KEY ("org_id", "op_id")
);

CREATE TABLE "org_users" (
  "org_id" uuid,
  "user_id" uuid,
  "category" text DEFAULT 'ALL',
  PRIMARY KEY ("org_id", "user_id")
);

CREATE TABLE "op_user_categories" (
  "op_id" uuid,
  "user_category" text DEFAULT 'ALL',
  PRIMARY KEY ("op_id", "user_category")
);

CREATE TABLE "op_users" (
  "op_id" uuid,
  "user_id" uuid,
  "access" boolean DEFAULT false,
  PRIMARY KEY ("op_id", "user_id")
);

CREATE TABLE "org_endpoint_categories" (
  "org_id" uuid,
  "endpoint_category" text DEFAULT 'ALL',
  PRIMARY KEY ("org_id", "endpoint_category")
);

CREATE TABLE "org_endpoints" (
  "org_id" uuid,
  "endpoint_id" uuid,
  "access" boolean DEFAULT false,
  PRIMARY KEY ("org_id", "endpoint_id")
);

CREATE TABLE "op_endpoint_categories" (
  "op_id" uuid,
  "endpoint_category" text DEFAULT 'ALL',
  PRIMARY KEY ("op_id", "endpoint_category")
);

CREATE TABLE "op_endpoints" (
  "op_id" uuid,
  "endpoint_id" uuid,
  "access" boolean DEFAULT false,
  PRIMARY KEY ("op_id", "endpoint_id")
);

CREATE TABLE "user_endpoint_categories" (
  "user_id" uuid,
  "endpoint_category" text DEFAULT 'ALL',
  PRIMARY KEY ("user_id", "endpoint_category")
);

CREATE TABLE "user_endpoints" (
  "user_id" uuid,
  "endpoint_id" uuid,
  "access" boolean DEFAULT false,
  PRIMARY KEY ("user_id", "endpoint_id")
);

CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid,
  "name" text,
  "type" text,
  "started_at" int,
  "ended_at" int
);

CREATE TABLE "session_events" (
  "user_id" uuid,
  "name" text,
  "started_at" int,
  "ended_at" int,
  "properties" jsonb,
  PRIMARY KEY ("user_id", "name", "started_at")
);

CREATE TABLE "session_records" (
  "user_id" uuid,
  "name" text,
  "timestamp" int,
  "value" float,
  "properties" jsonb,
  PRIMARY KEY ("user_id", "name", "timestamp")
);

CREATE TABLE "session_properties" (
  "session_id" uuid,
  "name" text,
  "value" text
);

ALTER TABLE "orgs" ADD FOREIGN KEY ("auth_id") REFERENCES "auths" ("id");

ALTER TABLE "ops" ADD FOREIGN KEY ("auth_id") REFERENCES "auths" ("id");

ALTER TABLE "users" ADD FOREIGN KEY ("auth_id") REFERENCES "auths" ("id");

ALTER TABLE "devices" ADD FOREIGN KEY ("org_id") REFERENCES "orgs" ("id");

ALTER TABLE "devices" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "org_ops" ADD FOREIGN KEY ("org_id") REFERENCES "orgs" ("id");

ALTER TABLE "org_ops" ADD FOREIGN KEY ("op_id") REFERENCES "ops" ("id");

ALTER TABLE "org_users" ADD FOREIGN KEY ("org_id") REFERENCES "orgs" ("id");

ALTER TABLE "org_users" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "op_user_categories" ADD FOREIGN KEY ("op_id") REFERENCES "ops" ("id");

ALTER TABLE "op_users" ADD FOREIGN KEY ("op_id") REFERENCES "ops" ("id");

ALTER TABLE "op_users" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "org_endpoint_categories" ADD FOREIGN KEY ("org_id") REFERENCES "orgs" ("id");

ALTER TABLE "org_endpoints" ADD FOREIGN KEY ("org_id") REFERENCES "orgs" ("id");

ALTER TABLE "org_endpoints" ADD FOREIGN KEY ("endpoint_id") REFERENCES "endpoints" ("id");

ALTER TABLE "op_endpoint_categories" ADD FOREIGN KEY ("op_id") REFERENCES "ops" ("id");

ALTER TABLE "op_endpoints" ADD FOREIGN KEY ("op_id") REFERENCES "ops" ("id");

ALTER TABLE "op_endpoints" ADD FOREIGN KEY ("endpoint_id") REFERENCES "endpoints" ("id");

ALTER TABLE "user_endpoint_categories" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_endpoints" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_endpoints" ADD FOREIGN KEY ("endpoint_id") REFERENCES "endpoints" ("id");

ALTER TABLE "sessions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "session_events" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "session_records" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "session_properties" ADD FOREIGN KEY ("session_id") REFERENCES "sessions" ("id");

CREATE INDEX ON "devices" ("device_id");

CREATE INDEX ON "sessions" ("ended_at");

CREATE INDEX ON "sessions" ("user_id");
