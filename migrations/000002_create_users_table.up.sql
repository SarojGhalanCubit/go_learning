CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    age INT ,
    email TEXT NOT NULL,
    phone_number TEXT ,
    password TEXT NOT NULL,
    role_id INT NOT NULL DEFAULT 2,
    is_email_verified BOOLEAN DEFAULT FALSE,
    is_phone_verified BOOLEAN DEFAULT FALSE,
    CONSTRAINT fk_role FOREIGN KEY (role_id) REFERENCES roles(id),
    CONSTRAINT user_email_unique UNIQUE (email),
    CONSTRAINT user_phone_unique UNIQUE (phone_number)
);
