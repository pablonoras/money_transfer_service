package ports

import (
	"context"
	"github.com/pablonoras/money_transfer_service/internal/core/domain"
)

type TransactionRepository interface {
	Get(ctx context.Context, userID string) ([]domain.Transaction, error)
	Create(ctx context.Context, transaction domain.Transaction) error
}

type UserRepository interface {
	GetAccess(ctx context.Context, userID string) error
	UpdateBalance(ctx context.Context, userID string, amount int) error
	GetUser(ctx context.Context, userID string)(*domain.User, error)
}
