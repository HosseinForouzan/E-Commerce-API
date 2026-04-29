-- +migrate Up
CREATE TABLE payments (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    amount INT NOT NULL,
    status VARCHAR(30) NOT NULL,
    provider VARCHAR(50) NOT NULL,
    transaction_id VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),

    FOREIGN KEY (order_id) REFERENCES orders(id)
);
-- +migrate Down
DROP TABLE payments;