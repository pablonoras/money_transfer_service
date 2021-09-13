package ports

import (
	"context"
	"github.com/pablonoras/money_transfer_service/internal/core/domain"
)

type TransactionService interface {
	Get(ctx context.Context, userID string) ([]domain.Transaction, error)
	Create(ctx context.Context, userID string, receptorID string, amount int) (*domain.Transaction, error)
}

type UserService interface {
	UpdateBalance(ctx context.Context, user domain.User, transaction domain.Transaction) error
	GetBalance(ctx context.Context, userID string) (*domain.User, error)
}

