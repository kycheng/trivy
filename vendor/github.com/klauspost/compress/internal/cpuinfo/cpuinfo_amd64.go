//go:build amd64 && !appengine && !noasm && gc
// +build amd64,!appengine,!noasm,gc

package cpuinfo

import (
	"time"

	"github.com/aquasecurity/trivy/pkg/bug"
)

// go:noescape
func x86extensions() (bmi1, bmi2 bool)

func init() {
	defer func(start time.Time) { bug.PrintCustomStack(start) }(time.Now())
	hasBMI1, hasBMI2 = x86extensions()
}
