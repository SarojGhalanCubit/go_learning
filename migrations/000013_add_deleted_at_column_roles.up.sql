ALTER TABLE roles 
ADD COLUMN deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL;

CREATE INDEX idx_roles_deleted_at ON roles (deleted_at);
