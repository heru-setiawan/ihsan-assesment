package databases

import (
	"assesment/pkg/configs"
	"context"
	"fmt"

	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel/trace"
)

func NewRedis(config configs.Redis, traceProvider trace.TracerProvider) (*redisDB, error) {
	database := new(redisDB)
	database.config = config
	if err := database.connect(traceProvider); err != nil {
		return nil, err
	}

	return database, nil
}

type redisDB struct {
	Conn   *redis.Client
	config configs.Redis
}

func (db *redisDB) connect(traceProvider trace.TracerProvider) error {
	url := fmt.Sprintf("redis://%s:%s@%s:%d", db.config.User, db.config.Password, db.config.Host, db.config.Port)
	opt, err := redis.ParseURL(url)
	if err != nil {
		return err
	}

	db.Conn = redis.NewClient(opt)
	if err := db.Conn.Ping(context.Background()).Err(); err != nil {
		return err
	}

	if err := redisotel.InstrumentTracing(db.Conn); err != nil {
		return err
	}

	if err := redisotel.InstrumentMetrics(db.Conn); err != nil {
		return err
	}

	return nil
}
