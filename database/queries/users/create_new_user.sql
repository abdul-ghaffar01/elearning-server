INSERT INTO users (fullname, email, profile)
VALUES ($1, $2, $3)
RETURNING id, fullname, email, profile, joined_at, profile_setupped;
