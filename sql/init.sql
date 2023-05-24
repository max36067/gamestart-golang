\c my_database;

CREATE TABLE "user" (
  "id" bigserial PRIMARY KEY,
  "email" varchar NOT NULL,
  "name" varchar NOT NULL,
  "password" varchar NOT NULL,
  "is_active" bool NOT NULL DEFAULT (true),
  "is_super_user" bool NOT NULL
);

CREATE TABLE "salt" (
  "id" bigserial PRIMARY KEY,
  "email" varchar NOT NULL,
  "salt" varchar NOT NULL
);

CREATE INDEX ON "user" ("email");

CREATE INDEX ON "salt" ("email");
