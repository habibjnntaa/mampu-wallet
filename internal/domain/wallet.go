package domain

import (
	"context"
)

type Wallet struct {
	ID      int64   `json:"id"`
	UserID  int64   `json:"user_id"`
	Balance float64 `json:"balance"`
}

type Withdraw struct {
	UserID int64   `json:"user_id" binding:"required"`
	Amount float64 `json:"amount" binding:"required"`
}

type WalletRepository interface {
	GetByUserID(ctx context.Context, userID int64) (*Wallet, int, error)
	UpdateBalance(ctx context.Context, userID int64, newBalance float64) (float64, int, error)
}

type WalletService interface {
	GetBalance(userID int64) (float64, int, error)
	Withdraw(form *Withdraw) (float64, int, error)
}
