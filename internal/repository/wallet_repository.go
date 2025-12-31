package repository

import (
	"context"
	"database/sql"
	"fmt"
	"mampu-wallet/internal/domain"
	"net/http"
)

type walletRepository struct {
	db *sql.DB
}

func NewWalletRepository(db *sql.DB) domain.WalletRepository {
	return &walletRepository{db}
}

func (r *walletRepository) GetByUserID(ctx context.Context, userID int64) (*domain.Wallet, int, error) {
	var w domain.Wallet
	query := "SELECT id, user_id, balance FROM wallets WHERE user_id = $1"
	err := r.db.QueryRowContext(ctx, query, userID).Scan(&w.ID, &w.UserID, &w.Balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, http.StatusNotFound, fmt.Errorf("wallet tidak ditemukan untuk user id %d", userID)
		}

		return nil, http.StatusInternalServerError, fmt.Errorf("wallet tidak ditemukan untuk user id %d", userID)
	}
	return &w, http.StatusOK, err
}

func (r *walletRepository) UpdateBalance(ctx context.Context, userID int64, newBalance float64) (float64, int, error) {
	query := "UPDATE wallets SET balance = $1 WHERE user_id = $2"
	_, err := r.db.ExecContext(ctx, query, newBalance, userID)
	if err != nil {
		return newBalance, http.StatusBadGateway, err
	}

	return newBalance, http.StatusAccepted, err
}
