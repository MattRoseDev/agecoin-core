-- Extentions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Enums 
CREATE TYPE user_roles AS ENUM ('USER', 'ADMIN');

-- Tables
CREATE TABLE "user" (
	"id" uuid NOT NULL DEFAULT uuid_generate_v4(),
	"username" varchar(32) NOT NULL UNIQUE,
	"email" varchar(100) NOT NULL UNIQUE,
	"fullname" varchar(64) NOT NULL,
	"role" user_roles DEFAULT 'USER',
	"birthday" TIMESTAMP,
	"max_age" integer DEFAULT NUll,
  "created_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"deleted_at" TIMESTAMP,
	CONSTRAINT "user_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

CREATE TABLE "password" (
	"id" uuid NOT NULL DEFAULT uuid_generate_v4(),
	"user_id" uuid NOT NULL,
	"password" TEXT NOT NULL,
	"created_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"deleted_at" TIMESTAMP,
 CONSTRAINT "password_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

CREATE TABLE "task" (
	"id" uuid NOT NULL DEFAULT uuid_generate_v4(),
	"user_id" uuid NOT NULL,
	"title" TEXT NOT NULL,
	"description" TEXT,
	"amount" integer, 
	"created_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"deleted_at" TIMESTAMP,
 	CONSTRAINT "task_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

CREATE TABLE "current_task" (
	"id" uuid NOT NULL DEFAULT uuid_generate_v4(),
	"user_id" uuid NOT NULL,
	"task_id" uuid NOT NULL,
	"default_amount" integer NOT NULL,
	"amount" integer,
	"status" integer NOT NULL,
	"description" TEXT,
	"active" BOOLEAN NOT NULL DEFAULT FALSE,
	"created_at" TIMESTAMP NOT NULL,
	"updated_at" TIMESTAMP NOT NULL,
	"deleted_at" TIMESTAMP,
	CONSTRAINT "current_task_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

-- Triggers
CREATE OR REPLACE FUNCTION trigger_set_updated_at()   
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at= now();
    RETURN NEW;   
END;
$$ language 'plpgsql';

CREATE TRIGGER update_user BEFORE UPDATE ON "user" FOR EACH ROW EXECUTE PROCEDURE  trigger_set_updated_at();
CREATE TRIGGER update_password BEFORE UPDATE ON "password" FOR EACH ROW EXECUTE PROCEDURE  trigger_set_updated_at();
CREATE TRIGGER update_task BEFORE UPDATE ON "task" FOR EACH ROW EXECUTE PROCEDURE  trigger_set_updated_at();
CREATE TRIGGER update_current_task BEFORE UPDATE ON "current_task" FOR EACH ROW EXECUTE PROCEDURE  trigger_set_updated_at();

-- Foreign keys
ALTER TABLE "password" ADD CONSTRAINT "password_fk0" FOREIGN KEY ("user_id") REFERENCES "user"("id");

ALTER TABLE "task" ADD CONSTRAINT "task_fk0" FOREIGN KEY ("user_id") REFERENCES "user"("id");

ALTER TABLE "current_task" ADD CONSTRAINT "current_task_fk0" FOREIGN KEY ("user_id") REFERENCES "user"("id");
ALTER TABLE "current_task" ADD CONSTRAINT "current_task_fk1" FOREIGN KEY ("task_id") REFERENCES "task"("id");
