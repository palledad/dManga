CREATE TABLE IF NOT EXISTS tags(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR (50) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_At TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);