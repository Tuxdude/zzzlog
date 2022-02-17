// Package zzzlog provides a minimalistic level logging library based on
// the zzzlogi level logging interface.
package zzzlog

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/tuxdude/zzzlogi"
)

const (
	timestampFormat = "2006-01-02T15:04:05.000Z0700"
)

type loggerImpl struct {
	writer   io.Writer
	levelStr []string
}

type loggerConfig struct {
	dest        io.Writer
	levelColors levelColorMap
}

// NewLogger instantiates a Logger.
func NewLogger() zzzlogi.Logger {
	return newLoggerForConfig(defaultLoggingConfig())
}

func newLoggerForConfig(config *loggerConfig) zzzlogi.Logger {
	logger := &loggerImpl{
		writer:   config.dest,
		levelStr: buildColoredLevels(config.levelColors),
	}
	return logger
}

func (l *loggerImpl) Fatal(args ...interface{}) {
	l.log(lvlFatal, 1, defaultFormat(len(args)), args...)
	l.write("\n%s\n", stackTraces())
	os.Exit(1)
}

func (l *loggerImpl) Fatalf(format string, args ...interface{}) {
	l.log(lvlFatal, 1, format, args...)
	l.write("\n%s\n", stackTraces())
	os.Exit(1)
}

func (l *loggerImpl) Error(args ...interface{}) {
	l.log(lvlError, 1, defaultFormat(len(args)), args...)
}

func (l *loggerImpl) Errorf(format string, args ...interface{}) {
	l.log(lvlError, 1, format, args...)
}

func (l *loggerImpl) Warn(args ...interface{}) {
	l.log(lvlWarn, 1, defaultFormat(len(args)), args...)
}

func (l *loggerImpl) Warnf(format string, args ...interface{}) {
	l.log(lvlWarn, 1, format, args...)
}

func (l *loggerImpl) Info(args ...interface{}) {
	l.log(lvlInfo, 1, defaultFormat(len(args)), args...)
}

func (l *loggerImpl) Infof(format string, args ...interface{}) {
	l.log(lvlInfo, 1, format, args...)
}

func (l *loggerImpl) Debug(args ...interface{}) {
	l.log(lvlDebug, 1, defaultFormat(len(args)), args...)
}

func (l *loggerImpl) Debugf(format string, args ...interface{}) {
	l.log(lvlDebug, 1, format, args...)
}

func (l *loggerImpl) Trace(args ...interface{}) {
	l.log(lvlTrace, 1, defaultFormat(len(args)), args...)
}

func (l *loggerImpl) Tracef(format string, args ...interface{}) {
	l.log(lvlTrace, 1, format, args...)
}

func (l *loggerImpl) log(lvl level, skipFrames int, format string, args ...interface{}) {
	f := "%s  %s  %-30s  " + format + "\n"
	a := []interface{}{
		time.Now().Format(timestampFormat),
		l.levelStr[lvl],
		callerInfo(skipFrames + 1),
	}
	a = append(a, args...)
	l.write(f, a...)
}

func (l *loggerImpl) write(format string, args ...interface{}) {
	fmt.Fprintf(l.writer, format, args...)
}
