package datastore

import (
	"assesment/pkg/logs"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func New(log logs.Logger, postgresDB *gorm.DB, redisDB *redis.Client) *datastore {
	return &datastore{
		log:        log,
		postgresDB: postgresDB,
		redisDB:    redisDB,
	}
}

type datastore struct {
	log logs.Logger

	postgresDB *gorm.DB
	redisDB    *redis.Client
}
