-- +migrate Up
CREATE TABLE orders (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  status VARCHAR(30) NOT NULL,
  total_amount INT NOT NULL,
  created_at TIMESTAMP DEFAULT NOW()
);
-- +migrate Down
DROP TABLE orders;