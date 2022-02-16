package zzzlog

import (
	"fmt"
	"runtime"
	"strings"
)

func callerInfo(skipFrames int) string {
	frame, ok := callerFrame(skipFrames + 1)
	if !ok {
		return formatCallerInfo("UnknownFile", 0)
	}

	ix := strings.LastIndexByte(frame.File, '/')
	if ix == -1 {
		return formatCallerInfo(frame.File, frame.Line)
	}
	ix = strings.LastIndexByte(frame.File[:ix], '/')
	if ix == -1 {
		return formatCallerInfo(frame.File, frame.Line)
	}
	return formatCallerInfo(frame.File[ix+1:], frame.Line)
}

func formatCallerInfo(file string, line int) string {
	caller := fmt.Sprintf("%s:%d", file, line)
	callerLen := len(caller)
	if callerLen > 30 {
		return caller[callerLen-30:]
	}
	return caller
}

func callerFrame(skipFrames int) (runtime.Frame, bool) {
	var frame runtime.Frame
	pc := make([]uintptr, 1)
	// We need to skip one frame to get this call site which invokes
	// runtime.Callers(), hence one more than that is what we care.
	numFrames := runtime.Callers(skipFrames+2, pc)
	if numFrames < 1 {
		return frame, false
	}

	frame, _ = runtime.CallersFrames(pc).Next()
	return frame, frame.PC != 0
}
