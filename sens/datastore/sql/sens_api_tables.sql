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

ALTER TABLE "session_properties" ADD FOREIGN KEY ("session_id") REFERENCES "sessions" ("id");

CREATE INDEX ON "sessions" ("ended_at");

CREATE INDEX ON "sessions" ("user_id");
