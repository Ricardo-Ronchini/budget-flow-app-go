CREATE TABLE IF NOT EXISTS expenses (
    expense_id text PRIMARY KEY,
    user_id text NOT NULL,
    name text NOT NULL,
    value numeric NOT NULL,
    date TIMESTAMP,
    created_at TIMESTAMP,
    modified_at TIMESTAMP
);
