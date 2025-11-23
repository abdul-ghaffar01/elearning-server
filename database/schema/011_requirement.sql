CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS requirement(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    tutId uuid REFERENCES tutorial(id) ON DELETE CASCADE,
    description TEXT NOT NULL,
    position smallint NOT NULL,
    createdAt TIMESTAMPTZ DEFAULT NOW()
)