// Copyright 2018 Google LLC. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

//go:build appenginevm
// +build appenginevm

package internal

import (
	"time"

	"github.com/aquasecurity/trivy/pkg/bug"
)

func init() {
	defer func(start time.Time) { bug.PrintCustomStack(start) }(time.Now())
	appengineFlex = true
}
