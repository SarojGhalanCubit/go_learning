ALTER TABLE categories 
ADD COLUMN deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL;

CREATE INDEX idx_categories_deleted_at ON categories (deleted_at);
