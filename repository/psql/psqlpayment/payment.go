package psqlpayment

import (
	"context"
	"errors"
	"fmt"

	"github.com/HosseinForouzan/E-Commerce-API/entity"
	"github.com/jackc/pgx/v5"
)

func (d *DB) Create(ctx context.Context, orderID uint, totalAmount uint) (uint, error) {
	var id uint
	err := d.conn.Conn().QueryRow(ctx,
		`INSERT INTO payments (order_id,amount,status,provider) VALUES ($1,$2,'pending','mock') RETURNING id`, orderID, totalAmount).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("can't create payment:%w", err)
	}

	return id, nil
}

func (d *DB) GetByID(ctx context.Context, paymentID uint) (entity.Payment, error) {
	var p entity.Payment

	err := d.conn.Conn().QueryRow(
		ctx,
		`
		SELECT
			id,
			order_id,
			amount,
			status,
			provider,
			transaction_id
		FROM payments
		WHERE id = $1
		`,
		paymentID,
	).Scan(
		&p.ID,
		&p.OrderID,
		&p.Amount,
		&p.Status,
		&p.Provider,
		&p.TransactionID,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Payment{}, nil
		}
		return entity.Payment{}, err
	}

	return p, nil
}

func (d *DB) MarkSuccessTx(
	ctx context.Context,
	tx pgx.Tx,
	paymentID uint,
	transactionID string,
) error {

	cmd, err := tx.Exec(
		ctx,
		`
		UPDATE payments
		SET status='success',
		    transaction_id=$1,
		    updated_at=NOW()
		WHERE id=$2
		  AND status='pending'
		`,
		transactionID,
		paymentID,
	)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return errors.New("payment already processed")
	}

	return nil
}


func (d *DB) MarkFailed(
	ctx context.Context,
	paymentID uint,
) error {

	cmd, err := d.conn.Conn().Exec(
		ctx,
		`
		UPDATE payments
		SET status='failed',
		    updated_at=NOW()
		WHERE id=$1
		  AND status='pending'
		`,
		paymentID,
	)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return errors.New("payment already processed")
	}

	return nil
}
