CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS curriculum(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    tutId uuid REFERENCES tutorial(id) ON DELETE CASCADE,
    heading VARCHAR(255) NOT NULL,
    isFree BOOLEAN DEFAULT FALSE,
    position smallint NOT NULL,
    createdAt TIMESTAMPTZ DEFAULT NOW()
)