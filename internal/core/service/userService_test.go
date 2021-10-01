package service_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/pablonoras/money_transfer_service/internal/core/domain"
	"github.com/pablonoras/money_transfer_service/internal/core/service"
	"gopkg.in/go-playground/assert.v1"

	"github.com/golang/mock/gomock"
	"github.com/pablonoras/money_transfer_service/mocks"
)

type dependencies struct {
	transactionRepo *mocks.MockTransactionRepository
	userRepo        *mocks.MockUserRepository
}

func makeDependencies(t *testing.T) dependencies {
	return dependencies{
		transactionRepo: mocks.NewMockTransactionRepository(gomock.NewController(t)),
		userRepo:        mocks.NewMockUserRepository(gomock.NewController(t)),
	}
}

func TestUserService_GetBalance(t *testing.T) {

	ctx := context.Background()
	user := domain.User{
		ID:      "11",
		Balance: 1000,
		Site:    "Arg",
	}

	type args struct {
		context context.Context
		userID  string
	}

	type want struct {
		user *domain.User
		err  error
	}

	tests := []struct {
		name string
		args args
		want want
		mock func(m dependencies, args args, want want)
	}{
		{
			name: "Get access error",
			mock: func(m dependencies, args args, want want) {
				m.userRepo.EXPECT().GetAccess(args.context, args.userID).Return(fmt.Errorf("time-out"))
			},
			args: args{
				context: ctx,
				userID:  "11",
			},
			want: want{
				user: nil,
				err:  fmt.Errorf("time-out"),
			},
		},
		{
			name: "Get user error",
			mock: func(m dependencies, args args, want want) {
				m.userRepo.EXPECT().GetAccess(args.context, args.userID).Return(nil)
				m.userRepo.EXPECT().GetUser(args.context, args.userID).Return(nil, fmt.Errorf("time-out"))
			},
			args: args{
				context: ctx,
				userID:  "11",
			},
			want: want{
				user: nil,
				err:  fmt.Errorf("time-out"),
			},
		},
		{
			name: "Get user ok",
			mock: func(m dependencies, args args, want want) {
				m.userRepo.EXPECT().GetAccess(args.context, args.userID).Return(nil)
				m.userRepo.EXPECT().GetUser(args.context, args.userID).Return(&user, nil)
			},
			args: args{
				context: ctx,
				userID:  "11",
			},
			want: want{
				user: &user,
				err:  fmt.Errorf("time-out"),
			},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			//
			// Setup
			//

			args, want := tt.args, tt.want
			m := makeDependencies(t)
			tt.mock(m, args, want)

			service := service.NewUserService(m.transactionRepo, m.userRepo)

			//
			// Execute
			//

			user, err := service.GetBalance(args.context, args.userID)

			//
			// Verify
			//

			if err != nil && want.err != nil {
				assert.Equal(t, want.err.Error(), err.Error())
			} else {
				assert.Equal(t, want.user, user)
			}
		})
	}
}
