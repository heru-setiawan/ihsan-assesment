package datastore

import (
	"assesment/internal/mutation/core/models"
	"assesment/pkg/exceptions"
	"context"
)

func (d *datastore) AddEntry(ctx context.Context, data models.Journal) error {
	if err := d.postgresDB.WithContext(ctx).Create(&data).Error; err != nil {
		d.log.Error(map[string]any{"data": data, "error": err.Error()}, "database error")
		return exceptions.Database{Code: 400, Message: "add entry failed"}
	}

	return nil
}
