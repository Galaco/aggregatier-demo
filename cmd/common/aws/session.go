package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

// NewSession
func NewSession(region string) *session.Session {
	cfg := aws.NewConfig().WithRegion(region)

	return session.Must(session.NewSession(cfg))
}
