package zzzlog

import "os"

var (
	defaultLevelColors = levelColorMap{
		lvlFatal: {
			color: colorCodeRed,
			bold: true,
		},
		lvlError: {
			color: colorCodeRed,
		},
		lvlWarn: {
			color: colorCodeYellow,
		},
		lvlInfo: {
			color: colorCodeBlue,
		},
		lvlDebug: {
			color: colorCodeGreen,
		},
		lvlTrace: {
			color: colorCodeMagenta,
		},
	}
)

func defaultLoggingConfig() *loggerConfig {
	return &loggerConfig{
		dest:        os.Stdout,
		levelColors: defaultLevelColors,
	}
}
