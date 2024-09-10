package zzzlog

import "strings"

var (
	defaultFormatStr = []string{
		buildDefaultFormat(0),
		buildDefaultFormat(1),
		buildDefaultFormat(2),
		buildDefaultFormat(3),
		buildDefaultFormat(4),
		buildDefaultFormat(5),
		buildDefaultFormat(6),
		buildDefaultFormat(7),
		buildDefaultFormat(8),
		buildDefaultFormat(9),
		buildDefaultFormat(10),
		buildDefaultFormat(11),
		buildDefaultFormat(12),
		buildDefaultFormat(13),
		buildDefaultFormat(14),
		buildDefaultFormat(15),
	}
)

func defaultFormat(count int) string {
	if count < len(defaultFormatStr) {
		return defaultFormatStr[count]
	}
	return buildDefaultFormat(count)
}

func buildDefaultFormat(count int) string {
	if count == 0 {
		return ""
	}
	var result strings.Builder
	result.WriteString("%v")
	count--
	for count > 0 {
		count--
		result.WriteString(" %v")
	}
	return result.String()
}
