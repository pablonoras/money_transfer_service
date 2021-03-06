package transaction

import (
	"context"
	"net/http"

	"github.com/pablonoras/money_transfer_service/internal/core/domain"
	"github.com/pablonoras/money_transfer_service/pkg/errors"
)

func (repo *mySQLTransactionRepo) Get(ctx context.Context, userID string) ([]domain.Transaction, error) {

	rows, err := repo.client.QueryContext(ctx, "SELECT *  FROM transactions WHERE user_id =?", userID)
	if err != nil {
		return nil, errors.NewError(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()

	transactions := make([]domain.Transaction, 0)

	for rows.Next() {
		var transaction domain.Transaction
		var id int
		if err := rows.Scan(
			&id,
			&transaction.TransactionID,
			&transaction.UserID,
			&transaction.ReceptorID,
			&transaction.Amount,
			&transaction.SiteFrom,
			&transaction.SiteTo,
			&transaction.Status,
			&transaction.CreationDate,
			&transaction.LastModifiedDate); err != nil {
			return nil, errors.NewError(http.StatusInternalServerError, err.Error())
		}

		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.NewError(http.StatusInternalServerError, err.Error())
	}

	return transactions, nil
}

func (repo *mySQLTransactionRepo) Create(ctx context.Context, transaction domain.Transaction) error {
	insert, err := repo.client.QueryContext(ctx, "INSERT INTO transactions (transaction_id, user_id , receptor_id , amount , site_from , site_to , creation_date , last_modified_date , status ) VALUES (?,?,?,?,?,?,?,?,?)", transaction.TransactionID, transaction.UserID, transaction.ReceptorID, transaction.Amount, transaction.SiteFrom, transaction.SiteTo, transaction.CreationDate, transaction.LastModifiedDate, transaction.Status)
	if err != nil {
		return errors.NewError(http.StatusInternalServerError, err.Error())
	}
	defer insert.Close()

	return nil
}
