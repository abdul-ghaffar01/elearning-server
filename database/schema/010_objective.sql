CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS objective(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    tutId uuid REFERENCES tutorial(id) ON DELETE CASCADE,
    heading TEXT NOT NULL,
    position smallint NOT NULL,
    createdAt TIMESTAMPTZ DEFAULT NOW()
)