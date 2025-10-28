CREATE TABLE IF NOT EXISTS users (
    user_id text PRIMARY KEY,
    name text NOT NULL,
    email text UNIQUE NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
