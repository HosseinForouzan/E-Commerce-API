-- +migrate Up
CREATE TABLE order_items (
  id SERIAL PRIMARY KEY,
  order_id INT NOT NULL,
  product_id INT NOT NULL,
  product_name VARCHAR(255) NOT NULL,
  unit_price INT NOT NULL,
  quantity INT NOT NULL,
  subtotal INT NOT NULL
);
-- +migrate Down
DROP TABLE order_items;