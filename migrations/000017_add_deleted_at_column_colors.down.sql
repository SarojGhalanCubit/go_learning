DROP INDEX IF EXISTS idx_colors_deleted_at;
ALTER TABLE colors DROP COLUMN IF EXISTS deleted_at;
