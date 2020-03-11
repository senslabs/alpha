CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid,
  "start_time" timestamp,
  "end_time" timestamp
);

CREATE INDEX ON "sessions" ("end_time");

CREATE INDEX ON "sessions" ("user_id");
