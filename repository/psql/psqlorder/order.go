package psqlorder

import (
	"context"
	"fmt"

	"github.com/HosseinForouzan/E-Commerce-API/param"
	"github.com/jackc/pgx/v5"
)

func (d *DB) CreateTx(ctx context.Context, tx pgx.Tx, userID uint, total uint) (uint, error) {
	var orderID uint

	err := tx.QueryRow(ctx,
		 `INSERT INTO orders(user_id, status, total_amount) VALUES($1, 'pending_payment' ,$2) RETURNING id`, userID, total).
		 Scan(&orderID)
	if err != nil {
		return 0, fmt.Errorf("can't insert order: %w", err)
	}

	return orderID, nil
}

func (d *DB) CreateItemTx(ctx context.Context, tx pgx.Tx, orderID uint, item param.CartItemResponse) error{
	_, err := tx.Exec(ctx,
	`INSERT INTO order_items(order_id, product_id, product_name, unit_price, quantity, subtotal) VALUES($1,$2,$3,$4,$5,$6)`,
						orderID, item.ProductID, item.Name, item.Price, item.Quantity, item.Subtotal)
	if err != nil {
		return fmt.Errorf("can't insert order items: %w", err)
	}

	return nil
}