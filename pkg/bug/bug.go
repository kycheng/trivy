package bug

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"time"
)

func PrintCustomStack(start time.Time) {
	if pc, file, line, ok := runtime.Caller(1); ok {
		fmt.Printf("spend %v\n", time.Since(start)/time.Nanosecond)
		fmt.Printf("▁▂▃▄▅▆▇ DEBUG.STACK ▇▆▅▄▃▂▁ %s ==> +%d %s\n%s\n", runtime.FuncForPC(pc).Name(), line, file, debug.Stack())
	}
}
