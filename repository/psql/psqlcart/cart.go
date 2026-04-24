package psqlcart

import (
	"context"
	"fmt"

	"github.com/HosseinForouzan/E-Commerce-API/param"
)

func (d *DB) AddItem(req param.AddItemRequest) error {
	query := `INSERT INTO cart_items (user_id, product_id, quantity)
				VALUES ($1, $2, $3)
				ON CONFLICT (user_id, product_id)
				DO UPDATE SET quantity = cart_items.quantity + EXCLUDED.quantity; `
	_, err := d.conn.Conn().Exec(context.Background(), query, req.UserID, req.ProductID, req.Quantity)
	if err != nil {
		return fmt.Errorf("can't insert item to cart: %w", err)
	}

	return nil
}