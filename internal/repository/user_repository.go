package repository

import (
	"context"
	"database/sql"
	"mampu-wallet/internal/domain"
	"mampu-wallet/internal/tools"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetAllUser(ctx context.Context, pagination *tools.Pagination) ([]*domain.User, *tools.Pagination, error) {
	query := `
		SELECT id, name, email 
		FROM users 
		ORDER BY id ASC 
		LIMIT $1 OFFSET $2`

	rows, err := r.db.QueryContext(ctx, query, pagination.Limit, pagination.Offset)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var allUsers []*domain.User
	for rows.Next() {
		var u domain.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, nil, err
		}
		allUsers = append(allUsers, &u)
	}

	// Cek error setelah loop selesai
	if err = rows.Err(); err != nil {
		return nil, nil, err
	}

	var total int
	query = "SELECT COUNT(id) FROM users"
	err = r.db.QueryRowContext(ctx, query).Scan(&total)
	if err != nil {
		return nil, nil, err
	}

	pagination.Count = total
	pagination = tools.Paging(pagination)

	return allUsers, pagination, nil
}
