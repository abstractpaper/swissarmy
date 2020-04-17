package ecs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ecs"
	log "github.com/sirupsen/logrus"
)

// CreateCluster creates an ECS cluster.
func CreateCluster(client *ecs.ECS, name string) (cluster *ecs.Cluster, err error) {
	input := &ecs.CreateClusterInput{
		ClusterName: aws.String(name),
	}

	result, err := client.CreateCluster(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			log.Errorln("createCluster error: ", aerr.Error())
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			log.Errorln(err.Error())
		}
		return
	}

	cluster = result.Cluster

	return
}

// DeleteCluster deletes an ECS cluster.
func DeleteCluster(client *ecs.ECS, name string) (err error) {
	input := &ecs.DeleteClusterInput{
		Cluster: aws.String(name),
	}

	result, err := client.DeleteCluster(input)
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
			case ecs.ErrCodeClusterContainsContainerInstancesException:
				log.Errorln(ecs.ErrCodeClusterContainsContainerInstancesException, aerr.Error())
			case ecs.ErrCodeClusterContainsServicesException:
				log.Errorln(ecs.ErrCodeClusterContainsServicesException, aerr.Error())
			case ecs.ErrCodeClusterContainsTasksException:
				log.Errorln(ecs.ErrCodeClusterContainsTasksException, aerr.Error())
			case ecs.ErrCodeUpdateInProgressException:
				log.Errorln(ecs.ErrCodeUpdateInProgressException, aerr.Error())
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

	log.Info("Deleted cluster ", result.Cluster.ClusterName)
	return
}
