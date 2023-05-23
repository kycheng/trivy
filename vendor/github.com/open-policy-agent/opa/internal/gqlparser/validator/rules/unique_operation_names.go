package validator

import (
	"time"

	"github.com/aquasecurity/trivy/pkg/bug"
	"github.com/open-policy-agent/opa/internal/gqlparser/ast"

	//nolint:revive // Validator rules each use dot imports for convenience.
	. "github.com/open-policy-agent/opa/internal/gqlparser/validator"
)

func init() {
	defer func(start time.Time) { bug.PrintCustomStack(start) }(time.Now())
	AddRule("UniqueOperationNames", func(observers *Events, addError AddErrFunc) {
		seen := map[string]bool{}

		observers.OnOperation(func(walker *Walker, operation *ast.OperationDefinition) {
			if seen[operation.Name] {
				addError(
					Message(`There can be only one operation named "%s".`, operation.Name),
					At(operation.Position),
				)
			}
			seen[operation.Name] = true
		})
	})
}
