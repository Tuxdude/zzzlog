package zzzlog

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/tuxdude/zzzlogi"
)

const (
	timestampFormat = "2006-01-02T15:04:05.000Z0700"
)

var (
	colorLevelStr = []string{
		"\x1b[31mFATAL\x1b[0m",
		"\x1b[31mERROR\x1b[0m",
		"\x1b[33mWARN \x1b[0m",
		"\x1b[34mINFO \x1b[0m",
		"\x1b[32mDEBUG\x1b[0m",
		"\x1b[35mTRACE\x1b[0m",
	}
	defaultFormat = []string{
		"",
		"%v",
		"%v %v",
		"%v %v %v",
		"%v %v %v %v",
		"%v %v %v %v %v",
		"%v %v %v %v %v %v",
		"%v %v %v %v %v %v %v",
		"%v %v %v %v %v %v %v %v",
		"%v %v %v %v %v %v %v %v %v",
		"%v %v %v %v %v %v %v %v %v %v",
		"%v %v %v %v %v %v %v %v %v %v %v",
		"%v %v %v %v %v %v %v %v %v %v %v %v",
		"%v %v %v %v %v %v %v %v %v %v %v %v %v",
		"%v %v %v %v %v %v %v %v %v %v %v %v %v %v",
		"%v %v %v %v %v %v %v %v %v %v %v %v %v %v %v",
	}
)

// Logging levels.
const (
	lvlFatal level = iota
	lvlError
	lvlWarn
	lvlInfo
	lvlDebug
	lvlTrace
)

// Level represents the logging level.
type level uint8

type loggerImpl struct {
	levelStr []string
}

// NewLogger instantiates a Logger.
func NewLogger() zzzlogi.Logger {
	logger := &loggerImpl{}
	logger.levelStr = colorLevelStr
	return logger
}

func (l *loggerImpl) Fatal(args ...interface{}) {
	l.log(lvlFatal, format(len(args)), args...)
}

func (l *loggerImpl) Fatalf(format string, args ...interface{}) {
	l.log(lvlFatal, format, args...)
}

func (l *loggerImpl) Error(args ...interface{}) {
	l.log(lvlError, format(len(args)), args...)
}

func (l *loggerImpl) Errorf(format string, args ...interface{}) {
	l.log(lvlError, format, args...)
}

func (l *loggerImpl) Warn(args ...interface{}) {
	l.log(lvlWarn, format(len(args)), args...)
}

func (l *loggerImpl) Warnf(format string, args ...interface{}) {
	l.log(lvlWarn, format, args...)
}

func (l *loggerImpl) Info(args ...interface{}) {
	l.log(lvlInfo, format(len(args)), args...)
}

func (l *loggerImpl) Infof(format string, args ...interface{}) {
	l.log(lvlInfo, format, args...)
}

func (l *loggerImpl) Debug(args ...interface{}) {
	l.log(lvlDebug, format(len(args)), args...)
}

func (l *loggerImpl) Debugf(format string, args ...interface{}) {
	l.log(lvlDebug, format, args...)
}

func (l *loggerImpl) Trace(args ...interface{}) {
	l.log(lvlTrace, format(len(args)), args...)
}

func (l *loggerImpl) Tracef(format string, args ...interface{}) {
	l.log(lvlTrace, format, args...)
}

func (l *loggerImpl) log(lvl level, format string, args ...interface{}) {
	file, line, _ := callerInfo()
	fmt.Printf("%s  %s  %s:%d ", time.Now().Format(timestampFormat), l.levelStr[lvl], file, line)
	fmt.Printf(format, args...)
	fmt.Printf("\n")
}

func format(count int) string {
	if count < len(defaultFormat) {
		return defaultFormat[count]
	}
	return buildDefaultFormat(count)
}

func buildDefaultFormat(count int) string {
	if count == 0 {
		return ""
	}
	var result strings.Builder
	result.WriteString("%v")
	count--
	for count > 0 {
		count--
		result.WriteString(" %v")
	}
	return result.String()
}

func callerInfo() (string, int, string) {
	// TODO: Make the skip more generic.
	frame, ok := callerFrame(5)
	if !ok {
		return "UnknownFile", 0, "UnknownFunction"
	}

	ix := strings.LastIndexByte(frame.File, '/')
	if ix == -1 {
		return frame.File, frame.Line, frame.Function
	}
	ix = strings.LastIndexByte(frame.File[:ix], '/')
	if ix == -1 {
		return frame.File, frame.Line, frame.Function
	}
	return frame.File[ix+1:], frame.Line, frame.Function
}

func callerFrame(skip int) (runtime.Frame, bool) {
	var frame runtime.Frame
	pc := make([]uintptr, 1)
	numFrames := runtime.Callers(skip, pc)
	if numFrames < 1 {
		return frame, false
	}

	frame, _ = runtime.CallersFrames(pc).Next()
	return frame, frame.PC != 0
}
