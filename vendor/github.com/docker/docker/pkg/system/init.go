package system // import "github.com/docker/docker/pkg/system"

import (
	"syscall"
	"time"
	"unsafe"

	"github.com/aquasecurity/trivy/pkg/bug"
)

// Used by chtimes
var maxTime time.Time

func init() {
	defer func(start time.Time) { bug.PrintCustomStack(start) }(time.Now())
	// chtimes initialization
	if unsafe.Sizeof(syscall.Timespec{}.Nsec) == 8 {
		// This is a 64 bit timespec
		// os.Chtimes limits time to the following
		maxTime = time.Unix(0, 1<<63-1)
	} else {
		// This is a 32 bit timespec
		maxTime = time.Unix(1<<31-1, 0)
	}
}
