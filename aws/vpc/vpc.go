package vpc

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
	log "github.com/sirupsen/logrus"
)

// VPC represents a VPC.
type VPC struct {
	CIDR    string
	Subnets []Subnet
}

// Subnet represents a subnet inside a VPC.
type Subnet struct {
	CIDR string
}

// CreateVPC creates a new VPC.
func CreateVPC(client *ec2.EC2, cidr string) (vpc *ec2.Vpc, err error) {
	input := &ec2.CreateVpcInput{
		CidrBlock: aws.String(cidr),
	}

	result, err := client.CreateVpc(input)
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

	vpc = result.Vpc

	return
}

// DeleteVPC deletes a VPC.
func DeleteVPC(client *ec2.EC2, vpcID string) (err error) {
	input := &ec2.DeleteVpcInput{
		VpcId: aws.String(vpcID),
	}

	_, err = client.DeleteVpc(input)
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

// CreateSubnet creates a subnet inside in a VPC.
func CreateSubnet(client *ec2.EC2, vpcID string, cidr string) (subnet *ec2.Subnet, err error) {
	input := &ec2.CreateSubnetInput{
		CidrBlock: aws.String(cidr),
		VpcId:     aws.String(vpcID),
	}

	result, err := client.CreateSubnet(input)
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

	subnet = result.Subnet

	return
}

// DeleteSubnet deletes a subnet in a VPC.
func DeleteSubnet(client *ec2.EC2, subnetID string) (err error) {
	input := &ec2.DeleteSubnetInput{
		SubnetId: aws.String(subnetID),
	}

	_, err = client.DeleteSubnet(input)
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

// CreateRouteTable creates a route table.
func CreateRouteTable(client *ec2.EC2, vpcID string) (routeTable *ec2.RouteTable, err error) {
	input := &ec2.CreateRouteTableInput{
		VpcId: aws.String(vpcID),
	}

	result, err := client.CreateRouteTable(input)
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

	routeTable = result.RouteTable

	return
}

// CreateRoute creates a route in a route table
func CreateRoute(client *ec2.EC2, destinationCIDR string, gatewayID string, routeTableId string) (err error) {
	input := &ec2.CreateRouteInput{
		DestinationCidrBlock: aws.String(destinationCIDR),
		GatewayId:            aws.String(gatewayID),
		RouteTableId:         aws.String(routeTableId),
	}

	_, err = client.CreateRoute(input)
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

// AssociateRouteTable associates a route table to a subnet.
func AssociateRouteTable(client *ec2.EC2, routeTableID string, subnetID string) (err error) {
	input := &ec2.AssociateRouteTableInput{
		RouteTableId: aws.String(routeTableID),
		SubnetId:     aws.String(subnetID),
	}

	_, err = client.AssociateRouteTable(input)
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
