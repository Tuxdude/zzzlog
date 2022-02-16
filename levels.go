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

// level represents the logging level.
type level uint8
