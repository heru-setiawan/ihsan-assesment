package logs

import (
	"assesment/pkg/configs"
	"runtime"

	"github.com/sirupsen/logrus"
)

func getCaller() (file, function string, line int) {
	pc, file, line, _ := runtime.Caller(3)
	function = runtime.FuncForPC(pc).Name()
	return
}

func NewLogger(config configs.Logrus, serviceName string) (logger *Logger) {
	log := logrus.New()

	formatter := new(logrus.TextFormatter)
	formatter.TimestampFormat = config.FormatTimestamp
	formatter.FullTimestamp = config.FullTimestamp
	formatter.ForceColors = config.ForceColors

	log.SetLevel(logrus.Level(config.Level))
	log.SetFormatter(formatter)

	return &Logger{
		log: log.WithField("service", serviceName),
	}
}

type Logger struct {
	log *logrus.Entry
}

func (l *Logger) BaseLog(data map[string]any) (log *logrus.Entry) {
	file, function, line := getCaller()
	log = l.log.WithFields(logrus.Fields{
		"file":     file,
		"line":     line,
		"function": function,
	}).WithFields(data)
	return
}

func (l *Logger) Info(data map[string]any, message string) {
	l.BaseLog(data).Info(message)
}

func (l *Logger) Warn(data map[string]any, message string) {
	l.BaseLog(data).Warn(message)
}

func (l *Logger) Error(data map[string]any, message string) {
	l.BaseLog(data).Error(message)
}

func (l *Logger) Fatal(data map[string]any, message string) {
	l.BaseLog(data).Fatal(message)
}

func (l *Logger) Panic(data map[string]any, message string) {
	l.BaseLog(data).Panic(message)
}
