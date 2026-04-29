package paymentservice

import (
	"context"
	"errors"
	"fmt"

	"github.com/HosseinForouzan/E-Commerce-API/entity"
	"github.com/HosseinForouzan/E-Commerce-API/param"
	"github.com/HosseinForouzan/E-Commerce-API/service/orderservice"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	Create(ctx context.Context, orderID uint, totalAmount uint) (uint, error)
	MarkSuccessTx(ctx context.Context, tx pgx.Tx, paymentID uint ,transactionID string) error
	GetByID(ctx context.Context, paymentID uint) (entity.Payment, error)
	MarkFailed(ctx context.Context ,paymentID uint) error
	
}

type Service struct {
	PaymentRepo Repository
	OrderRepo orderservice.Repository
	DB *pgxpool.Pool
}

func New(paymentRepo Repository, orderRepo orderservice.Repository, db *pgxpool.Pool) Service {
	return Service{
		PaymentRepo: paymentRepo,
		OrderRepo: orderRepo ,
		DB: db,
	}
}

func (s *Service) StartPayment(
	ctx context.Context,
	req param.PaymentRequest,
) (string, error) {

	order, err := s.OrderRepo.GetByID(ctx, req.OrderID, req.UserID)
	if err != nil {
		return "", err
	}

	if order.ID == 0 {
		return "", errors.New("order not found")
	}

	if order.Status != "pending_payment" {
		return "", errors.New("order is not payable")
	}

	paymentID, err := s.PaymentRepo.Create(
		ctx,
		req.OrderID,
		uint(order.TotalAmount),
	)

	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("/payment/mock/%d", paymentID)

	return url, nil
}


func (s *Service) Success(
	ctx context.Context,
	paymentID uint,
) error {

	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	payment, err := s.PaymentRepo.GetByID(ctx, paymentID)
	if err != nil {
		return err
	}

	if payment.ID == 0 {
		return errors.New("payment not found")
	}

	if payment.Status != "pending" {
		return errors.New("payment already processed")
	}

	transactionID := uuid.New().String()

	err = s.PaymentRepo.MarkSuccessTx(
		ctx,
		tx,
		paymentID,
		transactionID,
	)

	if err != nil {
		return err
	}

	err = s.OrderRepo.MarkPaidTx(
		ctx,
		tx,
		payment.OrderID,
	)

	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}



func (s Service) Fail(
	ctx context.Context,
	paymentID uint,
) error {

	return s.PaymentRepo.MarkFailed(
		ctx,
		paymentID,
	)
}