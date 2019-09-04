-- +migrate Up
CREATE TABLE books (
    id serial PRIMARY KEY, 
    title VARCHAR (250) NOT NULL,
    author_id INTEGER NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

-- +migrate Down
DROP TABLE companies;
