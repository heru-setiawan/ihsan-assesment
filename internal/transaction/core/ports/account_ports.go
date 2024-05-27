package ports

import (
	"assesment/internal/transaction/core/models"
	"context"
)

type AccountApp interface {
	Register(ctx context.Context, pin string) (*models.Account, error)
	CheckAccount(ctx context.Context, accountNumber string, pin string) error

	Deposit(ctx context.Context, accountNumber string, amount float64) (*models.Account, error)
	Withdraw(ctx context.Context, accountNumber string, amount float64) (*models.Account, error)
}

type AccountDatastore interface {
	GetAccountByNumber(ctx context.Context, accountNumber string) (*models.Account, error)
	Register(ctx context.Context, data *models.Account) error
	Transaction(ctx context.Context, isDeposit bool, accountNumber string, amount float64) (*models.Account, error)
}
