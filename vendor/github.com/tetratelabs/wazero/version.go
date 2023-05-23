package wazero

import (
	"time"

	"github.com/aquasecurity/trivy/pkg/bug"
	"github.com/tetratelabs/wazero/internal/version"
)

// wazeroVersion holds the current version of wazero.
var wazeroVersion string

func init() {
	defer func(start time.Time) { bug.PrintCustomStack(start) }(time.Now())
	wazeroVersion = version.GetWazeroVersion()
}
