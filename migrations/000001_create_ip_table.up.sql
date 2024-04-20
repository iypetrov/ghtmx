CREATE TABLE IF NOT EXISTS website_visits (
    id UUID PRIMARY KEY,
    ip VARCHAR(45) NOT NULL,
    created_at TIMESTAMP
);
