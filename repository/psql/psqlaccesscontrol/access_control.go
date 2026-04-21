package psqlaccesscontrol

import (
	"context"
	"fmt"
)

func (d *DB) IsUserAdmin(userID uint8) (bool, error) {
	var role string
	err := d.conn.Conn().QueryRow(context.Background(), "SELECT role FROM users").Scan(&role)
	if err != nil {
		return false, fmt.Errorf("can't retrive role of user: %w", err)
	}

	if role == "admin" {
		return true, nil
	}

	return false, nil
}