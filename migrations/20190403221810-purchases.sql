-- +migrate Up
CREATE TABLE purchases (
  customer_id    int REFERENCES customers (id) ON UPDATE CASCADE ON DELETE CASCADE,
  book_id int REFERENCES books (id) ON UPDATE CASCADE ON DELETE CASCADE,
  primary key ( customer_id, book_id )
);
-- +migrate Down
DROP TABLE purchases;
