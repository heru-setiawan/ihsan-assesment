package configs

import (
	"os"
	"strconv"

	"github.com/fatih/structs"
)

type Logrus struct {
	Level           int    `default:"4"`
	FormatTimestamp string `default:"2006-01-02 15:04:05"`
	FullTimestamp   bool   `default:"true"`
	ForceColors     bool   `default:"true"`
}

func (cfg *Logrus) defaultValue(key string) string {
	s := structs.New(cfg)
	field := s.Field(key)
	if field == nil {
		return ""
	}
	return field.Tag("default")
}

func (cfg *Logrus) loadEnv() error {
	osLevel := os.Getenv("LOG_LOGRUS_LEVEL")
	if osLevel == "" {
		osLevel = cfg.defaultValue("Level")
	}

	level, err := strconv.Atoi(osLevel)
	if err != nil {
		return err
	}
	cfg.Level = level

	osFormatTimestamp := os.Getenv("LOG_LOGRUS_TIMESTAMP_FORMAT")
	if osFormatTimestamp == "" {
		osFormatTimestamp = cfg.defaultValue("FormatTimestamp")
	}
	cfg.FormatTimestamp = osFormatTimestamp

	osFullTimestamp := os.Getenv("LOG_LOGRUS_TIMESTAMP_FULL")
	if osFullTimestamp == "" {
		osFullTimestamp = cfg.defaultValue("FullTimestamp")
	}

	fullTimestamp, err := strconv.ParseBool(osFullTimestamp)
	if err != nil {
		return err
	}
	cfg.FullTimestamp = fullTimestamp

	osForceColors := os.Getenv("LOG_LOGRUS_FORCE_COLORS")
	if osForceColors == "" {
		osForceColors = cfg.defaultValue("ForceColors")
	}

	ForceColors, err := strconv.ParseBool(osForceColors)
	if err != nil {
		return err
	}
	cfg.ForceColors = ForceColors

	return nil
}
