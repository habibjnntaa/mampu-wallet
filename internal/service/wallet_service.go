package service

import (
	"context"
	"errors"
	"mampu-wallet/internal/domain"
	"net/http"
)

type walletService struct {
	repo domain.WalletRepository
	ctx  context.Context
}

func NewWalletService(repo domain.WalletRepository, ctx context.Context) domain.WalletService {
	return &walletService{repo, ctx}
}

func (s *walletService) GetBalance(userID int64) (float64, int, error) {
	wallet, code, err := s.repo.GetByUserID(s.ctx, userID)
	if err != nil {
		return 0, code, err
	}

	return wallet.Balance, code, nil
}

func (s *walletService) Withdraw(form *domain.Withdraw) (float64, int, error) {
	wallet, code, err := s.repo.GetByUserID(s.ctx, form.UserID)
	if err != nil {
		return 0, code, err
	}

	if wallet.Balance < form.Amount {
		return 0, http.StatusBadRequest, errors.New("saldo tidak mencukupi")
	}

	newBalance := wallet.Balance - form.Amount
	return s.repo.UpdateBalance(s.ctx, form.UserID, newBalance)
}
