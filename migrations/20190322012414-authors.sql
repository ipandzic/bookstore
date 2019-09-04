-- +migrate Up
CREATE TABLE authors (
    id serial PRIMARY KEY, 
    first_name VARCHAR (250) NOT NULL,
    last_name VARCHAR (250) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

-- +migrate Down
DROP TABLE authors;