UPDATE users
SET profile = $2
WHERE email = $1
RETURNING id, fullname, email, profile, joined_at
	