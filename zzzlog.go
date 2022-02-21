// Package zzzlog provides a minimalistic level logging library based on
// the zzzlogi level logging interface.
package zzzlog

import (
	"io"

	"github.com/tuxdude/zzzlogi"
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
}

// NewLogger instantiates a Logger.
func NewLogger(userConfig *Config) zzzlogi.Logger {
	c := defaultLoggingConfig()
	c.dest = userConfig.Dest
	c.maxLevel = userConfig.MaxLevel
	return newLoggerForConfig(c)
}
