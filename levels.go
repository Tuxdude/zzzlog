package zzzlog

// Logging levels.
const (
	lvlFatal level = iota
	lvlError
	lvlWarn
	lvlInfo
	lvlDebug
	lvlTrace
)

var (
	orderedLevels = []level{
		lvlFatal,
		lvlError,
		lvlWarn,
		lvlInfo,
		lvlDebug,
		lvlTrace,
	}
	levelName = map[level]string{
		lvlFatal: "FATAL",
		lvlError: "ERROR",
		lvlWarn:  "WARN ",
		lvlInfo:  "INFO ",
		lvlDebug: "DEBUG",
		lvlTrace: "TRACE",
	}
)

// level represents the logging level.
type level uint8
