SELECT id, fullname, email, profile, joined_at, profile_setupped
FROM users WHERE email = $1