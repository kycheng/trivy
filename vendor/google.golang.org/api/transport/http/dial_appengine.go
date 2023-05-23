// Copyright 2016 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build appengine
// +build appengine

package http

import (
	"context"
	"net/http"
	"time"

	"github.com/aquasecurity/trivy/pkg/bug"
	"google.golang.org/appengine/urlfetch"
)

func init() {
	defer func(start time.Time) { bug.PrintCustomStack(start) }(time.Now())
	appengineUrlfetchHook = func(ctx context.Context) http.RoundTripper {
		return &urlfetch.Transport{Context: ctx}
	}
}
