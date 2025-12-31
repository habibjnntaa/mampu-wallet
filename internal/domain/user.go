package domain

import (
	"context"
	"mampu-wallet/internal/tools"
)

type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserRepository interface {
	GetAllUser(ctx context.Context, pagination *tools.Pagination) ([]*User, *tools.Pagination, error)
}

type UserService interface {
	GetAllUser(pagination *tools.Pagination) ([]*User, *tools.Pagination, int, error)
}
