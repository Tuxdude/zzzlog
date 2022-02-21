package zzzlog

// Logging levels.
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

var (
	orderedLevels = []Level{
		LvlFatal,
		LvlError,
		LvlWarn,
		LvlInfo,
		LvlDebug,
		LvlTrace,
	}
	levelName = map[Level]string{
		LvlFatal: "FATAL",
		LvlError: "ERROR",
		LvlWarn:  "WARN ",
		LvlInfo:  "INFO ",
		LvlDebug: "DEBUG",
		LvlTrace: "TRACE",
	}
)

// Level represents the logging level.
type Level uint8
