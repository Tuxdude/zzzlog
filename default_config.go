package zzzlog

import "os"

var (
	colorizedDefault = levelColorMap{
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
	nonColorizedDefault = levelColorMap{
		lvlFatal: {
			color: colorCodeNone,
		},
		lvlError: {
			color: colorCodeNone,
		},
		lvlWarn: {
			color: colorCodeNone,
		},
		lvlInfo: {
			color: colorCodeNone,
		},
		lvlDebug: {
			color: colorCodeNone,
		},
		lvlTrace: {
			color: colorCodeNone,
		},
	}
)

func defaultLoggingConfig() *loggerConfig {
	c := &loggerConfig{
		dest: os.Stdout,
	}
	if isTTY() {
		c.levelColors = colorizedDefault
	} else {
		c.levelColors = nonColorizedDefault
	}
	return c
}

func isTTY() bool {
	f, err := os.Stdin.Stat()
	if err != nil {
		// Unable to get FileInfo for stdin, will fall back to assuming
		// it is not a TTY.
		return false
	}
	if (f.Mode() & os.ModeCharDevice) == 0 {
		return false
	}
	return true
}
