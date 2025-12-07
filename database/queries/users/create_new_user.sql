INSERT INTO users (
    fullname,
    email,
    password,
    profile
)
VALUES (
    $1,               -- fullname
    $2,               -- email
    $3,               -- hashed password (nullable)
    $4                -- profile (nullable)
)
RETURNING 
    id,
    fullname,
    email,
    profile,
    profile_setupped;
