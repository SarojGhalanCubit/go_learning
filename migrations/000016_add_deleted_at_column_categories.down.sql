DROP INDEX IF EXISTS idx_categories_deleted_at;
ALTER TABLE categories DROP COLUMN IF EXISTS deleted_at;
