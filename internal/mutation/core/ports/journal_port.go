package ports

import (
	"assesment/internal/mutation/core/models"
	"context"
)

type JournalApp interface {
	AddEntry(ctx context.Context, data models.Journal) error
}

type JournalDatastore interface {
	AddEntry(ctx context.Context, data models.Journal) error
}
