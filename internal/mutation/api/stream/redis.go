package stream

import (
	"assesment/internal/mutation/core/models"
	"assesment/internal/mutation/core/ports"
	"context"
	"log"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

func NewRedis(redis *redis.Client, journalApp ports.JournalApp) (*redisStream, error) {
	stream := new(redisStream)
	stream.redis = redis
	stream.journalApp = journalApp

	stream.createGroup()

	return stream, nil
}

type redisStream struct {
	redis *redis.Client

	journalApp ports.JournalApp
}

func (s *redisStream) createGroup() {
	s.redis.XGroupCreate(
		context.Background(),
		"transaction",
		"mutation",
		"0",
	)
}

func (s *redisStream) Start() {
	for {
		entries, err := s.redis.XRead(context.Background(), &redis.XReadArgs{
			Streams: []string{"transaction", "0-0"},
			Count:   2,
			Block:   0,
		}).Result()

		if err != nil {
			log.Fatal(err)
		}

		s.handler(entries)
	}
}

func (s *redisStream) handler(entries []redis.XStream) {
	for i := 0; i < len(entries[0].Messages); i++ {
		messageID := entries[0].Messages[i].ID
		values := entries[0].Messages[i].Values

		data := new(models.Journal)

		data.AccountNumber = values["account_number"].(string)

		valueAmount := values["amount"].(string)
		amount, _ := strconv.ParseFloat(valueAmount, 64)
		data.Amount = amount

		valueDate := values["date"].(string)
		date, _ := time.Parse("2006-01-02 15:04:05", valueDate)
		data.Date = date

		valueIsDeposit := values["is_deposit"].(string)
		isDeposit, _ := strconv.ParseBool(valueIsDeposit)
		data.IsDeposit = isDeposit

		s.journalApp.AddEntry(context.Background(), *data)
		s.redis.XDel(context.Background(), "transaction", messageID)
	}
}
