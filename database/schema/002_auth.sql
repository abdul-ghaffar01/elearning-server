CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS auth(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    userId uuid REFERENCES users(id) ON DELETE CASCADE,
    refreshToken TEXT NOT NULL,
    isExpired BOOLEAN DEFAULT FALSE,
    ip inet,
    loginAt TIMESTAMPTZ DEFAULT NOW(),
    logoutAt TIMESTAMPTZ,
    deviceType VARCHAR(100),
    os VARCHAR(100),
    browser VARCHAR(100),
    country VARCHAR(100),
    city VARCHAR(100),
    expireTime TIMESTAMPTZ NOT NULL
)