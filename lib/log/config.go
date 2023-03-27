package log

import "fmt"

// Config of log
type LogConfig struct {
	LogLevel      string
	Path          string
	ErrorFileName string
	InfoFileName  string
}

const (
	// Default Value
	DefaultLogLevel      = "debug"
	DefaultFilePath      = "log"
	DefaultErrorFileName = "error"
	DefaultInfoFileName  = "info"
)

// setDefaultConfig for set default value log config
func (lc *LogConfig) setDefaultConfig() {
	if lc.LogLevel == "" {
		lc.LogLevel = DefaultLogLevel
	}

	if lc.Path == "" {
		lc.Path = DefaultFilePath
	}

	if lc.ErrorFileName == "" {
		lc.ErrorFileName = DefaultErrorFileName
	}

	if lc.InfoFileName == "" {
		lc.InfoFileName = DefaultInfoFileName
	}
}

// getLogErrorPath for get log error file path
func (lc *LogConfig) getLogErrorFilePath() string {
	return fmt.Sprintf("./%s/%s.log", lc.Path, lc.ErrorFileName)
}

// getLogInfoFilePath for get log info file path
func (lc *LogConfig) getLogInfoFilePath() string {
	return fmt.Sprintf("./%s/%s.log", lc.Path, lc.InfoFileName)
}
