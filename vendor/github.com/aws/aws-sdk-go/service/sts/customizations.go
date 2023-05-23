package sts

import (
	"time"

	"github.com/aquasecurity/trivy/pkg/bug"
	"github.com/aws/aws-sdk-go/aws/request"
)

func init() {
	defer func(start time.Time) { bug.PrintCustomStack(start) }(time.Now())
	initRequest = customizeRequest
}

func customizeRequest(r *request.Request) {
	r.RetryErrorCodes = append(r.RetryErrorCodes, ErrCodeIDPCommunicationErrorException)
}
