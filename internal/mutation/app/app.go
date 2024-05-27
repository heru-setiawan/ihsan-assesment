package app

import (
	"assesment/internal/mutation/core/ports"
	"assesment/pkg/logs"

	"go.opentelemetry.io/otel/trace"
)

func New(log logs.Logger, trace trace.Tracer, journalDatastore ports.JournalDatastore) *app {
	return &app{
		log:              log,
		trace:            trace,
		journalDatastore: journalDatastore,
	}
}

type app struct {
	log   logs.Logger
	trace trace.Tracer

	journalDatastore ports.JournalDatastore
}
