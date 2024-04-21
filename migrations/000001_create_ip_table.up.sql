CREATE TABLE IF NOT EXISTS request_ips (
    id UUID PRIMARY KEY,
    ip VARCHAR(45) NOT NULL,
    created_at TIMESTAMP
);
