DROP INDEX IF EXISTS idx_roles_deleted_at;
ALTER TABLE roles DROP COLUMN IF EXISTS deleted_at;
