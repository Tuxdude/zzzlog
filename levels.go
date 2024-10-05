package zzzlog

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
	stringerLevelname = map[Level]string{
		LvlFatal: "Fatal",
		LvlError: "Error",
		LvlWarn:  "Warn",
		LvlInfo:  "Info",
		LvlDebug: "Debug",
		LvlTrace: "Trace",
	}
)
