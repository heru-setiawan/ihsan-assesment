package app

import (
	"assesment/internal/mutation/core/models"
	"context"
	"encoding/json"

	"go.opentelemetry.io/otel/attribute"
)

func (a *app) AddEntry(ctx context.Context, data models.Journal) error {
	ctx, span := a.trace.Start(ctx, "app/add-entry")
	dataString, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		span.RecordError(err)
		a.log.Warn(map[string]any{"data": data, "error": err.Error()}, "marshal error")
	}
	span.SetAttributes(attribute.String("data", string(dataString)))
	defer span.End()

	if err := a.journalDatastore.AddEntry(ctx, data); err != nil {
		span.RecordError(err)
		return err
	}

	return nil
}
