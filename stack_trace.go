package zzzlog

import "runtime"

func stackTraces() []byte {
	buf := make([]byte, 1<<16)
	size := runtime.Stack(buf, true)
	return buf[:size]
}
