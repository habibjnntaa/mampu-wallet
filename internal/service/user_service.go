package service

import (
	"context"
	"mampu-wallet/internal/domain"
	"mampu-wallet/internal/tools"
	"net/http"
)

type userService struct {
	repo domain.UserRepository
	ctx  context.Context
}

func NewUserService(repo domain.UserRepository, ctx context.Context) domain.UserService {
	return &userService{repo, ctx}
}

func (s *userService) GetAllUser(pagination *tools.Pagination) ([]*domain.User, *tools.Pagination, int, error) {
	result, pagination, err := s.repo.GetAllUser(s.ctx, pagination)
	if err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}

	return result, pagination, http.StatusOK, nil
}
