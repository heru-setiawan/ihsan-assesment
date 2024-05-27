package configs

import "github.com/joho/godotenv"

type Config struct {
	HashBcrypt         Bcrypt
	DatabaseRedis      Redis
	DatabasePostgres   Postgres
	LogLogrus          Logrus
	TraceTelemetry     Telemetry
	ServiceTransaction TransactionService
}

func (cfg *Config) LoadFromEnvFile(files ...string) error {
	if len(files) != 0 {
		for _, file := range files {
			if err := godotenv.Load(file); err != nil {
				return err
			}
		}
	} else {
		godotenv.Load()
	}

	if err := cfg.DatabaseRedis.loadEnv(); err != nil {
		return err
	}

	if err := cfg.DatabasePostgres.loadEnv(); err != nil {
		return err
	}

	if err := cfg.LogLogrus.loadEnv(); err != nil {
		return err
	}

	if err := cfg.TraceTelemetry.loadEnv(); err != nil {
		return err
	}

	if err := cfg.ServiceTransaction.loadEnv(); err != nil {
		return err
	}

	if err := cfg.ServiceTransaction.loadEnv(); err != nil {
		return err
	}

	if err := cfg.HashBcrypt.loadEnv(); err != nil {
		return err
	}

	return nil
}
