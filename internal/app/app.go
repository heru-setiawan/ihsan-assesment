package app

import (
	"assesment/internal/core/ports"
	"assesment/pkg/logs"
)

func New(log logs.Logger, nasabahDatabase ports.NasabahDatabase, rekeningDatabase ports.RekeningDatabase) *app {
	return &app{
		log:              log,
		nasabahDatabase:  nasabahDatabase,
		rekeningDatabase: rekeningDatabase,
	}
}

type app struct {
	log logs.Logger

	nasabahDatabase  ports.NasabahDatabase
	rekeningDatabase ports.RekeningDatabase
}
