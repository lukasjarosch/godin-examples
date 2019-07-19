package log

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-godin/log/level"
	"github.com/go-kit/kit/log"
)

// Logger is the default logging interface which is used throughout all godin services.
type Logger interface {
	Log(keyvals ...interface{})
	Debug(message string, keyvals ...interface{})
	Info(message string, keyvals ...interface{})
	Warning(message string, keyvals ...interface{})
	Error(message string, keyvals ...interface{})
	With(keyvals ...interface{}) Log
}

const (
	LevelDebug          = "debug"
	LevelInfo           = "info"
	LevelWarning        = "warning"
	LevelError          = "error"
	MessageKey          = "message"
	EnvironmentVariable = "LOG_LEVEL"
)

type Log struct {
	kitLogger log.Logger
}

// NewLogger creates a new, leveled Log. The given level is the allowed minimal level.
func NewLogger(logLevel string) Log {
	levelOpt, err := evaluateLogLevel(logLevel)

	var kitLogger log.Logger
	kitLogger = log.NewJSONLogger(log.NewSyncWriter(os.Stdout))
	kitLogger = level.NewFilter(kitLogger, levelOpt)

	log := Log{
		kitLogger: kitLogger,
	}

	// the error from evaluateLogLevel needs to be logged
	if err != nil {
		log.Warning("", "err", err)
	}

	return log
}

// NewLoggerFromEnv creates a new Log, configuring the log level using an environment variable.
func NewLoggerFromEnv() Log {
	levelStr := os.Getenv(EnvironmentVariable)
	return NewLogger(levelStr)
}

// Log redirects to go-kit/log.Log
func (l Log) Log(keyvals ...interface{}) {
	_ = l.kitLogger.Log(keyvals...)
}

// Debug will log a message and arbitrary key-value pairs
func (l Log) Debug(message string, keyvals ...interface{}) {
	_ = level.Debug(l.kitLogger).Log(l.mergeKeyValues(message, keyvals)...)
}

// Info will log a message and arbitrary key-value pairs
func (l Log) Info(message string, keyvals ...interface{}) {
	_ = level.Info(l.kitLogger).Log(l.mergeKeyValues(message, keyvals)...)
}

// Warning will log a message and arbitrary key-value pairs
func (l Log) Warning(message string, keyvals ...interface{}) {
	_ = level.Warn(l.kitLogger).Log(l.mergeKeyValues(message, keyvals)...)
}

// Error will log a message and arbitrary key-value pairs
func (l Log) Error(message string, keyvals ...interface{}) {
	_ = level.Error(l.kitLogger).Log(l.mergeKeyValues(message, keyvals)...)
}

func (l Log) With(keyvals ...interface{}) Log {
	if len(keyvals) == 0 {
		return l
	}

	kitLogger := log.With(l.kitLogger, keyvals...)

	return Log{
		kitLogger: kitLogger,
	}
}

// evaluateLogLevel maps a given logLevel as string (e.g. from an ENV variable) to a level Option.
// If the passed logLevel does not exist, all levels will be enabled by default.
func evaluateLogLevel(logLevel string) (level.Option, error) {
	logLevel = strings.ToLower(logLevel)
	switch logLevel {
	case LevelDebug:
		return level.AllowDebug(), nil
	case LevelInfo:
		return level.AllowInfo(), nil
	case LevelWarning:
		return level.AllowWarn(), nil
	case LevelError:
		return level.AllowError(), nil
	default:
		return level.AllowAll(), fmt.Errorf("no log-level passed, falling back to debug")
	}
}

// mergeKeyValues will append the level and message field to already existing keyvals
func (l Log) mergeKeyValues(message string, keyvals []interface{}) []interface{} {
	var list []interface{}
	var levelData []interface{}

	if message != "" {
		levelData = append(levelData, MessageKey)
		levelData = append(levelData, message)
	}

	list = append(list, levelData...)
	list = append(list, keyvals...)

	return list
}
