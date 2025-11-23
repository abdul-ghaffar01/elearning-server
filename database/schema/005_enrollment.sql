CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS enrollment(
    tutId uuid REFERENCES tutorial(id) ON DELETE CASCADE,
    userId uuid REFERENCES users(id) ON DELETE CASCADE,
    lesson_completed smallint DEFAULT 0,
    dateEnrolled TIMESTAMPTZ DEFAULT NOW(),
    completedAt TIMESTAMPTZ,
    PRIMARY KEY (tutId, userId)
)