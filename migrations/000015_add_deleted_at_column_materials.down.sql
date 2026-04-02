DROP INDEX IF EXISTS idx_materials_deleted_at;
ALTER TABLE materials DROP COLUMN IF EXISTS deleted_at;
