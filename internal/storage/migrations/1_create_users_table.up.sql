CREATE TYPE role_enum AS ENUM (
    'employee',
    'moderator'
);

CREATE TABLE IF NOT EXISTS users
(
    id SERIAL PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    role role_enum NOT NULL
);