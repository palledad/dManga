CREATE TABLE IF NOT EXISTS users(
    wallet_address VARCHAR (42) UNIQUE PRIMARY KEY,
    name VARCHAR (50) UNIQUE NOT NULL,
    icon_url VARCHAR (300),
    biography VARCHAR (160)
);