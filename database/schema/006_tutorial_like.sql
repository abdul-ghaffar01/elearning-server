
CREATE TABLE IF NOT EXISTS tut_like(
    tutId uuid REFERENCES tutorial(id) ON DELETE CASCADE,
    userId uuid REFERENCES users(id) ON DELETE CASCADE,
    likedAt TIMESTAMPTZ DEFAULT NOW(),
    PRIMARY KEY (tutId, userId)
)