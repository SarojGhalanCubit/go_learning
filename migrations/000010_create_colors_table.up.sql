CREATE TABLE IF NOT EXISTS colors (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    hex_code VARCHAR(7) NOT NULL, -- e.g., #FFFFFF
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
