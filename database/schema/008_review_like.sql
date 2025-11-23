
CREATE TABLE IF NOT EXISTS review_like(
    reviewId uuid REFERENCES review(id) ON DELETE CASCADE,
    userId uuid REFERENCES users(id) ON DELETE CASCADE,
    likedAt TIMESTAMPTZ DEFAULT NOW(),
    PRIMARY KEY (reviewId, userId)
)