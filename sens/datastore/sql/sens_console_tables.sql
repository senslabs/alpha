CREATE TABLE "auths" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "email" text UNIQUE,
  "mobile" text UNIQUE NOT NULL,
  "social" text UNIQUE,
  "first_name" text,
  "last_name" text,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "orgs" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "auth_id" uuid,
  "name" text UNIQUE,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "ops" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "auth_id" uuid,
  "created_at" timestamp,
  "updated_at" timestamp,
  "status" text
);

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "auth_id" uuid,
  "created_at" timestamp,
  "updated_at" timestamp,
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
  "created_at" timestamp DEFAULT (now()),
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
