CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS category(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    prentId uuid REFERENCES category(id) ON DELETE CASCADE,
    createdAt TIMESTAMPTZ DEFAULT NOW()
)