package zzzlog

import "os"

var (
	colorizedDefault = levelColorMap{
		LvlFatal: {
			color: colorCodeRed,
			bold:  true,
		},
		LvlError: {
			color: colorCodeRed,
		},
		LvlWarn: {
			color: colorCodeYellow,
		},
		LvlInfo: {
			color: colorCodeBlue,
		},
		LvlDebug: {
			color: colorCodeGreen,
		},
		LvlTrace: {
			color: colorCodeMagenta,
		},
	}
	nonColorizedDefault = levelColorMap{
		LvlFatal: {
			color: colorCodeNone,
		},
		LvlError: {
			color: colorCodeNone,
		},
		LvlWarn: {
			color: colorCodeNone,
		},
		LvlInfo: {
			color: colorCodeNone,
		},
		LvlDebug: {
			color: colorCodeNone,
		},
		LvlTrace: {
			color: colorCodeNone,
		},
	}
)

func defaultLoggingConfig() *configInternal {
	c := &configInternal{
		dest:           os.Stdout,
		maxLevel:       LvlInfo,
		skipCallerInfo: false,
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
