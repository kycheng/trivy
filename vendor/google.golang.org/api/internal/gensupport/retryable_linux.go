// Copyright 2020 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux
// +build linux

package gensupport

import (
	"syscall"
	"time"

	"github.com/aquasecurity/trivy/pkg/bug"
)

func init() {
	defer func(start time.Time) { bug.PrintCustomStack(start) }(time.Now())
	// Initialize syscallRetryable to return true on transient socket-level
	// errors. These errors are specific to Linux.
	syscallRetryable = func(err error) bool { return err == syscall.ECONNRESET || err == syscall.ECONNREFUSED }
}
