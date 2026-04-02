ALTER TABLE product_variants 
ADD COLUMN deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL;

CREATE INDEX idx_product_variants_deleted_at ON product_variants (deleted_at);
