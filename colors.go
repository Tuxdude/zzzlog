package zzzlog

import "fmt"

const (
	colorPrefix = "\x1b["
	colorBold   = "1;"
	colorReset  = "\x1b[0m"
)

// nolint (deadcode)
const (
	colorCodeBlack colorCode = iota + 30
	colorCodeRed
	colorCodeGreen
	colorCodeYellow
	colorCodeBlue
	colorCodeMagenta
	colorCodeCyan
	colorCodeWhite
)

type colorCode uint8

var (
	colorLevelStr = []string{
		coloredText("FATAL", colorCodeRed, true),
		coloredText("ERROR", colorCodeRed, false),
		coloredText("WARN ", colorCodeYellow, false),
		coloredText("INFO ", colorCodeBlue, false),
		coloredText("DEBUG", colorCodeGreen, false),
		coloredText("TRACE", colorCodeMagenta, false),
	}
)

func coloredText(text string, color colorCode, bold bool) string {
	b := ""
	if bold {
		b = colorBold
	}
	return fmt.Sprintf("%s%s%dm%s%s", colorPrefix, b, color, text, colorReset)
}
