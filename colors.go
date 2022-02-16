package zzzlog

import (
	"fmt"
	"strings"
)

const (
	formatPrefix    = "\x1b["
	formatSuffix    = "m"
	optionSeperator = ";"
	optionReset     = "0"
	optionBold      = "1"
	optionUnderline = "4"
	optionInvert    = "7"
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
	colorCodeNone = 0
)

type colorCode uint8

type levelColorMap map[level]*levelColorInfo

type levelColorInfo struct {
	color     colorCode
	bold      bool
	underline bool
	invert    bool
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
		result[idx] = coloredText(n, c)
	}
	return result
}

func coloredText(text string, format *levelColorInfo) string {
	if format.color == colorCodeNone && !format.bold && !format.underline && !format.invert {
		return text
	}

	firstOption := true
	var result strings.Builder
	result.WriteString(formatPrefix)
	if format.bold {
		if !firstOption {
			result.WriteString(optionSeperator)
		}
		firstOption = false
		result.WriteString(optionBold)
	}
	if format.underline {
		if !firstOption {
			result.WriteString(optionSeperator)
		}
		firstOption = false
		result.WriteString(optionUnderline)
	}
	if format.invert {
		if !firstOption {
			result.WriteString(optionSeperator)
		}
		firstOption = false
		result.WriteString(optionInvert)
	}
	if format.color != colorCodeNone {
		if !firstOption {
			result.WriteString(optionSeperator)
		}
		result.WriteString(fmt.Sprintf("%d", format.color))
	}
	result.WriteString(formatSuffix)
	result.WriteString(text)
	result.WriteString(formatPrefix)
	result.WriteString(optionReset)
	result.WriteString(formatSuffix)
	return result.String()
}
