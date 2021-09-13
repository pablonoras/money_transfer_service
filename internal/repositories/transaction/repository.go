package transaction

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type mySQLTransactionRepo struct {
	client *sql.DB
}

func NewMySQLRepository(client *sql.DB) *mySQLTransactionRepo{
	return &mySQLTransactionRepo{
		client: client,
	}
}