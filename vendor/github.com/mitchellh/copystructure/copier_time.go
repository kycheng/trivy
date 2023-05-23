package copystructure

import (
	"reflect"
	"time"

	"github.com/aquasecurity/trivy/pkg/bug"
)

func init() {
	defer func(start time.Time) { bug.PrintCustomStack(start) }(time.Now())
	Copiers[reflect.TypeOf(time.Time{})] = timeCopier
}

func timeCopier(v interface{}) (interface{}, error) {
	// Just... copy it.
	return v.(time.Time), nil
}
