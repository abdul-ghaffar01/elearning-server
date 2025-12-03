-- Enable UUID function
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    fullname VARCHAR(100) NOT NULL,
    dob DATE,
    country VARCHAR(50),
    email VARCHAR(100) UNIQUE NOT NULL,
    role VARCHAR(20),
    profile VARCHAR(255),
    password varchar(512),
    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deactivated BOOLEAN DEFAULT FALSE,
    profile_setupped BOOLEAN DEFAULT FALSE
);
