package cmd

import (
	"database/sql"
	"money_transfer_service/internal/core/ports"
	"money_transfer_service/internal/core/service"
	"money_transfer_service/internal/handlers/transaction"
	"money_transfer_service/internal/handlers/user"
	trxRepo "money_transfer_service/internal/repositories/transaction"
	userRepo  "money_transfer_service/internal/repositories/user"
	"net/http"
	"os"
)

var Dependencies = Definition{}

type Definition struct{
	//
	// Repositories
	//

	TransactionRepository ports.TransactionRepository
	UserRepository ports.UserRepository

	//
	// Core
	//

	TransactionService ports.TransactionService
	UserService ports.UserService

	// Handlers

	TransactionHandler *transaction.HTTPHandler
	UserHandler *user.HTTPHandler
}

func NewByEnvironment() Definition {
	var environment string

	switch os.Getenv("GO_ENVIRONMENT") {
	case "production":
		environment = "prod"
	default:
		environment = "local"
	}

	configs := environmentConfigs[environment]

	mySQLClient, err := initMySQL(configs.mySql["test_mysql"])
	if err != nil {
		panic(err)
	}

	d := initDependencies(mySQLClient)

	return d
}

func initDependencies(mySQLClient *sql.DB) Definition {
	d := Definition{}

	d.TransactionRepository = trxRepo.NewMySQLRepository(mySQLClient)
	d.UserRepository = userRepo.NewUserRepository(http.Client{},mySQLClient)

	d.UserService = service.NewUserService(d.TransactionRepository, d.UserRepository)
	d.TransactionService = service.NewTransactionService(d.TransactionRepository, d.UserRepository, d.UserService)

	d.TransactionHandler = transaction.NewTransactionHandler(d.TransactionService)
	d.UserHandler = user.NewUserHandler(d.UserService)

	return d
}