-- Down Migration
DROP TRIGGER IF EXISTS update_users_updated_at ON users;
DROP FUNCTION IF EXISTS update_updated_at_column();

ALTER TABLE users 
DROP COLUMN IF EXISTS created_at,
DROP COLUMN IF EXISTS updated_at;
