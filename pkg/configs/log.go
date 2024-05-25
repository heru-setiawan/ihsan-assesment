package configs

import (
	"fmt"
	"os"
	"strconv"
)

type log struct {
	Level           int
	FormatTimestamp string
	FullTimestamp   bool
	ForceColors     bool
}

func (l *log) loadFromEnv(defaultValue log) {
	osLevel := os.Getenv("LOG_LEVEL")
	if osLevel == "" {
		osLevel = fmt.Sprintf("%d", defaultValue.Level)
	}

	level, err := strconv.Atoi(osLevel)
	if err != nil {
		level = defaultValue.Level
	}
	l.Level = level

	osFormatTimestamp := os.Getenv("LOG_TIMESTAMP_FORMAT")
	if osFormatTimestamp == "" {
		osFormatTimestamp = defaultValue.FormatTimestamp
	}
	l.FormatTimestamp = osFormatTimestamp

	osFullTimestamp := os.Getenv("LOG_TIMESTAMP_FULL")
	if osFullTimestamp == "" {
		osFullTimestamp = strconv.FormatBool(defaultValue.FullTimestamp)
	}

	fullTimestamp, err := strconv.ParseBool(osFullTimestamp)
	if err != nil {
		fullTimestamp = defaultValue.FullTimestamp
	}
	l.FullTimestamp = fullTimestamp

	osForceColors := os.Getenv("LOG_FORCE_COLORS")
	if osForceColors == "" {
		osForceColors = strconv.FormatBool(defaultValue.ForceColors)
	}

	ForceColors, err := strconv.ParseBool(osForceColors)
	if err != nil {
		ForceColors = defaultValue.ForceColors
	}
	l.ForceColors = ForceColors
}
