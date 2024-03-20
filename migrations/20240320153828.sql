-- Create "users" table
CREATE TABLE "users" (
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "name" text NULL,
  "surname" text NULL,
  "id" text NOT NULL,
  PRIMARY KEY ("id")
);
