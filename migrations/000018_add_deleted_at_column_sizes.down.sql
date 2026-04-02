DROP INDEX IF EXISTS idx_sizes_deleted_at;
ALTER TABLE sizes DROP COLUMN IF EXISTS deleted_at;
