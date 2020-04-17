package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
	log "github.com/sirupsen/logrus"
)

// NewSession takes a region and pair of credentials and returns an AWS session.
func NewSession(region string, accessKey string, secretKey string) (sess *session.Session, err error) {
	sess, err = session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})
	if err != nil {
		log.Errorln("Error creating a session: ", err)
	}
	return
}

// NewClient creates a new ECS client from a session.
func NewClient(session *session.Session) (client *ecs.ECS) {
	client = ecs.New(session)
	return
}
