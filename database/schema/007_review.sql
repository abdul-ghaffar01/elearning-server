CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS review(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    tutorialId uuid REFERENCES tutorial(id) ON DELETE CASCADE,
    userId uuid REFERENCES users(id) ON DELETE CASCADE,
    parentId uuid REFERENCES review(id) ON DELETE CASCADE,
    rating smallint CHECK (rating >= 1 AND rating <= 5),
    content TEXT,
    reviewedAt TIMESTAMPTZ DEFAULT NOW()
)