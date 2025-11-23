CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- CREATE TYPE contenttype AS ENUM ('Video', 'Text');

CREATE TABLE IF NOT EXISTS lesson(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    curriculumId uuid REFERENCES curriculum(id) ON DELETE CASCADE,
    contentType VARCHAR(20) NOT NULL CHECK(contentType IN ('Video', 'Text')),
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    videoUrl TEXT,
    textContent TEXT,
    position smallint NOT NULL,
    duration varchar(100),
    createdAt TIMESTAMPTZ DEFAULT NOW()
)