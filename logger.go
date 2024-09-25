package zzzlog

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/tuxdude/zzzlogi"
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
	// skipTimestamp set to true skips logging the timestamp in the logs.
	skipTimestamp bool
	// skipLogLevel seto true skips logging the log level in the logs.
	skipLogLevel bool
	// skipCallerInfo set to true skips logging the call site information.
	skipCallerInfo bool
	// panicInFatal set to true causes the log message to be emitted
	// through panic() after logging, instead of the default behavior of
	// exiting with a status code 1 when using Fatal or FatalF logging methods.
	panicInFatal bool
	// timestampLoggingFormat determines the format for logging the timestamps.
	timestampLoggingFormat string
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

	if l.config.panicInFatal {
		panic(fmt.Sprintf(defaultFormat(len(args)), args...))
	} else {
		os.Exit(1)
	}
}

func (l *loggerImpl) Fatalf(format string, args ...interface{}) {
	l.log(LvlFatal, 1, format, args...)
	l.write("\n%s\n", stackTraces())

	if l.config.panicInFatal {
		panic(fmt.Sprintf(format, args...))
	} else {
		os.Exit(1)
	}
}

func (l *loggerImpl) Error(args ...interface{}) {
	l.log(LvlError, 1, defaultFormat(len(args)), args...)
}

func (l *loggerImpl) Errorf(format string, args ...interface{}) {
	l.log(LvlError, 1, format, args...)
}

func (l *loggerImpl) ErrorEmpty() {
	l.logEmpty(LvlError)
}

func (l *loggerImpl) Warn(args ...interface{}) {
	l.log(LvlWarn, 1, defaultFormat(len(args)), args...)
}

func (l *loggerImpl) Warnf(format string, args ...interface{}) {
	l.log(LvlWarn, 1, format, args...)
}

func (l *loggerImpl) WarnEmpty() {
	l.logEmpty(LvlWarn)
}

func (l *loggerImpl) Info(args ...interface{}) {
	l.log(LvlInfo, 1, defaultFormat(len(args)), args...)
}

func (l *loggerImpl) Infof(format string, args ...interface{}) {
	l.log(LvlInfo, 1, format, args...)
}

func (l *loggerImpl) InfoEmpty() {
	l.logEmpty(LvlInfo)
}

func (l *loggerImpl) Debug(args ...interface{}) {
	l.log(LvlDebug, 1, defaultFormat(len(args)), args...)
}

func (l *loggerImpl) Debugf(format string, args ...interface{}) {
	l.log(LvlDebug, 1, format, args...)
}

func (l *loggerImpl) DebugEmpty() {
	l.logEmpty(LvlDebug)
}

func (l *loggerImpl) Trace(args ...interface{}) {
	l.log(LvlTrace, 1, defaultFormat(len(args)), args...)
}

func (l *loggerImpl) Tracef(format string, args ...interface{}) {
	l.log(LvlTrace, 1, format, args...)
}

func (l *loggerImpl) TraceEmpty() {
	l.logEmpty(LvlTrace)
}

func (l *loggerImpl) Print(args ...interface{}) {
	l.write(fmt.Sprintf("%s\n", defaultFormat(len(args))), args...)
}

func (l *loggerImpl) Printf(format string, args ...interface{}) {
	l.write(fmt.Sprintf("%s\n", format), args...)
}

func (l *loggerImpl) log(lvl Level, skipFrames int, format string, args ...interface{}) {
	if lvl > l.config.maxLevel {
		return
	}

	var f strings.Builder
	var a []interface{}

	if !l.config.skipTimestamp {
		f.WriteString("%s  ")
		a = append(a, time.Now().Format(l.config.timestampLoggingFormat))
	}
	if !l.config.skipLogLevel {
		f.WriteString("%s  ")
		a = append(a, l.levelStr[lvl])
	}
	if !l.config.skipCallerInfo {
		f.WriteString("%-40s  ")
		a = append(a, callerInfo(skipFrames+1))
	}
	f.WriteString(format)
	f.WriteString("\n")
	a = append(a, args...)
	l.write(f.String(), a...)
}

func (l *loggerImpl) logEmpty(lvl Level) {
	if lvl > l.config.maxLevel {
		return
	}
	l.write("\n")
}

func (l *loggerImpl) write(format string, args ...interface{}) {
	fmt.Fprintf(l.config.dest, format, args...)
}
