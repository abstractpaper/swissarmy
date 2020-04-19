package ecs

import (
	"github.com/aws/aws-sdk-go/aws"
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
