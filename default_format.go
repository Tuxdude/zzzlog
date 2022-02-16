package zzzlog

import "strings"

var (
	defaultFormatStr = []string{
		"",
		"%v",
		"%v %v",
		"%v %v %v",
		"%v %v %v %v",
		"%v %v %v %v %v",
		"%v %v %v %v %v %v",
		"%v %v %v %v %v %v %v",
		"%v %v %v %v %v %v %v %v",
		"%v %v %v %v %v %v %v %v %v",
		"%v %v %v %v %v %v %v %v %v %v",
		"%v %v %v %v %v %v %v %v %v %v %v",
		"%v %v %v %v %v %v %v %v %v %v %v %v",
		"%v %v %v %v %v %v %v %v %v %v %v %v %v",
		"%v %v %v %v %v %v %v %v %v %v %v %v %v %v",
		"%v %v %v %v %v %v %v %v %v %v %v %v %v %v %v",
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
