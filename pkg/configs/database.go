package configs

import (
	"fmt"
	"os"
	"strconv"
)

type database struct {
	Driver   string
	Host     string
	Port     uint16
	User     string
	Password string
	Database string
	Schema   string
}

func (d *database) loadFromEnv(defaultValue database) {
	osDriver := os.Getenv("DB_DRIVER")
	if osDriver == "" {
		osDriver = defaultValue.Driver
	}
	d.Driver = osDriver

	osHost := os.Getenv("DB_HOST")
	if osHost == "" {
		osHost = defaultValue.Host
	}
	d.Host = osHost

	osPort := os.Getenv("DB_PORT")
	if osPort == "" {
		osPort = fmt.Sprintf("%d", defaultValue.Port)
	}

	port, err := strconv.Atoi(osPort)
	if err != nil {
		port = int(defaultValue.Port)
	}
	d.Port = uint16(port)

	osUser := os.Getenv("DB_USER")
	if osUser == "" {
		osUser = defaultValue.User
	}
	d.User = osUser

	osPassword := os.Getenv("DB_PASSWORD")
	if osPassword == "" {
		osPassword = defaultValue.Password
	}
	d.Password = osPassword

	osDatabase := os.Getenv("DB_DATABASE")
	if osDatabase == "" {
		osDatabase = defaultValue.Database
	}
	d.Database = osDatabase

	osSchema := os.Getenv("DB_SCHEMA")
	if osSchema == "" {
		osSchema = defaultValue.Schema
	}
	d.Schema = osSchema
}
