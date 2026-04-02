ALTER TABLE colors 
ADD COLUMN deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL;

CREATE INDEX idx_colors_deleted_at ON colors (deleted_at);
