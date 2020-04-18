package ecr

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
)

// NewClient creates a new ECR client from a session.
func NewClient(session *session.Session) (client *ecr.ECR) {
	client = ecr.New(session)
	return
}
