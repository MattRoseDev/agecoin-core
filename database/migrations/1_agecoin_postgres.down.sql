ALTER TABLE "password" DROP CONSTRAINT IF EXISTS "password_fk0";

ALTER TABLE "task" DROP CONSTRAINT IF EXISTS "task_fk0";

ALTER TABLE "task" DROP CONSTRAINT IF EXISTS "task_fk1";

ALTER TABLE "current_task" DROP CONSTRAINT IF EXISTS "current_task_fk0";

ALTER TABLE "current_task" DROP CONSTRAINT IF EXISTS "current_task_fk1";

ALTER TABLE "current_task_history" DROP CONSTRAINT IF EXISTS "current_task_history_fk0";

ALTER TABLE "current_task_history" DROP CONSTRAINT IF EXISTS "current_task_history_fk1";

DROP TABLE IF EXISTS "user";

DROP TABLE IF EXISTS "password";

DROP TABLE IF EXISTS "task";

DROP TABLE IF EXISTS "current_task";

DROP TABLE IF EXISTS "current_task_history";

DROP TYPE IF EXISTS "user_roles";

DROP TYPE IF EXISTS "current_task_history_type";