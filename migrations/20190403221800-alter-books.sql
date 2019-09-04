-- +migrate Up
ALTER TABLE books
ADD FOREIGN KEY (author_id) REFERENCES authors(id) ON DELETE CASCADE;
-- +migrate Down
ALTER TABLE books
DROP CONSTRAINT books_author_id_fkey;