package psqluser

import (
	"context"
	"database/sql"
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

func (d *DB) GetUserByEmail(email string) (entity.User, error) {
	var user entity.User
	var phone sql.NullString

	user.PhoneNumber = ""
	err := d.conn.Conn().QueryRow(context.Background(),
									"SELECT * FROM users WHERE email=$1", email).Scan(
									&user.ID, &user.Name, &user.Password, &phone, &user.Email, &user.CreatedAt, &user.UpdatedAt)	
	if err != nil {
		return entity.User{}, fmt.Errorf("can't retrieve user by email %w", err)
	}

	if phone.Valid {
		user.PhoneNumber = phone.String
	}

	return user, nil
}