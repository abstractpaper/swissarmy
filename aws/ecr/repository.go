package ecr

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ecr"
	log "github.com/sirupsen/logrus"
)

// CreateRepository creates a repository in an ECR registry.
func CreateRepository(client *ecr.ECR, name string) (err error) {
	input := &ecr.CreateRepositoryInput{
		RepositoryName: aws.String(name),
	}

	result, err := client.CreateRepository(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ecr.ErrCodeServerException:
				log.Errorln(ecr.ErrCodeServerException, aerr.Error())
			case ecr.ErrCodeInvalidParameterException:
				log.Errorln(ecr.ErrCodeInvalidParameterException, aerr.Error())
			case ecr.ErrCodeInvalidTagParameterException:
				log.Errorln(ecr.ErrCodeInvalidTagParameterException, aerr.Error())
			case ecr.ErrCodeTooManyTagsException:
				log.Errorln(ecr.ErrCodeTooManyTagsException, aerr.Error())
			case ecr.ErrCodeRepositoryAlreadyExistsException:
				log.Warnln(ecr.ErrCodeRepositoryAlreadyExistsException, aerr.Error())
			case ecr.ErrCodeLimitExceededException:
				log.Errorln(ecr.ErrCodeLimitExceededException, aerr.Error())
			default:
				log.Errorln(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			log.Errorln(err.Error())
		}
		return
	}

	log.Info("Created repository: ", *result.Repository.RepositoryName)
	return
}
