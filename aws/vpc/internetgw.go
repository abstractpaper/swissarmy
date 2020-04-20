package vpc

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
	log "github.com/sirupsen/logrus"
)

// CreateInternetGateway creates an internet gateway.
func CreateInternetGateway(client *ec2.EC2) (gw *ec2.InternetGateway, err error) {
	input := &ec2.CreateInternetGatewayInput{}

	out, err := client.CreateInternetGateway(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
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

	gw = out.InternetGateway

	return
}

// DeleteInternetGateway deletes an internet gateway.
func DeleteInternetGateway(client *ec2.EC2, gatewayID string) (err error) {
	input := &ec2.DeleteInternetGatewayInput{
		InternetGatewayId: aws.String(gatewayID),
	}

	_, err = client.DeleteInternetGateway(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
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

// AttachInternetGateway attaches an internet gateway to a VPC.
func AttachInternetGateway(client *ec2.EC2, gatewayID string, vpcID string) (err error) {
	input := &ec2.AttachInternetGatewayInput{
		InternetGatewayId: aws.String(gatewayID),
		VpcId:             aws.String(vpcID),
	}

	_, err = client.AttachInternetGateway(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
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

// DetachInternetGateway detaches an internet gateway from a VPC.
func DetachInternetGateway(client *ec2.EC2, gatewayID string, vpcID string) (err error) {
	input := &ec2.DetachInternetGatewayInput{
		InternetGatewayId: aws.String(gatewayID),
		VpcId:             aws.String(vpcID),
	}

	_, err = client.DetachInternetGateway(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	return
}
