package service

import (
	"context"
	"fmt"
	"github.com/oklog/ulid"
	"github.com/pablonoras/money_transfer_service/internal/core/domain"
	"github.com/pablonoras/money_transfer_service/internal/core/ports"
	"math/rand"
	"time"
)

type transactionService struct {
	transactionRepo ports.TransactionRepository
	userRepo ports.UserRepository
	userService ports.UserService
}

func NewTransactionService(
	transactionRepo ports.TransactionRepository,
	userRepo ports.UserRepository,
	userService ports.UserService) *transactionService {

	return &transactionService{
		transactionRepo: transactionRepo,
		userRepo: userRepo,
		userService: userService,
	}
}

func (srv *transactionService) Get(ctx context.Context, userID string) ([]domain.Transaction, error) {

	err := srv.userRepo.GetAccess(ctx, userID)
	if err != nil{
		return nil, err
	}

	transactions, err :=srv.transactionRepo.Get(ctx, userID)
	if err != nil {
		// TODO: structures logging and metrics should be added
		return nil, err
	}
	return transactions, nil
}

func (srv *transactionService) 	Create(ctx context.Context, userID string, receptorID string, amount int) (*domain.Transaction, error){

	if err := srv.userRepo.GetAccess(ctx, userID); err != nil{
		return nil, err
	}

	user, err := srv.userRepo.GetUser(ctx, userID)
	if err != nil{
		// TODO: structured logging and metrics should be added
		return nil, err
	}

	if amount > user.Balance{
		// TODO: structured logging and metrics should be added
		return nil, fmt.Errorf("cannot create the transaction, insufficient funds")
	}

	receptor, err := srv.userRepo.GetUser(ctx, receptorID)
	if err != nil{
		// TODO: structured logging and metrics should be added
		return nil, err
	}

	// TODO: Some business logic must be added in order to validate if this transaction could be processed.
	transaction := buildNewTransaction(*user, receptorID, receptor.Site, amount)

	if err := srv.transactionRepo.Create(ctx, transaction); err != nil{
		// TODO: structured logging and metrics should be added
		return nil, err
	}

	if err := srv.userService.UpdateBalance(ctx, *user, transaction); err != nil {
		// TODO: structured logging and metrics should be added
		return nil, err
	}

	// TODO: structured logging and metrics should be added
	return &transaction, nil
}

func buildNewTransaction(user domain.User, receptorID string, siteTo string,amount int) domain.Transaction {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return domain.Transaction{
		UserID: user.ID,
		ReceptorID: receptorID,
		TransactionID: id.String(),
		Amount: amount,
		SiteFrom: user.Site,
		SiteTo: siteTo,
		CreationDate: time.Now().UTC().String(),
		LastModifiedDate: time.Now().UTC().String(),
		Status: "pending",
	}
}