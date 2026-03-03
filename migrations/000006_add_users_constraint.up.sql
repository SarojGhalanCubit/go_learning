ALTER TABLE users
ADD CONSTRAINT user_email_unique UNIQUE (email),
ADD CONSTRAINT user_phone_unique UNIQUE (phone_number);
