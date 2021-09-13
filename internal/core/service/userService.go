package service

import (
	"context"
	"github.com/pablonoras/money_transfer_service/internal/core/domain"
	"github.com/pablonoras/money_transfer_service/internal/core/ports"
)

type userService struct{
	transactionRepo ports.TransactionRepository
	userRepo ports.UserRepository
}

func NewUserService(
	transactionRepo ports.TransactionRepository,
	userRepo ports.UserRepository) *userService {

	return &userService{
		transactionRepo: transactionRepo,
		userRepo: userRepo,
	}
}

func (srv *userService) UpdateBalance(ctx context.Context, user domain.User, transaction domain.Transaction) error{

	err := srv.userRepo.GetAccess(ctx, user.ID)
	if err != nil{
		return nil
	}

	receptor, err := srv.userRepo.GetUser(ctx, transaction.ReceptorID)
	if err != nil {
		// TODO: structured logging and metrics should be added
		return nil
	}

	updatedUserBalance := user.Balance - transaction.Amount
	updatedReceptorBalance := receptor.Balance + transaction.Amount

	err = srv.userRepo.UpdateBalance(ctx, user.ID, updatedUserBalance )
	if err != nil {
		// TODO: structured logging and metrics should be added
		return err
	}

	err = srv.userRepo.UpdateBalance(ctx, transaction.ReceptorID, updatedReceptorBalance)
	if err != nil {
		// TODO: structured logging and metrics should be added
		return err
	}
	return nil
}

func (srv *userService) GetBalance(ctx context.Context, userID string) (*domain.User, error) {
	err := srv.userRepo.GetAccess(ctx, userID)
	if err != nil{
		return nil, err
	}

	user, err := srv.userRepo.GetUser(ctx, userID)
	if err != nil{
		// TODO: structured logging and metrics should be added
		return nil, err
	}
	return user, err
}