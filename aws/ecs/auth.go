package ecs

import (
	"encoding/base64"
	"strings"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"

	log "github.com/sirupsen/logrus"
)

// AuthECR authenticates with AWS ECR
func AuthECR(sess *session.Session) (authorizationToken string, endpoint string, username string, password string, err error) {
	svc := ecr.New(sess)
	input := &ecr.GetAuthorizationTokenInput{}

	result, err := svc.GetAuthorizationToken(input)
	if err != nil {
		log.Errorln(err)
	}

	authorizationToken = *result.AuthorizationData[0].AuthorizationToken
	endpoint = *result.AuthorizationData[0].ProxyEndpoint

	decoded, err := base64.StdEncoding.DecodeString(authorizationToken)
	if err != nil {
		log.Errorln(err)
	}
	s := strings.Split(string(decoded), ":")
	username, password = s[0], s[1]

	return
}
