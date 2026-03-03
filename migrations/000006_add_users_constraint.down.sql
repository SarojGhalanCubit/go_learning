ALTER TABLE users
DROP CONSTRAINT IF EXISTS user_email_unique,
DROP CONSTRAINT IF EXISTS user_phone_unique;
