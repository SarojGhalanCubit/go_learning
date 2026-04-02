DROP INDEX IF EXISTS idx_product_variants_deleted_at;
ALTER TABLE product_variants DROP COLUMN IF EXISTS deleted_at;
