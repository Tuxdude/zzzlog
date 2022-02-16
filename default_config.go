package zzzlog

var (
	defaultLevelColors = levelColorMap{
		lvlFatal: {
			code: colorCodeRed,
			bold: true,
		},
		lvlError: {
			code: colorCodeRed,
			bold: false,
		},
		lvlWarn: {
			code: colorCodeYellow,
			bold: false,
		},
		lvlInfo: {
			code: colorCodeBlue,
			bold: false,
		},
		lvlDebug: {
			code: colorCodeGreen,
			bold: false,
		},
		lvlTrace: {
			code: colorCodeMagenta,
			bold: false,
		},
	}
)
