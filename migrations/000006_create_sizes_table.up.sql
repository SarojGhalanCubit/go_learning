CREATE TABLE IF NOT EXISTS sizes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL, -- e.g., "Large", "XL", "42"
    sort_order INTEGER DEFAULT 0, -- Used to sort S, M, L correctly in UI
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
