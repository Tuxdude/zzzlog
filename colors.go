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

type levelColorMap map[level]*levelColorInfo

type levelColorInfo struct {
	code colorCode
	bold bool
}

func buildColoredLevels(colorMap levelColorMap) []string {
	result := make([]string, len(orderedLevels))
	for idx, lvl := range orderedLevels {
		n, ok := levelName[lvl]
		if !ok {
			panic("Level name %q present in orderedLevels is not part of levelName map")
		}
		c, ok := colorMap[lvl]
		if !ok {
			panic("Level name %q present in orderedLevels is not part of color map")
		}
		result[idx] = coloredText(n, c.code, c.bold)
	}
	return result
}

func coloredText(text string, color colorCode, bold bool) string {
	b := ""
	if bold {
		b = colorBold
	}
	return fmt.Sprintf("%s%s%dm%s%s", colorPrefix, b, color, text, colorReset)
}
