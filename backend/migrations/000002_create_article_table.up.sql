CREATE TABLE IF NOT EXISTS articles(
    id UUID PRIMARY KEY,
    title VARCHAR (50) NOT NULL,
    content JSON NOT NULL,
    author VARCHAR (42) NOT NULL,
    alias VARCHAR (50) UNIQUE,
    created_at TIMESTAMP NOT NULL,
    updated_At TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);