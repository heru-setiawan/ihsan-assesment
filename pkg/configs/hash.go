package configs

import (
	"os"
	"strconv"

	"github.com/fatih/structs"
)

type Bcrypt struct {
	Salt int `default:"10"`
}

func (cfg *Bcrypt) defaultValue(key string) string {
	s := structs.New(cfg)
	field := s.Field(key)
	if field == nil {
		return ""
	}
	return field.Tag("default")
}

func (cfg *Bcrypt) loadEnv() error {
	osSalt := os.Getenv("HASH_BCRYPT_SALT")
	if osSalt == "" {
		osSalt = cfg.defaultValue("Salt")
	}

	salt, err := strconv.Atoi(osSalt)
	if err != nil {
		return err
	}
	cfg.Salt = salt

	return nil
}
