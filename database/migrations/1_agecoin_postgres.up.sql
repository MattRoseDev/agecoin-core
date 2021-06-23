-- Extentions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Enums 
DROP TYPE IF EXISTS "user_roles";
DROP TYPE IF EXISTS "current_task_history_type";
CREATE TYPE user_roles AS ENUM ('USER', 'ADMIN');
CREATE TYPE current_task_history_type AS ENUM ('START', 'PAUSE', 'FINISH', 'CANCEL');

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
	"default_coins" integer NOT NULL,
	"created_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"deleted_at" TIMESTAMP,
 	CONSTRAINT "task_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);
-- Current Task Fields
-- status: 0: Not started, 1: started, 2: finished, 3: canceled 

CREATE TABLE "current_task" (
	"id" uuid NOT NULL DEFAULT uuid_generate_v4(),
	"user_id" uuid NOT NULL,
	"task_id" uuid NOT NULL,
	"default_coins" integer NOT NULL,
	"coins" integer,
	"status" integer NOT NULL DEFAULT 0,
	"description" TEXT,
	"active" BOOLEAN DEFAULT FALSE,
	"created_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"deleted_at" TIMESTAMP,
	CONSTRAINT "current_task_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

CREATE TABLE "current_task_history" (
	"id" uuid NOT NULL DEFAULT uuid_generate_v4(),
	"user_id" uuid NOT NULL,
	"current_task_id" uuid NOT NULL,
	"type" current_task_history_type NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"deleted_at" TIMESTAMP,
	CONSTRAINT "current_task_history_pk" PRIMARY KEY ("id")
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
CREATE TRIGGER update_current_task_history BEFORE UPDATE ON "current_task_history" FOR EACH ROW EXECUTE PROCEDURE  trigger_set_updated_at();

-- Foreign keys
ALTER TABLE "password" ADD CONSTRAINT "password_fk0" FOREIGN KEY ("user_id") REFERENCES "user"("id");

ALTER TABLE "task" ADD CONSTRAINT "task_fk0" FOREIGN KEY ("user_id") REFERENCES "user"("id");

ALTER TABLE "current_task" ADD CONSTRAINT "current_task_fk0" FOREIGN KEY ("user_id") REFERENCES "user"("id");
ALTER TABLE "current_task" ADD CONSTRAINT "current_task_fk1" FOREIGN KEY ("task_id") REFERENCES "task"("id");

ALTER TABLE "current_task_history" ADD CONSTRAINT "current_task_history_fk0" FOREIGN KEY ("user_id") REFERENCES "user"("id");
ALTER TABLE "current_task_history" ADD CONSTRAINT "current_task_history_fk1" FOREIGN KEY ("current_task_id") REFERENCES "current_task"("id");
