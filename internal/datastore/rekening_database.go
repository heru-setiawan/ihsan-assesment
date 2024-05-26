package datastore

import (
	"assesment/internal/core/models"
	"assesment/pkg/exceptions"
	"context"
)

func (d *database) Transaksi(ctx context.Context, noRekening string, nominal float64) (*models.Rekening, error) {
	tx := d.db.WithContext(ctx).Begin()

	rekening := new(models.Rekening)
	rekening.No = noRekening

	if err := tx.Where(&rekening).First(&rekening).Error; err != nil {
		tx.Rollback()
		return nil, exceptions.Database{
			Code:    400,
			Message: "rekening tidak ditemukan",
		}
	}

	rekening.Saldo += nominal
	if rekening.Saldo < 0 {
		return nil, exceptions.Database{
			Code:    400,
			Message: "saldo tidak cukup",
		}
	}

	if err := tx.Save(&rekening).Error; err != nil {
		tx.Rollback()
		return nil, exceptions.Database{
			Code:    400,
			Message: "transaksi gagal",
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, exceptions.Database{
			Code:    400,
			Message: "transaksi gagal",
		}
	}

	return rekening, nil
}

func (d *database) CekSaldo(ctx context.Context, noRekening string) (*models.Rekening, error) {
	rekening := new(models.Rekening)
	rekening.No = noRekening

	if err := d.db.WithContext(ctx).Where(&rekening).First(&rekening).Error; err != nil {
		return nil, exceptions.Database{
			Code:    400,
			Message: "rekening tidak ditemukan",
		}
	}

	return rekening, nil
}
