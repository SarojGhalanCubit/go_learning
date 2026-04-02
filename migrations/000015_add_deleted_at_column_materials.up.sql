ALTER TABLE materials 
ADD COLUMN deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL;

CREATE INDEX idx_materials_deleted_at ON materials (deleted_at);
