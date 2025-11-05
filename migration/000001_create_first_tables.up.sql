CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    lastname       TEXT        NOT NULL,
    firstname       TEXT        NOT NULL,
    email      TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);