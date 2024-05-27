package configs

import (
	"os"
	"strconv"

	"github.com/fatih/structs"
)

type Postgres struct {
	Host     string `default:"localhost"`
	Port     uint16 `default:"5432"`
	User     string `default:"postgres"`
	Password string `default:""`
	Database string `default:"postgres"`
	Schema   string `default:"public"`
}

func (cfg *Postgres) defaultValue(key string) string {
	s := structs.New(cfg)
	field := s.Field(key)
	if field == nil {
		return ""
	}
	return field.Tag("default")
}

func (cfg *Postgres) loadEnv() error {
	osHost := os.Getenv("DB_POSTGRES_HOST")
	if osHost == "" {
		osHost = cfg.defaultValue("Host")
	}
	cfg.Host = osHost

	osPort := os.Getenv("DB_POSTGRES_PORT")
	if osPort == "" {
		osPort = cfg.defaultValue("Port")
	}

	port, err := strconv.Atoi(osPort)
	if err != nil {
		return err
	}
	cfg.Port = uint16(port)

	osUser := os.Getenv("DB_POSTGRES_USER")
	if osUser == "" {
		osUser = cfg.defaultValue("User")
	}
	cfg.User = osUser

	osPassword := os.Getenv("DB_POSTGRES_PASSWORD")
	if osPassword == "" {
		osPassword = cfg.defaultValue("Password")
	}
	cfg.Password = osPassword

	osDatabase := os.Getenv("DB_POSTGRES_DATABASE")
	if osDatabase == "" {
		osDatabase = cfg.defaultValue("Database")
	}
	cfg.Database = osDatabase

	osSchema := os.Getenv("DB_POSTGRES_SCHEMA")
	if osSchema == "" {
		osSchema = cfg.defaultValue("Schema")
	}
	cfg.Schema = osSchema

	return nil
}

type Redis struct {
	Host     string `default:"localhost"`
	Port     uint16 `default:"6379"`
	User     string `default:"default"`
	Password string `default:""`
}

func (cfg *Redis) defaultValue(key string) string {
	s := structs.New(cfg)
	field := s.Field(key)
	if field == nil {
		return ""
	}
	return field.Tag("default")
}

func (cfg *Redis) loadEnv() error {
	osHost := os.Getenv("DB_REDIS_HOST")
	if osHost == "" {
		osHost = cfg.defaultValue("Host")
	}
	cfg.Host = osHost

	osPort := os.Getenv("DB_REDIS_PORT")
	if osPort == "" {
		osPort = cfg.defaultValue("Port")
	}

	port, err := strconv.Atoi(osPort)
	if err != nil {
		return err
	}
	cfg.Port = uint16(port)

	osUser := os.Getenv("DB_REDIS_USER")
	if osUser == "" {
		osUser = cfg.defaultValue("User")
	}
	cfg.User = osUser

	osPassword := os.Getenv("DB_REDIS_PASSWORD")
	if osPassword == "" {
		osPassword = cfg.defaultValue("Password")
	}
	cfg.Password = osPassword

	return nil
}
