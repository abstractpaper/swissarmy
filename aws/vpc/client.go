package vpc

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// NewClient creates a new EC2 client from a session.
func NewClient(session *session.Session) (client *ec2.EC2) {
	client = ec2.New(session)
	return
}
