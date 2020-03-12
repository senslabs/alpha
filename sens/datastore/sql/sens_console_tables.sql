CREATE TABLE "auths" (
  "id" uuid PRIMARY KEY,
  "email" text UNIQUE NOT NULL,
  "mobile" text UNIQUE NOT NULL,
  "social" text UNIQUE NOT NULL,
  "first_name" text,
  "last_name" text,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "orgs" (
  "id" uuid PRIMARY KEY,
  "name" text UNIQUE,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "ops" (
  "id" uuid PRIMARY KEY,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "endpoints" (
  "id" uuid PRIMARY KEY,
  "group" text,
  "path" text,
  "secure" boolean,
  "next_endpoint" text
);

CREATE TABLE "org_auths" (
  "org_id" uuid,
  "auth_id" uuid,
  PRIMARY KEY ("org_id", "auth_id")
);

CREATE TABLE "op_auths" (
  "op_id" uuid,
  "auth_id" uuid,
  PRIMARY KEY ("op_id", "auth_id")
);

CREATE TABLE "user_auths" (
  "user_id" uuid,
  "auth_id" uuid,
  PRIMARY KEY ("user_id", "auth_id")
);

CREATE TABLE "org_ops" (
  "org_id" uuid,
  "op_id" uuid,
  PRIMARY KEY ("org_id", "op_id")
);

CREATE TABLE "org_users" (
  "org_id" uuid,
  "user_id" uuid,
  PRIMARY KEY ("org_id", "user_id")
);

CREATE TABLE "op_users" (
  "op_id" uuid,
  "user_id" uuid,
  PRIMARY KEY ("op_id", "user_id")
);

CREATE TABLE "org_endpoints" (
  "org_id" uuid,
  "endpoint_id" uuid,
  PRIMARY KEY ("org_id", "endpoint_id")
);

CREATE TABLE "op_endpoints" (
  "op_id" uuid,
  "endpoint_id" uuid,
  PRIMARY KEY ("op_id", "endpoint_id")
);

CREATE TABLE "user_endpoints" (
  "user_id" uuid,
  "endpoint_id" uuid,
  PRIMARY KEY ("user_id", "endpoint_id")
);

ALTER TABLE "org_auths" ADD FOREIGN KEY ("org_id") REFERENCES "orgs" ("id");

ALTER TABLE "org_auths" ADD FOREIGN KEY ("auth_id") REFERENCES "auths" ("id");

ALTER TABLE "op_auths" ADD FOREIGN KEY ("op_id") REFERENCES "ops" ("id");

ALTER TABLE "op_auths" ADD FOREIGN KEY ("auth_id") REFERENCES "auths" ("id");

ALTER TABLE "user_auths" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_auths" ADD FOREIGN KEY ("auth_id") REFERENCES "auths" ("id");

ALTER TABLE "org_ops" ADD FOREIGN KEY ("org_id") REFERENCES "orgs" ("id");

ALTER TABLE "org_ops" ADD FOREIGN KEY ("op_id") REFERENCES "ops" ("id");

ALTER TABLE "org_users" ADD FOREIGN KEY ("org_id") REFERENCES "orgs" ("id");

ALTER TABLE "org_users" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "op_users" ADD FOREIGN KEY ("op_id") REFERENCES "ops" ("id");

ALTER TABLE "op_users" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "org_endpoints" ADD FOREIGN KEY ("org_id") REFERENCES "orgs" ("id");

ALTER TABLE "org_endpoints" ADD FOREIGN KEY ("endpoint_id") REFERENCES "endpoints" ("id");

ALTER TABLE "op_endpoints" ADD FOREIGN KEY ("op_id") REFERENCES "ops" ("id");

ALTER TABLE "op_endpoints" ADD FOREIGN KEY ("endpoint_id") REFERENCES "endpoints" ("id");

ALTER TABLE "user_endpoints" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_endpoints" ADD FOREIGN KEY ("endpoint_id") REFERENCES "endpoints" ("id");

CREATE UNIQUE INDEX ON "endpoints" ("group", "path");
