CREATE TABLE "UploadedFile" (
  "id" string PRIMARY KEY,
  "URL" string,
  "size" integer,
  "type" string
);

CREATE TABLE "MeditationGroup" (
  "id" string PRIMARY KEY,
  "title" string,
  "inProduction" boolean,
  "inBeta" boolean,
  "key" string
);

CREATE TABLE "Meditation" (
  "id" string PRIMARY KEY,
  "key" string,
  "groupId" string,
  "benefits" array,
  "access" boolean,
  "index" integer,
  "description" text,
  "title" string,
  "inProduction" boolean,
  "collaboration" boolean,
  "premium" boolean,
  "pricing" string,
  "collaborator" string
);

CREATE TABLE "Collaborator" (
  "id" string PRIMARY KEY,
  "type" string,
  "description" text,
  "collaborator" string
);

CREATE TABLE "Pricing" (
  "id" string PRIMARY KEY,
  "price" integer,
  "type" string,
  "ratio" integer
);

CREATE TABLE "MeditationSettings" (
  "id" string PRIMARY KEY,
  "meditationId" string,
  "duration" boolean,
  "level" boolean,
  "loop" boolean,
  "dozee" boolean,
  "musicJump" boolean
);

CREATE TABLE "MeditationImages" (
  "id" string PRIMARY KEY,
  "meditationId" string,
  "header" string,
  "playerOverlay" string,
  "card" string
);

CREATE TABLE "MeditationColors" (
  "id" string PRIMARY KEY,
  "meditationId" string,
  "button" string,
  "alarmBackground" string,
  "progressBar" string,
  "cardBackground" string,
  "text" string,
  "title" string,
  "playerBackground" string
);

CREATE TABLE "MeditationLevels" (
  "id" string PRIMARY KEY,
  "meditationId" string,
  "key" string,
  "title" string,
  "minDuration" integer,
  "duration" integer,
  "MGTime" integer,
  "OBTime" integer,
  "file" string,
  "sessionEnd" integer
);

CREATE TABLE "SleepMusicGroup" (
  "id" string PRIMARY KEY,
  "premium" boolean,
  "pricing" string,
  "album" string,
  "albumId" string,
  "inProduction" boolean,
  "thumbnail" string,
  "headerImage" string
);

CREATE TABLE "SleepMusic" (
  "id" string PRIMARY KEY,
  "groupId" string,
  "duration" integer,
  "thumbnail" string,
  "file" string,
  "name" string,
  "key" string,
  "inProduction" boolean,
  "premium" boolean,
  "pricing" string
);

CREATE TABLE "Session" (
  "id" string PRIMARY KEY,
  "userId" string,
  "startTime" datetime,
  "endTime" datetime,
  "deviceId" string,
  "score" integer,
  "type" string,
  "email" string
);

CREATE TABLE "SessionGain" (
  "id" string PRIMARY KEY,
  "sessionId" string,
  "gain" integer,
  "type" string
);

CREATE TABLE "SessionFeeback" (
  "id" string PRIMARY KEY,
  "sessionId" string,
  "points" integer,
  "emoticon" string,
  "notes" array,
  "feedback" string
);

CREATE TABLE "SleepSessionAttributes" (
  "id" string PRIMARY KEY,
  "sessionId" string,
  "sleepTime" datetime,
  "wakeupTime" datetime,
  "sunriseTime" datetime,
  "apneaAHI" float64
);

CREATE TABLE "MeditationSessionAttributes" (
  "id" string,
  "sessionId" string,
  "level" string,
  "levelTitle" string,
  "title" string,
  "meditationId" string,
  "duration" integer,
  "pausePairs" array
);

CREATE TABLE "MeditationSessionSettings" (
  "id" string,
  "duration" integer,
  "withDozee" boolean,
  "deviceDisconnect" boolean,
  "wakeupAlarm" boolean
);

CREATE TABLE "SessionHeartRate" (
  "id" string PRIMARY KEY,
  "sessionId" string,
  "min" integer,
  "max" integer,
  "avg" integer,
  "coherence" integer,
  "start" integer,
  "end" integer,
  "violations" integer
);

CREATE TABLE "SessionBreathRate" (
  "id" string PRIMARY KEY,
  "sessionId" string,
  "min" integer,
  "max" integer,
  "avg" integer,
  "coherence" integer,
  "start" integer,
  "end" integer,
  "violations" integer
);

CREATE TABLE "SessionStress" (
  "id" string PRIMARY KEY,
  "sessionId" string,
  "min" integer,
  "max" integer,
  "avg" integer,
  "coherence" integer,
  "start" integer,
  "end" integer,
  "violations" integer
);

CREATE TABLE "SessionMovementParameters" (
  "id" string PRIMARY KEY,
  "sessionId" string,
  "duration" integer,
  "instances" integer
);

CREATE TABLE "SessionVitals" (
  "id" string PRIMARY KEY,
  "sessionId" string,
  "time" timestamp,
  "heartRate" float64,
  "breathRate" float64,
  "stress" float64,
  "recovery" float64
);

CREATE TABLE "SessionMovements" (
  "id" string PRIMARY KEY,
  "sessionId" string,
  "time" timestamp,
  "value" integer
);

