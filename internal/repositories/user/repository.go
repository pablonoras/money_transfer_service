package user

import (
	"database/sql"
	"net/http"
)

type userRepository struct {
	httpClient http.Client
	mySqlClient *sql.DB
}

func NewUserRepository(client http.Client, mySqlClient *sql.DB) *userRepository{
	return &userRepository{
		httpClient: client,
		mySqlClient: mySqlClient,
	}
}