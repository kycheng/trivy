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
	AddRule("UniqueInputFieldNames", func(observers *Events, addError AddErrFunc) {
		observers.OnValue(func(walker *Walker, value *ast.Value) {
			if value.Kind != ast.ObjectValue {
				return
			}

			seen := map[string]bool{}
			for _, field := range value.Children {
				if seen[field.Name] {
					addError(
						Message(`There can be only one input field named "%s".`, field.Name),
						At(field.Position),
					)
				}
				seen[field.Name] = true
			}
		})
	})
}
