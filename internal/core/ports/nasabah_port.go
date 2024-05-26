package ports

import (
	"assesment/internal/core/models"
	"context"
)

type NasabahApp interface {
	Daftar(ctx context.Context, data models.Nasabah) (*models.Nasabah, error)
}

type NasabahDatabase interface {
	Daftar(ctx context.Context, data models.Nasabah) (*models.Nasabah, error)
}
