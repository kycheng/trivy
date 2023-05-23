// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build appengine
// +build appengine

package internal

import (
	"time"

	"github.com/aquasecurity/trivy/pkg/bug"
	"google.golang.org/appengine/urlfetch"
)

func init() {
	defer func(start time.Time) { bug.PrintCustomStack(start) }(time.Now())
	appengineClientHook = urlfetch.Client
}
