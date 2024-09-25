// Package zzzlog provides a minimalistic level logging library based on
// the zzzlogi level logging interface.
package zzzlog

import (
	"io"
	"os"

	"github.com/tuxdude/zzzlogi"
)

const (
	defaultTimestampFormat = "2006-01-02T15:04:05.000Z0700"
)

// Logging levels used by the zzzlog logger.
const (
	// LvlFatal represents the Fatal log level.
	LvlFatal Level = iota
	// LvlError represents the Error log level.
	LvlError
	// LvlWarn represents the Warn log level.
	LvlWarn
	// LvlInfo represents the Info log level.
	LvlInfo
	// LvlDebug represents the Debug log level.
	LvlDebug
	// LvlTrace represents the Trace log level.
	LvlTrace
)

// Level represents the logging level used by the zzzlog logger.
type Level uint8

// Config contains the configuration for the logger.
type Config struct {
	// Dest is the logging destination for the logs.
	Dest io.Writer
	// Level determines the maximum logging level.
	MaxLevel Level
	// SkipTimestamp set to true skips logging the timestamp in the logs.
	SkipTimestamp bool
	// SkipLogLevel seto to true skips logging the log level in the logs.
	SkipLogLevel bool
	// SkipCallerInfo set to true skips logging the call site information.
	SkipCallerInfo bool
	// PanicInFatal set to true causes the log message to be emitted
	// through panic() after logging, instead of the default behavior of
	// exiting with a status code 1 when using Fatal or FatalF logging methods.
	PanicInFatal bool
	// TimestampLoggingFormat determines the format for logging the timestamps.
	TimestampLoggingFormat string
}

// NewLogger instantiates a Logger.
func NewLogger(userConfig *Config) zzzlogi.Logger {
	c := defaultLoggingConfig()
	c.dest = userConfig.Dest
	c.maxLevel = userConfig.MaxLevel
	c.skipTimestamp = userConfig.SkipTimestamp
	c.skipLogLevel = userConfig.SkipLogLevel
	c.skipCallerInfo = userConfig.SkipCallerInfo
	c.panicInFatal = userConfig.PanicInFatal
	if userConfig.TimestampLoggingFormat != "" {
		c.timestampLoggingFormat = userConfig.TimestampLoggingFormat
	} else {
		c.timestampLoggingFormat = defaultTimestampFormat
	}
	return newLoggerForConfig(c)
}

// NewConsoleLoggerConfig returns a logger configuration for
// logging to stdout with the maximum logging level set to Info.
func NewConsoleLoggerConfig() *Config {
	return &Config{
		Dest:     os.Stdout,
		MaxLevel: LvlInfo,
	}
}

func NewVanillaLoggerConfig() *Config {
	return &Config{
		Dest:           os.Stdout,
		MaxLevel:       LvlInfo,
		SkipTimestamp:  true,
		SkipLogLevel:   true,
		SkipCallerInfo: true,
	}
}
