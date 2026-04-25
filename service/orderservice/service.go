package orderservice

import (
	"context"
	"fmt"

	"github.com/HosseinForouzan/E-Commerce-API/param"
	"github.com/HosseinForouzan/E-Commerce-API/service/cartservice"
	"github.com/HosseinForouzan/E-Commerce-API/service/productservice"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	CreateTx(ctx context.Context, tx pgx.Tx, userID uint, total uint) (uint, error)
	CreateItemTx(ctx context.Context, tx pgx.Tx, orderID uint, item param.CartItemResponse) error
}

type Service struct {
	DB *pgxpool.Pool
	OrderRepo Repository
	CartRepo cartservice.Repository
	ProductRepo productservice.Repository
}

func New(orderRepo Repository, cartRepo cartservice.Repository, productRepo productservice.Repository, db *pgxpool.Pool) Service {
	return Service{
		OrderRepo: orderRepo,
		CartRepo: cartRepo,
		ProductRepo: productRepo,
		DB: db,
	}
}

func (s Service) Checkout(ctx context.Context, userID uint) (param.OrderResponse, error) {
	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return param.OrderResponse{}, fmt.Errorf("unexpected error: %w", err)
	}
	defer tx.Rollback(ctx)

	items, err := s.CartRepo.GetItemsTx(ctx, tx, userID)
	if err != nil {
		return param.OrderResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	if len(items.Items) == 0 {
		return param.OrderResponse{}, fmt.Errorf("Cart is empty.")
	}

	total := 0
	for _, item := range items.Items {
		total += int(item.Price) * int(item.Quantity)
	}

	orderID, err := s.OrderRepo.CreateTx(ctx, tx, userID, uint(total))
	if err != nil {
		return param.OrderResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	for _, item := range items.Items {
		err = s.OrderRepo.CreateItemTx(ctx, tx, orderID, item)
		if err != nil {
			return param.OrderResponse{}, fmt.Errorf("unexpected error: %w", err)
		}

		err = s.ProductRepo.DecreaseStockTx(ctx, tx, item.ProductID, item.Quantity)
		if err != nil {
			return param.OrderResponse{}, fmt.Errorf("unexpected error: %w", err)
		}
	}

	err = s.CartRepo.ClearTx(ctx, tx, userID)
	if err != nil {
		return param.OrderResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return param.OrderResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	return param.OrderResponse{
		ID: int(orderID),
		Status: "pending_payment",
		Total: total,
	}, nil


}