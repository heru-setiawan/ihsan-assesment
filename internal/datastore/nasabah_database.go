package datastore

import (
	"assesment/internal/core/models"
	"assesment/pkg/exceptions"
	"context"
	"strings"
)

func (d *database) Daftar(ctx context.Context, data models.Nasabah) (*models.Nasabah, error) {
	if err := d.db.WithContext(ctx).Create(&data).Error; err != nil {
		if strings.Contains(err.Error(), "nasabah_pkey") {
			return nil, exceptions.Database{
				Code:    400,
				Message: "nama sudah digunakan",
			}
		}

		if strings.Contains(err.Error(), "no_hp") {
			return nil, exceptions.Database{
				Code:    400,
				Message: "no hp sudah digunakan",
			}
		}

		return nil, exceptions.Database{
			Code:    400,
			Message: "pendaftaran nasabah gagal",
		}
	}

	return &data, nil
}
