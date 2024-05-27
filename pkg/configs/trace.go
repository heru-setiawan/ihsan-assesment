package configs

import (
	"os"
	"strconv"

	"github.com/fatih/structs"
)

type Telemetry struct {
	Host string `default:"localhost"`
	Port uint16 `default:"6831"`
}

func (cfg *Telemetry) defaultValue(key string) string {
	s := structs.New(cfg)
	field := s.Field(key)
	if field == nil {
		return ""
	}
	return field.Tag("default")
}

func (cfg *Telemetry) loadEnv() error {
	osHost := os.Getenv("TRACE_TELEMETRY_HOST")
	if osHost == "" {
		osHost = cfg.defaultValue("Host")
	}
	cfg.Host = osHost

	osPort := os.Getenv("TRACE_TELEMETRY_PORT")
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