CREATE TABLE "SessionPoints" (
  "id" string PRIMARY KEY,
  "sessionId" string,
  "stress" integer,
  "deep" integer,
  "restfullness" integer,
  "bedtimeResistance" integer,
  "HRV" integer,
  "heart" integer,
  "breath" integer,
  "duration" integer,
  "circadianRhythm" integer,
  "awakening" integer,
  "snoring" integer,
  "wakeupAfterSleepOnset" integer,
  "REM" integer,
  "movement" integer,
  "feedback" integer,
  "vitals" integer,
  "sleepLatency" integer,
  "sleepEfficiency" integer,
  "sleepHygiene" integer,
  "sleepQuality" integer
);

CREATE TABLE "SessionDurations" (
  "id" string PRIMARY KEY,
  "sessionId" string,
  "deep" integer,
  "REM" integer,
  "awake" integer,
  "total" integer,
  "totalAwake" integer
);

CREATE TABLE "SessionStages" (
  "id" string PRIMARY KEY,
  "sessionId" string,
  "time" timestamp,
  "value" integer
);

CREATE TABLE "SessionSnoring" (
  "id" string PRIMARY KEY,
  "sessionId" string,
  "time" timestamp,
  "value" value
);

CREATE TABLE "SessionHRV30Parameters" (
  "id" string PRIMARY KEY,
  "sessionId" string,
  "time" timestamp,
  "PNN50" float,
  "RMSSD" float,
  "SDNN" float
);

ALTER TABLE "Meditation" ADD FOREIGN KEY ("groupId") REFERENCES "MeditationGroup" ("id");

ALTER TABLE "Meditation" ADD FOREIGN KEY ("pricing") REFERENCES "Pricing" ("id");

ALTER TABLE "Meditation" ADD FOREIGN KEY ("collaborator") REFERENCES "Collaborator" ("id");

ALTER TABLE "MeditationSettings" ADD FOREIGN KEY ("meditationId") REFERENCES "Meditation" ("id");

ALTER TABLE "MeditationImages" ADD FOREIGN KEY ("meditationId") REFERENCES "Meditation" ("id");

ALTER TABLE "MeditationColors" ADD FOREIGN KEY ("meditationId") REFERENCES "Meditation" ("id");

ALTER TABLE "MeditationLevels" ADD FOREIGN KEY ("meditationId") REFERENCES "Meditation" ("id");

ALTER TABLE "MeditationLevels" ADD FOREIGN KEY ("file") REFERENCES "UploadedFile" ("id");

ALTER TABLE "SleepMusicGroup" ADD FOREIGN KEY ("pricing") REFERENCES "Pricing" ("id");

ALTER TABLE "SleepMusicGroup" ADD FOREIGN KEY ("thumbnail") REFERENCES "UploadedFile" ("id");

ALTER TABLE "SleepMusicGroup" ADD FOREIGN KEY ("headerImage") REFERENCES "UploadedFile" ("id");

ALTER TABLE "SleepMusic" ADD FOREIGN KEY ("groupId") REFERENCES "SleepMusicGroup" ("id");

ALTER TABLE "SleepMusic" ADD FOREIGN KEY ("thumbnail") REFERENCES "UploadedFile" ("id");

ALTER TABLE "SleepMusic" ADD FOREIGN KEY ("file") REFERENCES "UploadedFile" ("id");

ALTER TABLE "SleepMusic" ADD FOREIGN KEY ("pricing") REFERENCES "Pricing" ("id");

ALTER TABLE "SessionGain" ADD FOREIGN KEY ("sessionId") REFERENCES "Session" ("id");

ALTER TABLE "SessionFeeback" ADD FOREIGN KEY ("sessionId") REFERENCES "Session" ("id");

ALTER TABLE "SleepSessionAttributes" ADD FOREIGN KEY ("sessionId") REFERENCES "Session" ("id");

ALTER TABLE "MeditationSessionAttributes" ADD FOREIGN KEY ("sessionId") REFERENCES "Session" ("id");

ALTER TABLE "MeditationSessionAttributes" ADD FOREIGN KEY ("meditationId") REFERENCES "Meditation" ("id");

ALTER TABLE "MeditationSessionSettings" ADD FOREIGN KEY ("id") REFERENCES "MeditationSessionAttributes" ("id");

ALTER TABLE "SessionHeartRate" ADD FOREIGN KEY ("sessionId") REFERENCES "Session" ("id");

ALTER TABLE "SessionBreathRate" ADD FOREIGN KEY ("sessionId") REFERENCES "Session" ("id");

ALTER TABLE "SessionStress" ADD FOREIGN KEY ("sessionId") REFERENCES "Session" ("id");

ALTER TABLE "SessionMovementParameters" ADD FOREIGN KEY ("sessionId") REFERENCES "Session" ("id");

ALTER TABLE "SessionVitals" ADD FOREIGN KEY ("sessionId") REFERENCES "Session" ("id");

ALTER TABLE "SessionMovements" ADD FOREIGN KEY ("sessionId") REFERENCES "Session" ("id");

ALTER TABLE "SessionPoints" ADD FOREIGN KEY ("sessionId") REFERENCES "Session" ("id");

ALTER TABLE "SessionDurations" ADD FOREIGN KEY ("sessionId") REFERENCES "Session" ("id");

ALTER TABLE "SessionStages" ADD FOREIGN KEY ("sessionId") REFERENCES "Session" ("id");

ALTER TABLE "SessionSnoring" ADD FOREIGN KEY ("sessionId") REFERENCES "Session" ("id");

ALTER TABLE "SessionHRV30Parameters" ADD FOREIGN KEY ("sessionId") REFERENCES "Session" ("id");
