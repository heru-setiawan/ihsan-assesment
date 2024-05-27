package datastore

import (
	"assesment/internal/transaction/core/models"
	"assesment/pkg/exceptions"
	"context"
	"math"
	"time"

	"github.com/redis/go-redis/v9"
)

func (d *datastore) GetAccountByNumber(ctx context.Context, accountNumber string) (*models.Account, error) {
	data := new(models.Account)
	if err := d.postgresDB.WithContext(ctx).Where(models.Account{Number: accountNumber}).First(&data).Error; err != nil {
		d.log.Error(map[string]any{"account_number": accountNumber, "error": err.Error()}, "database error")
		return nil, exceptions.Database{
			Code:    404,
			Message: "account not found",
		}
	}

	return data, nil
}

func (d *datastore) Register(ctx context.Context, data *models.Account) error {
	if err := d.postgresDB.WithContext(ctx).Create(&data).Error; err != nil {
		d.log.Error(map[string]any{"data": data, "error": err.Error()}, "database error")
		return exceptions.Database{
			Code:    400,
			Message: "transaction failed",
		}
	}

	return nil
}

func (d *datastore) Transaction(ctx context.Context, isDeposit bool, accountNumber string, amount float64) (*models.Account, error) {
	tx := d.postgresDB.WithContext(ctx).Begin()

	var data models.Account
	if err := tx.Where(models.Account{Number: accountNumber}).First(&data).Error; err != nil {
		d.log.Error(map[string]any{"account_number": accountNumber, "amount": amount, "error": err.Error()}, "database error")
		tx.Rollback()
		return nil, exceptions.Database{
			Code:    400,
			Message: "transaction failed",
		}
	}

	data.Balance += amount
	if data.Balance < 0 {
		d.log.Error(map[string]any{"account_number": accountNumber, "amount": amount, "error": "balance is not enough"}, "validation error")
		tx.Rollback()
		return nil, exceptions.Validation{
			Code:    400,
			Message: "balance is not enough",
		}
	}

	if err := tx.Where(models.Account{Number: accountNumber}).Save(&data).Error; err != nil {
		d.log.Error(map[string]any{"account_number": accountNumber, "amount": amount, "error": err.Error()}, "database error")
		tx.Rollback()
		return nil, exceptions.Validation{
			Code:    400,
			Message: "transaction failed",
		}
	}

	err := d.redisDB.XAdd(ctx, &redis.XAddArgs{
		Stream: "transaction",
		Values: map[string]any{"date": time.Now().Format("2006-01-02 15:04:05"), "is_deposit": isDeposit, "amount": math.Abs(amount), "account_number": accountNumber},
	}).Err()
	if err != nil {
		d.log.Error(map[string]any{"account_number": accountNumber, "amount": amount, "error": err}, "database error")
		tx.Rollback()
		return nil, exceptions.Validation{
			Code:    400,
			Message: "transaction failed",
		}
	}

	if err := tx.Commit().Error; err != nil {
		d.log.Error(map[string]any{"account_number": accountNumber, "amount": amount, "error": err.Error()}, "database error")
		tx.Rollback()
		return nil, exceptions.Validation{
			Code:    400,
			Message: "transaction failed",
		}
	}

	return &data, nil
}
