package datastore

import (
	"assesment/pkg/logs"

	"gorm.io/gorm"
)

func New(log logs.Logger, postgresDB *gorm.DB) *datastore {
	return &datastore{
		log:        log,
		postgresDB: postgresDB,
	}
}

type datastore struct {
	log logs.Logger

	postgresDB *gorm.DB
}
