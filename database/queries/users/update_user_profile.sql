UPDATE users
SET profile = $2
WHERE email = $1
RETURNING id, name, email, picture_url, created_at
	