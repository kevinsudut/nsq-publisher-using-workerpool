package log

import (
	"time"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

type Log struct {
	logger *logrus.Logger
}

// InitLog for init log object
func InitLog(cfg LogConfig) (Logger, error) {
	cfg.setDefaultConfig()

	return initLogger(cfg)
}

// initLoger for init log object
func initLogger(cfg LogConfig) (Logger, error) {
	level, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		return nil, err
	}

	logger := logrus.New()
	logger.SetLevel(level)
	logger.Hooks.Add(lfshook.NewHook(
		getPathMap(
			cfg.getLogErrorFilePath(),
			cfg.getLogInfoFilePath(),
		),
		getTextFormatter(),
	))

	log := &Log{
		logger: logger,
	}

	return log, nil
}

// getPathMap for mapping log level to path file
func getPathMap(errorFilePath string, infoFilePath string) lfshook.PathMap {
	return lfshook.PathMap{
		logrus.PanicLevel: errorFilePath,
		logrus.ErrorLevel: errorFilePath,
		logrus.FatalLevel: errorFilePath,
		logrus.WarnLevel:  errorFilePath,
		logrus.InfoLevel:  infoFilePath,
		logrus.DebugLevel: infoFilePath,
	}
}

// getTextFormatter for get logrus text formatter
func getTextFormatter() *logrus.TextFormatter {
	return &logrus.TextFormatter{
		TimestampFormat: time.RFC3339,
		FullTimestamp:   true,
	}
}
