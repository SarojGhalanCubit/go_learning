ALTER TABLE sizes 
ADD COLUMN deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL;

CREATE INDEX idx_sizes_deleted_at ON sizes (deleted_at);
