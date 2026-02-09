CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    password VARCHAR(255),
    phone_number VARCHAR(13) UNIQUE,
    email VARCHAR(255) UNIQUE,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);