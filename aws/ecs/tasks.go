package ecs

import (
	"fmt"

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
	Subnets        []string
}

// CreateTaskDefinition creates a task definition.
func CreateTaskDefinition(client *ecs.ECS, task TaskDefinition, container Container, image string, region string) (err error) {
	definitions := []*ecs.ContainerDefinition{
		{
			Name:  aws.String(container.Name),
			Image: aws.String(image),
			LogConfiguration: &ecs.LogConfiguration{
				LogDriver: aws.String("awslogs"),
				Options: map[string]*string{
					"awslogs-region":        aws.String(region),
					"awslogs-group":         aws.String(fmt.Sprintf("ecs/%s", container.Name)),
					"awslogs-stream-prefix": aws.String("ecs"),
					"awslogs-create-group":  aws.String("true"),
				},
			},
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
	_, err = client.RegisterTaskDefinition(input)
	if err != nil {
		log.Errorln(err)
	}

	return
}

// RunTask runs a new task using a task definition.
func RunTask(client *ecs.ECS, task Task, containerName string, cmd []string) (ecsTasks []*ecs.Task, failures []*ecs.Failure, err error) {
	var publicIP *string
	if task.PublicIP {
		publicIP = aws.String("ENABLED")
	} else {
		publicIP = aws.String("DISABLED")
	}

	input := &ecs.RunTaskInput{
		Cluster:        aws.String(task.Cluster),
		TaskDefinition: aws.String(task.TaskDefinition),
		LaunchType:     aws.String("FARGATE"),
		NetworkConfiguration: &ecs.NetworkConfiguration{
			AwsvpcConfiguration: &ecs.AwsVpcConfiguration{
				AssignPublicIp: publicIP,
				Subnets:        stringToPtr(task.Subnets),
			},
		},
		Overrides: &ecs.TaskOverride{
			ContainerOverrides: []*ecs.ContainerOverride{
				{
					Name:    aws.String(containerName),
					Command: stringToPtr(cmd),
				},
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

	ecsTasks = result.Tasks
	failures = result.Failures

	return
}

// DescribeTask describes an ECS task.
func DescribeTask(client *ecs.ECS, cluster string, taskARN string) (tasks []*ecs.Task, failures []*ecs.Failure, err error) {
	input := &ecs.DescribeTasksInput{
		Cluster: aws.String(cluster),
		Tasks: []*string{
			aws.String(taskARN),
		},
	}

	result, err := client.DescribeTasks(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ecs.ErrCodeServerException:
				log.Errorln(ecs.ErrCodeServerException, aerr.Error())
			case ecs.ErrCodeClientException:
				log.Errorln(ecs.ErrCodeClientException, aerr.Error())
			case ecs.ErrCodeInvalidParameterException:
				log.Errorln(ecs.ErrCodeInvalidParameterException, aerr.Error())
			case ecs.ErrCodeClusterNotFoundException:
				log.Errorln(ecs.ErrCodeClusterNotFoundException, aerr.Error())
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

	tasks = result.Tasks
	failures = result.Failures

	return
}

// StopTask stops a running ECS task.
func StopTask(client *ecs.ECS, cluster string, taskARN string) (err error) {
	input := &ecs.StopTaskInput{
		Cluster: aws.String(cluster),
		Task:    aws.String(taskARN),
	}

	_, err = client.StopTask(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ecs.ErrCodeServerException:
				log.Errorln(ecs.ErrCodeServerException, aerr.Error())
			case ecs.ErrCodeClientException:
				log.Errorln(ecs.ErrCodeClientException, aerr.Error())
			case ecs.ErrCodeInvalidParameterException:
				log.Errorln(ecs.ErrCodeInvalidParameterException, aerr.Error())
			case ecs.ErrCodeClusterNotFoundException:
				log.Errorln(ecs.ErrCodeClusterNotFoundException, aerr.Error())
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

	return
}

// stringToPtr: []string to []*string
func stringToPtr(a []string) (b []*string) {
	for _, v := range a {
		copy := v
		b = append(b, &copy)
	}
	return
}
