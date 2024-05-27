package app

import (
	"assesment/internal/transaction/core/models"
	"assesment/pkg/exceptions"
	"context"

	"go.opentelemetry.io/otel/attribute"
)

func (a *app) Register(ctx context.Context, pin string) (*models.Account, error) {
	ctx, span := a.trace.Start(ctx, "app/register")
	span.SetAttributes(attribute.String("pin", pin))
	defer span.End()

	if pin == "" {
		err := exceptions.Validation{
			Code:    400,
			Message: "pin is required",
		}
		span.RecordError(err)
		a.log.Warn(map[string]any{"pin": pin}, "validation error")
		return nil, err
	}

	newAccount := new(models.Account)
	newAccount.PIN = a.hash.Encrypt(pin)
	newAccount.GenerateNumber()

	err := a.accountDatabase.Register(ctx, newAccount)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	a.log.Info(nil, "register success")
	return newAccount, nil
}

func (a *app) CheckAccount(ctx context.Context, accountNumber string, pin string) error {
	ctx, span := a.trace.Start(ctx, "app/deposit")
	span.SetAttributes(attribute.String("account_number", accountNumber))
	span.SetAttributes(attribute.String("pin", pin))
	defer span.End()

	if accountNumber == "" {
		err := exceptions.Validation{
			Code:    400,
			Message: "account number is required",
		}
		span.RecordError(err)
		a.log.Warn(map[string]any{"account_number": accountNumber, "pin": pin}, "validation error")
		return err
	}

	if pin == "" {
		err := exceptions.Validation{
			Code:    400,
			Message: "pin is required",
		}
		span.RecordError(err)
		a.log.Warn(map[string]any{"account_number": accountNumber, "pin": pin}, "validation error")
		return err
	}

	data, err := a.accountDatabase.GetAccountByNumber(ctx, accountNumber)
	if err != nil {
		return err
	}

	if !a.hash.Compare(pin, data.PIN) {
		err := exceptions.Validation{
			Code:    400,
			Message: "invalid credential",
		}
		span.RecordError(err)
		a.log.Warn(map[string]any{"account_number": accountNumber, "pin": pin}, "compare pin error")
		return err
	}

	return nil
}

func (a *app) Deposit(ctx context.Context, accountNumber string, amount float64) (*models.Account, error) {
	ctx, span := a.trace.Start(ctx, "app/deposit")
	span.SetAttributes(attribute.String("account_number", accountNumber))
	span.SetAttributes(attribute.Float64("amount", amount))
	defer span.End()

	if accountNumber == "" {
		err := exceptions.Validation{
			Code:    400,
			Message: "account number is required",
		}
		span.RecordError(err)
		a.log.Warn(map[string]any{"account_number": accountNumber, "amount": amount}, "validation error")
		return nil, err
	}

	if amount < 0 {
		err := exceptions.Validation{
			Code:    400,
			Message: "amount must be greater than 0",
		}
		span.RecordError(err)
		a.log.Warn(map[string]any{"account_number": accountNumber, "amount": amount}, "validation error")
		return nil, err
	}

	data, err := a.accountDatabase.Transaction(ctx, true, accountNumber, amount)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	a.log.Info(nil, "deposit success")
	return data, nil
}

func (a *app) Withdraw(ctx context.Context, accountNumber string, amount float64) (*models.Account, error) {
	ctx, span := a.trace.Start(ctx, "app/withdraw")
	span.SetAttributes(attribute.String("account_number", accountNumber))
	span.SetAttributes(attribute.Float64("amount", amount))
	defer span.End()

	if accountNumber == "" {
		err := exceptions.Validation{
			Code:    400,
			Message: "account number is required",
		}
		span.RecordError(err)
		a.log.Warn(map[string]any{"account_number": accountNumber, "amount": amount}, "validation error")
		return nil, err
	}

	if amount < 0 {
		err := exceptions.Validation{
			Code:    400,
			Message: "amount must be greater than 0",
		}
		span.RecordError(err)
		a.log.Warn(map[string]any{"account_number": accountNumber, "amount": amount}, "validation error")
		return nil, err
	}

	amount = amount * -1

	data, err := a.accountDatabase.Transaction(ctx, false, accountNumber, amount)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	a.log.Info(nil, "withdraw success")
	return data, nil
}
