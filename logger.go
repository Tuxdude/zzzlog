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

// loggerImpl is the implementation of the level logger based on
// zzzlogi.Logger interface.
type loggerImpl struct {
	// config contains the logger configuration.
	config *configInternal
	// levelStr contains the colorized (if configured) log level strings.
	levelStr []string
}

// configInternal is the internal logger configuration used that is not
// exported to the callers of this library.
type configInternal struct {
	// dest is the logging destination for the logs.
	dest io.Writer
	// maxLevel determines the maximum logging level.
	maxLevel Level
	// levelColors contains the color configuration for each log level.
	levelColors levelColorMap
}

// newLoggerForConfig builds a logger based on the specified config.
func newLoggerForConfig(config *configInternal) zzzlogi.Logger {
	logger := &loggerImpl{
		config:   config,
		levelStr: buildColoredLevels(config.levelColors),
	}
	return logger
}

func (l *loggerImpl) Fatal(args ...interface{}) {
	l.log(LvlFatal, 1, defaultFormat(len(args)), args...)
	l.write("\n%s\n", stackTraces())
	os.Exit(1)
}

func (l *loggerImpl) Fatalf(format string, args ...interface{}) {
	l.log(LvlFatal, 1, format, args...)
	l.write("\n%s\n", stackTraces())
	os.Exit(1)
}

func (l *loggerImpl) Error(args ...interface{}) {
	l.log(LvlError, 1, defaultFormat(len(args)), args...)
}

func (l *loggerImpl) Errorf(format string, args ...interface{}) {
	l.log(LvlError, 1, format, args...)
}

func (l *loggerImpl) Warn(args ...interface{}) {
	l.log(LvlWarn, 1, defaultFormat(len(args)), args...)
}

func (l *loggerImpl) Warnf(format string, args ...interface{}) {
	l.log(LvlWarn, 1, format, args...)
}

func (l *loggerImpl) Info(args ...interface{}) {
	l.log(LvlInfo, 1, defaultFormat(len(args)), args...)
}

func (l *loggerImpl) Infof(format string, args ...interface{}) {
	l.log(LvlInfo, 1, format, args...)
}

func (l *loggerImpl) Debug(args ...interface{}) {
	l.log(LvlDebug, 1, defaultFormat(len(args)), args...)
}

func (l *loggerImpl) Debugf(format string, args ...interface{}) {
	l.log(LvlDebug, 1, format, args...)
}

func (l *loggerImpl) Trace(args ...interface{}) {
	l.log(LvlTrace, 1, defaultFormat(len(args)), args...)
}

func (l *loggerImpl) Tracef(format string, args ...interface{}) {
	l.log(LvlTrace, 1, format, args...)
}

func (l *loggerImpl) log(lvl Level, skipFrames int, format string, args ...interface{}) {
	if lvl > l.config.maxLevel {
		return
	}

	f := "%s  %s  %-40s  " + format + "\n"
	a := []interface{}{
		time.Now().Format(timestampFormat),
		l.levelStr[lvl],
		callerInfo(skipFrames + 1),
	}
	a = append(a, args...)
	l.write(f, a...)
}

func (l *loggerImpl) write(format string, args ...interface{}) {
	fmt.Fprintf(l.config.dest, format, args...)
}
