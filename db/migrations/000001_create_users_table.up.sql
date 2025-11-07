CREATE TABLE IF NOT EXISTS users (
    user_id text PRIMARY KEY,
    name text NOT NULL,
    email text UNIQUE NOT NULL,
    username TEXT,
    password_hash TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
