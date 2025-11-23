CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- CREATE TYPE tutorial_difficulty AS ENUM ('Beginner', 'Intermediate', 'Advanced');
-- CREATE TYPE tutorial_visibility AS ENUM ('Public', 'Private');
-- CREATE TYPE tutorial_status AS ENUM ('Draft', 'Published', 'Archived');


CREATE TABLE IF NOT EXISTS tutorial(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    instructorId uuid REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    categoryId uuid REFERENCES category(id) ON DELETE SET NULL,
    -- difficulty tutorial_difficulty DEFAULT 'Beginner',
    -- visibility tutorial_visibility DEFAULT 'Public',
    -- status tutorial_status DEFAULT 'Draft',
    originalPrice NUMERIC(10,2) DEFAULT 0.00,
    sellingPrice NUMERIC(10,2) DEFAULT 0.00 CHECK (sellingPrice <= originalPrice),
    thumbnailUrl TEXT,
    tags TEXT[],
    difficulty VARCHAR(20) NOT NULL DEFAULT 'Beginner' CHECK (difficulty IN ('Beginner', 'Intermediate', 'Advanced')),
    visibility VARCHAR(20) NOT NULL DEFAULT 'Public' CHECK (visibility IN ('Public', 'Private')),
    status VARCHAR(20) NOT NULL DEFAULT 'Draft' CHECK (status IN ('Draft', 'Published', 'Archived')),
    createdAt TIMESTAMPTZ DEFAULT NOW(),
    updatedAt TIMESTAMPTZ DEFAULT NOW()
);
