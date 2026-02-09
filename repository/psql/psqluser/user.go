package psqluser

import (
	"context"
	"fmt"

	"github.com/HosseinForouzan/E-Commerce-API/entity"
)

func (d *DB) Register(user entity.User) (entity.User, error) {
	var id uint
	err := d.conn.Conn().QueryRow(context.Background(), 
							`INSERT INTO users(name,password,email,created_at,updated_at)
							VALUES($1,$2,$3,$4,$5) RETURNING id`, user.Name, user.Password, user.Email, user.CreatedAt, user.UpdatedAt).Scan(&id)
	if err != nil {
		return entity.User{}, fmt.Errorf("can't execute query %w", err)
	}

	user.ID = id

	return user, nil
}