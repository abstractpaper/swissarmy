package ecs

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
)

// NewClient creates a new ECS client from a session.
func NewClient(session *session.Session) (client *ecs.ECS) {
	client = ecs.New(session)
	return
}
