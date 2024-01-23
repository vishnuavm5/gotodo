-- +goose Up

CREATE TABLE users(
    id UUID PRIMARY KEY,
    name VARCHAR(60) NOT NULL,
    username VARCHAR(60) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down

DROP TABLE users;