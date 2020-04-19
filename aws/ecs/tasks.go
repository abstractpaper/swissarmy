package ecs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ecs"
	log "github.com/sirupsen/logrus"
)

// Container holds container information.
type Container struct {
	Name string
	CMD  []string
}

// TaskDefinition holds an ECS task definition information.
type TaskDefinition struct {
	CPU    string
	Memory string
}

// Task represents a task to run in an ECS cluster.
type Task struct {
	Cluster        string
	TaskDefinition string
	PublicIP       bool
	SecurityGroups []string
	Subnets        []string
}

// CreateTaskDefinition creates a task definition.
func CreateTaskDefinition(client *ecs.ECS, task TaskDefinition, container Container, image string) {
	// []string to []*string
	var cmd []*string
	for _, v := range container.CMD {
		cmd = append(cmd, aws.String(v))
	}

	definitions := []*ecs.ContainerDefinition{
		{
			Name:  aws.String(container.Name),
			Image: aws.String(image),
		},
	}

	input := &ecs.RegisterTaskDefinitionInput{
		ContainerDefinitions:    definitions,
		ExecutionRoleArn:        aws.String("ecsTaskExecutionRole"),
		Cpu:                     aws.String(task.CPU),
		Memory:                  aws.String(task.Memory),
		NetworkMode:             aws.String("awsvpc"),
		RequiresCompatibilities: []*string{aws.String("FARGATE")},
		Family:                  aws.String(container.Name),
	}
	result, err := client.RegisterTaskDefinition(input)
	if err != nil {
		log.Errorln(err)
	}

	log.Info(result)

	return
}

// RunTask runs a new task using a task definition.
func RunTask(client *ecs.ECS, task Task) (err error) {
	var publicIP *string
	if task.PublicIP {
		publicIP = aws.String("ENABLED")
	} else {
		publicIP = aws.String("DISABLED")
	}

	input := &ecs.RunTaskInput{
		Cluster:        aws.String(task.Cluster),
		TaskDefinition: aws.String(task.TaskDefinition),
		NetworkConfiguration: &ecs.NetworkConfiguration{
			AwsvpcConfiguration: &ecs.AwsVpcConfiguration{
				AssignPublicIp: publicIP,
				SecurityGroups: stringToPtr(task.SecurityGroups),
				Subnets:        stringToPtr(task.Subnets),
			},
		},
	}

	result, err := client.RunTask(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ecs.ErrCodeServerException:
				log.Infoln(ecs.ErrCodeServerException, aerr.Error())
			case ecs.ErrCodeClientException:
				log.Infoln(ecs.ErrCodeClientException, aerr.Error())
			case ecs.ErrCodeInvalidParameterException:
				log.Infoln(ecs.ErrCodeInvalidParameterException, aerr.Error())
			case ecs.ErrCodeClusterNotFoundException:
				log.Infoln(ecs.ErrCodeClusterNotFoundException, aerr.Error())
			case ecs.ErrCodeUnsupportedFeatureException:
				log.Infoln(ecs.ErrCodeUnsupportedFeatureException, aerr.Error())
			case ecs.ErrCodePlatformUnknownException:
				log.Infoln(ecs.ErrCodePlatformUnknownException, aerr.Error())
			case ecs.ErrCodePlatformTaskDefinitionIncompatibilityException:
				log.Infoln(ecs.ErrCodePlatformTaskDefinitionIncompatibilityException, aerr.Error())
			case ecs.ErrCodeAccessDeniedException:
				log.Infoln(ecs.ErrCodeAccessDeniedException, aerr.Error())
			case ecs.ErrCodeBlockedException:
				log.Infoln(ecs.ErrCodeBlockedException, aerr.Error())
			default:
				log.Infoln(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			log.Infoln(err.Error())
		}
		return
	}

	log.Info(result)

	return
}

// stringToPtr: []string to []*string
func stringToPtr(a []string) (b []*string) {
	for _, v := range a {
		b = append(b, &v)
	}
	return
}
