package configs

import (
	"fmt"
	"os"
	"strconv"
)

type service struct {
	Port uint16

	Name string
}

func (s *service) loadFromEnv(defaultValue service) {
	osPort := os.Getenv("SERVICE_PORT")
	if osPort == "" {
		osPort = fmt.Sprintf("%d", defaultValue.Port)
	}

	port, err := strconv.Atoi(osPort)
	if err != nil {
		port = int(defaultValue.Port)
	}
	s.Port = uint16(port)

	osName := os.Getenv("SERVICE_NAME")
	if osName == "" {
		osName = defaultValue.Name
	}
	s.Name = osName
}
