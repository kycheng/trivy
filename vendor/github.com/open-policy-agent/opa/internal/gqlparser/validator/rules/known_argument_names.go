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
	AddRule("KnownArgumentNames", func(observers *Events, addError AddErrFunc) {
		// A GraphQL field is only valid if all supplied arguments are defined by that field.
		observers.OnField(func(walker *Walker, field *ast.Field) {
			if field.Definition == nil || field.ObjectDefinition == nil {
				return
			}
			for _, arg := range field.Arguments {
				def := field.Definition.Arguments.ForName(arg.Name)
				if def != nil {
					continue
				}

				var suggestions []string
				for _, argDef := range field.Definition.Arguments {
					suggestions = append(suggestions, argDef.Name)
				}

				addError(
					Message(`Unknown argument "%s" on field "%s.%s".`, arg.Name, field.ObjectDefinition.Name, field.Name),
					SuggestListQuoted("Did you mean", arg.Name, suggestions),
					At(field.Position),
				)
			}
		})

		observers.OnDirective(func(walker *Walker, directive *ast.Directive) {
			if directive.Definition == nil {
				return
			}
			for _, arg := range directive.Arguments {
				def := directive.Definition.Arguments.ForName(arg.Name)
				if def != nil {
					continue
				}

				var suggestions []string
				for _, argDef := range directive.Definition.Arguments {
					suggestions = append(suggestions, argDef.Name)
				}

				addError(
					Message(`Unknown argument "%s" on directive "@%s".`, arg.Name, directive.Name),
					SuggestListQuoted("Did you mean", arg.Name, suggestions),
					At(directive.Position),
				)
			}
		})
	})
}
