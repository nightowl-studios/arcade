package log

// Copy pasted from
// https://www.mountedthoughts.com/golang-logger-interface/

import "errors"

// A global variable so that log functions can be directly accessed
var log Logger

//Fields Type to pass when we want to call WithFields for structured logging
type Fields map[string]interface{}

const (
	//Debug has verbose message
	Debug = "debug"
	//Info is default log level
	Info = "info"
	//Warn is for logging messages about possible issues
	Warn = "warn"
	//Error is for logging errors
	Error = "error"
	//Fatal is for logging fatal messages. The sytem shutsdown after logging the message.
	Fatal = "fatal"
)

const (
	InstanceZapLogger int = iota
	InstanceLogrusLogger
)

var (
	errInvalidLoggerInstance = errors.New("Invalid logger instance")
	DefaultTimeFormat        = "2006-01-02 15:04:05"
)

//Logger is our contract for the logger
type Logger interface {
	Debugf(format string, args ...interface{})

	Infof(format string, args ...interface{})

	Warnf(format string, args ...interface{})

	Errorf(format string, args ...interface{})

	Fatalf(format string, args ...interface{})

	Panicf(format string, args ...interface{})

	WithFields(keyValues Fields) Logger
}

// Configuration stores the config for the logger
// For some loggers there can only be one level across writers, for such the level of Console is picked by default
type Configuration struct {
	EnableConsole     bool
	ConsoleJSONFormat bool
	ConsoleLevel      string
	EnableFile        bool
	FileJSONFormat    bool
	FileLevel         string
	FileLocation      string
}

//NewLogger returns an instance of logger
func NewLogger(config Configuration, loggerInstance int) error {
	switch loggerInstance {
	case InstanceZapLogger:
		logger, err := newZapLogger(config)
		if err != nil {
			return err
		}
		log = logger
		return nil

	case InstanceLogrusLogger:
		logger, err := newLogrusLogger(config)
		if err != nil {
			return err
		}
		log = logger
		return nil

	default:
		return errInvalidLoggerInstance
	}
}

func Debugf(format string, args ...interface{}) {
	initIfEmpty()
	log.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	initIfEmpty()
	log.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	initIfEmpty()
	log.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	initIfEmpty()
	log.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	initIfEmpty()
	log.Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	initIfEmpty()
	log.Panicf(format, args...)
}

func WithFields(keyValues Fields) Logger {
	initIfEmpty()
	return log.WithFields(keyValues)
}

func initIfEmpty() {
	if log == nil {
		NewLogger(
			Configuration{
				EnableConsole:     true,
				ConsoleJSONFormat: false,
				ConsoleLevel:      "info",
				EnableFile:        false,
				FileJSONFormat:    false,
			},
			InstanceLogrusLogger,
		)
	}
}
