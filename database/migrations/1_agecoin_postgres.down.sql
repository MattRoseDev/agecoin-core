ALTER TABLE "password" DROP CONSTRAINT IF EXISTS "password_fk0";

ALTER TABLE "task" DROP CONSTRAINT IF EXISTS "task_fk0";

ALTER TABLE "task" DROP CONSTRAINT IF EXISTS "task_fk1";

ALTER TABLE "user_task" DROP CONSTRAINT IF EXISTS "user_task_fk0";

ALTER TABLE "user_task" DROP CONSTRAINT IF EXISTS "user_task_fk1";

ALTER TABLE "current_task" DROP CONSTRAINT IF EXISTS "current_task_fk0";

ALTER TABLE "current_task" DROP CONSTRAINT IF EXISTS "current_task_fk1";

DROP TABLE IF EXISTS "user";

DROP TABLE IF EXISTS "password";

DROP TABLE IF EXISTS "task";

DROP TABLE IF EXISTS "user_task";

DROP TABLE IF EXISTS "current_task";

DROP TYPE IF EXISTS "user_roles";