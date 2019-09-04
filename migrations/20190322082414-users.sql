-- +migrate Up
CREATE TABLE customers (
    id serial PRIMARY KEY, 
    email VARCHAR (250) NOT NULL UNIQUE,
    first_name VARCHAR (250) NOT NULL,
    last_name VARCHAR (250) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

-- +migrate Down
DROP TABLE customers;