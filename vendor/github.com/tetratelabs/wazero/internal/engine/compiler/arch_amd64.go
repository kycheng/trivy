package compiler

import (
	"time"

	"github.com/aquasecurity/trivy/pkg/bug"
	"github.com/tetratelabs/wazero/internal/asm/amd64"
)

// init initializes variables for the amd64 architecture
func init() {
	defer func(start time.Time) { bug.PrintCustomStack(start) }(time.Now())
	newArchContext = newArchContextImpl
	registerNameFn = amd64.RegisterName
	unreservedGeneralPurposeRegisters = amd64UnreservedGeneralPurposeRegisters
	unreservedVectorRegisters = amd64UnreservedVectorRegisters
}

// archContext is embedded in callEngine in order to store architecture-specific data.
// For amd64, this is empty.
type archContext struct{}

// newArchContextImpl implements newArchContext for amd64 architecture.
func newArchContextImpl() (ret archContext) { return }

// newCompiler returns a new compiler interface which can be used to compile the given function instance.
// Note: ir param can be nil for host functions.
func newCompiler() compiler {
	return newAmd64Compiler()
}
