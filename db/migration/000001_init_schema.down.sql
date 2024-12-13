-- Drop foreign key constraints first
ALTER TABLE "verify_emails" DROP CONSTRAINT "verify_emails_username_fkey";
ALTER TABLE "sessions" DROP CONSTRAINT "sessions_username_fkey";

-- Drop tables in reverse order of creation
DROP TABLE IF EXISTS "sessions";
DROP TABLE IF EXISTS "verify_emails";
DROP TABLE IF EXISTS "users";