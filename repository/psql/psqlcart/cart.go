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

func (d *DB) GetCart(userID uint) (param.CartResponse, error) {
	query := `SELECT
			c.product_id,
			c.quantity,
			p.name,
			p.price
		FROM cart_items c
		JOIN products p ON p.id = c.product_id
		WHERE c.user_id = $1;`
	
	rows, err := d.conn.Conn().Query(context.Background(), query, userID)
	if err != nil {
		return param.CartResponse{}, fmt.Errorf("can't get items: %w", err)
	}

	items := []param.CartItemResponse{}

	for rows.Next() {
		var item param.CartItemResponse
		err := rows.Scan(
			&item.ProductID,
			&item.Quantity,
			&item.Name,
			&item.Price,
		)
		if err != nil {
			return param.CartResponse{}, fmt.Errorf("can't scan items: %w", err)
		}
		items = append(items, item)
	}

	return param.CartResponse{
		Items: items,
		Total: 0,
	}, nil

}