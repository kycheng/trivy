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
	AddRule("LoneAnonymousOperation", func(observers *Events, addError AddErrFunc) {
		observers.OnOperation(func(walker *Walker, operation *ast.OperationDefinition) {
			if operation.Name == "" && len(walker.Document.Operations) > 1 {
				addError(
					Message(`This anonymous operation must be the only defined operation.`),
					At(operation.Position),
				)
			}
		})
	})
}
