CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS mcq(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    lessonId uuid REFERENCES lesson(id) ON DELETE CASCADE,
    question TEXT NOT NULL,
    position smallint NOT NULL,
    createdAt TIMESTAMPTZ DEFAULT NOW()
)