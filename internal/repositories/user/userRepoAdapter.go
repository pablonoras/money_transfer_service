package user

import (
	"context"
	"money_transfer_service/internal/core/domain"
)

func (repo *userRepository)GetAccess(ctx context.Context, userID string) error {
	//TODO: Apicall to the authenticated HTTP-based API to valid userID
	return nil
}

func (repo *userRepository) UpdateBalance(ctx context.Context, userID string, balance int) error{
	_,err :=repo.mySqlClient.Exec( "UPDATE users SET balance =? WHERE user_id =?;", balance, userID)
	if err != nil{
		return err
	}

	return nil
}

func (repo *userRepository) GetUser(ctx context.Context, userID string)(*domain.User, error) {
	rows, err := repo.mySqlClient.QueryContext(ctx, "SELECT *  FROM 	users WHERE user_id =?", userID)
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	user:= domain.User{}

	for rows.Next() {

		if err := rows.Scan(&user.ID, &user.Site, &user.Balance); err != nil {
			return nil, err
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &user, nil
}
