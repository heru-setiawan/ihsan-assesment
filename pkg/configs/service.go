package configs

import (
	"os"
	"strconv"

	"github.com/fatih/structs"
)

type TransactionService struct {
	Port uint16 `default:"8000"`
}

func (cfg *TransactionService) defaultValue(key string) string {
	s := structs.New(cfg)
	field := s.Field(key)
	if field == nil {
		return ""
	}
	return field.Tag("default")
}

func (cfg *TransactionService) loadEnv() error {
	osPort := os.Getenv("SERVICE_TRANSACTION_PORT")
	if osPort == "" {
		osPort = cfg.defaultValue("Port")
	}

	port, err := strconv.Atoi(osPort)
	if err != nil {
		return err
	}
	cfg.Port = uint16(port)

	return nil
}
