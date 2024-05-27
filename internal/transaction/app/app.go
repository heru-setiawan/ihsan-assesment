package app

import (
	"assesment/internal/transaction/core/ports"
	"assesment/pkg/hash"
	"assesment/pkg/logs"

	"go.opentelemetry.io/otel/trace"
)

func New(log logs.Logger, trace trace.Tracer, hash hash.Bcrypt, accountDatabase ports.AccountDatastore) *app {
	return &app{
		log:             log,
		trace:           trace,
		accountDatabase: accountDatabase,
	}
}

type app struct {
	log   logs.Logger
	trace trace.Tracer
	hash  hash.Bcrypt

	accountDatabase ports.AccountDatastore
}
