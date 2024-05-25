package configs

import (
	"fmt"
	"os"
	"strconv"
)

type telemetry struct {
	Host string
	Port uint16
}

func (t *telemetry) loadFromEnv(defaultValue telemetry) {
	osHost := os.Getenv("TELEMETRY_HOST")
	if osHost == "" {
		osHost = defaultValue.Host
	}
	t.Host = osHost

	osPort := os.Getenv("TELEMETRY_PORT")
	if osPort == "" {
		osPort = fmt.Sprintf("%d", defaultValue.Port)
	}

	port, err := strconv.Atoi(osPort)
	if err != nil {
		port = int(defaultValue.Port)
	}
	t.Port = uint16(port)
}
