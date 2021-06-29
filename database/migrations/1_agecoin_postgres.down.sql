ALTER TABLE "password" DROP CONSTRAINT IF EXISTS "password_fk0";

ALTER TABLE "task" DROP CONSTRAINT IF EXISTS "task_fk0";

ALTER TABLE "task_history" DROP CONSTRAINT IF EXISTS "task_history_fk0";

ALTER TABLE "task_history" DROP CONSTRAINT IF EXISTS "task_history_fk1";

DROP TABLE IF EXISTS "user";

DROP TABLE IF EXISTS "password";

DROP TABLE IF EXISTS "task";

DROP TABLE IF EXISTS "task_history";

DROP TYPE IF EXISTS "user_roles";

DROP TYPE IF EXISTS "task_history_type";