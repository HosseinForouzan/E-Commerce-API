# E-Commerce API

A modular backend API for an e-commerce platform built with Go, PostgreSQL, and Echo.

This project was developed to practice backend system design, clean architecture, transactional workflows, and real-world e-commerce features such as authentication, shopping cart, checkout, orders, and payments.

---

## Tech Stack

- Go
- Echo Framework
- PostgreSQL
- pgx
- JWT Authentication
- Koanf (configuration management)

---

## Features

### Authentication & Authorization

- User registration
- User login
- JWT-based authentication
- Role-based access control (`user`, `admin`)

### Products

- Create product (admin)
- Update product (admin)
- Delete product (admin)
- Get all products
- Get single product
- Search products
- Filter products by category

### Categories

- Product categorization support

### Shopping Cart

- Add item to cart
- Update quantity
- Remove item
- Get current cart
- Clear cart

### Checkout & Orders

- Transactional checkout flow
- Create order from cart
- Create order items
- Reduce product stock
- Clear cart after successful checkout
- Get user orders
- Get order details
- Cancel order

### Payments

- Mock payment gateway
- Hosted fake payment page
- Success / failure callback flow
- Payment status tracking
- Order marked as paid after successful payment

---

## Project Structure

```text
  users/
  products/
  cart/
  orders/
  payment/
  config/
  handler/
    middleware/
  ```

Each module follows layered architecture:
```text
handler -> service -> repository -> database
```

## Database Design

Main tables:

```text
users
products
categories
cart_items
orders
order_items
payments
```

## Checkout Flow

``` text
Add products to cart
↓
Checkout
↓
Create order
↓
Create order items
↓
Decrease stock
↓
Clear cart
↓
Pay order
```
All critical checkout operations are executed using PostgreSQL transactions.

## Payment Flow

```text
POST /orders/:id/pay
↓
returns payment_url
↓
Open fake gateway page
↓
Success / Fail
↓
Callback
↓
Update payment + order status
```

## Example API Routes

### Auth

```http
POST /register
POST /login
GET /profile
```

### Products

```http
GET /products
GET /products/:id
POST /admin/products
PATCH /admin/products/:id
DELETE /admin/products/:id
```

### Cart
```http
GET /cart
POST /cart/items
PATCH /cart/items/:id
DELETE /cart/items/:id
DELETE /cart
```

### Orders

```http
POST /checkout
GET /orders
GET /orders/:id
PATCH /orders/:id/cancel
```

### Payments
```http
POST /orders/:id/pay
GET /payment/mock/:paymentID
POST /payment/mock/:paymentID/success
POST /payment/mock/:paymentID/fail
```

## What I Learned
- Building scalable REST APIs in Go
- Clean architecture patterns
- Authentication & authorization
- PostgreSQL schema design
- Database transactions
- Order lifecycle management
- Payment workflow design
- Service modularization

## Future Imporvments
- Docker support
- Swagger docs
- Unit / integration tests
- Redis caching
- Real payment gateway integration
- Email notifications
- Admin dashboard





