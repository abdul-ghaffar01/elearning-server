CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS choice(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    mcqId uuid REFERENCES mcq(id) ON DELETE CASCADE,
    choiceText TEXT NOT NULL,
    createdAt TIMESTAMPTZ DEFAULT NOW()
)