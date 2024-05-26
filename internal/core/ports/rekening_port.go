package ports

import (
	"assesment/internal/core/models"
	"context"
)

type RekeningApp interface {
	Tabung(ctx context.Context, noRekening string, nominal float64) (*models.Rekening, error)
	Tarik(ctx context.Context, noRekening string, nominal float64) (*models.Rekening, error)
	CekSaldo(ctx context.Context, noRekening string) (*models.Rekening, error)
}

type RekeningDatabase interface {
	Transaksi(ctx context.Context, noRekening string, nominal float64) (*models.Rekening, error)
	CekSaldo(ctx context.Context, noRekening string) (*models.Rekening, error)
}
