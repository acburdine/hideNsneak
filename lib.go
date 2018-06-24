package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		log.Fatal()
	}
}

func execCmd(binary string, args []string) {
	var stdout, stderr bytes.Buffer

	cmd := exec.Command(binary, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(stderr.String())
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	fmt.Println(stdout.String())
}

func terraformApply() {
	binary, err := exec.LookPath("terraform")

	checkErr(err)

	//Initializing Terraform
	fmt.Println("init")
	args := []string{"init", "-input=false", "terraform"}
	execCmd(binary, args)

	//Planning Terraform changes and saving plan to file tfplan
	args = []string{"plan", "-out=terraform/tfplan", "-input=false", "-var-file=terraform/terraform.tfvars", "terraform"}
	execCmd(binary, args)

	//Applying Changes Identified in tfplan
	args = []string{"apply", "-input=false", "terraform/tfplan"}
	execCmd(binary, args)

}

func removeSpaces(input string) (newString string) {
	newString = strings.ToLower(input)
	newString = strings.Replace(newString, " ", "_", -1)

	return
}

func providerCheck(providerArray []string) bool {
	for _, p := range providerArray {
		if strings.ToUpper(p) != "EC2" &&
			strings.ToUpper(p) != "DO" &&
			strings.ToUpper(p) != "GOOGLE" &&
			strings.ToUpper(p) != "AZURE" &&
			strings.ToUpper(p) != "AZURECDN" &&
			strings.ToUpper(p) != "APIGATEWAY" {
			fmt.Println(unknownProvider)
			return false
		}
	}
	return true
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

//checkEc2KeyExistence queries the Amazon EC2 API for the keypairs with the specified keyname
//Returns true if the resulting array is > 0, false otherwise
func checkEC2KeyExistence(secret string, accessID string, region string, keyName string) bool {
	svc := ec2.New(session.New(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessID, secret, ""),
	}))
	keyPairOutput, _ := svc.DescribeKeyPairs(&ec2.DescribeKeyPairsInput{
		KeyNames: aws.StringSlice([]string{keyName}),
	})

	if len(keyPairOutput.KeyPairs) == 0 {
		return false
	}
	return true
}

//checkEc2KeyExistence queries the Amazon EC2 API for the security groups
//with the specified name
//Returns true if the resulting array is > 0, false otherwise
func checkEC2SecurityGroupExistence(secret string, accessID string, region string, securityGroupName string) bool {
	svc := ec2.New(session.New(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessID, secret, ""),
	}))
	securityGroupOutput, _ := svc.DescribeSecurityGroups(&ec2.DescribeSecurityGroupsInput{
		GroupNames: aws.StringSlice([]string{securityGroupName}),
	})

	if len(securityGroupOutput.SecurityGroups) == 0 {
		return false
	}
	return true
}

//createSingleSOCKS initiates a SOCKS Proxy on the local host with the specifed ipv4 address
//returns a pointer to the specified OS process so that we can kill it effictively
func createSingleSOCKS(privateKey string, username string, ipv4 string, port int) *os.Process {
	portString := strconv.Itoa(port)
	cmd := exec.Command("ssh", "-N", "-D", portString, "-o", "StrictHostKeyChecking=no", "-i", privateKey, fmt.Sprintf(username+"@%s", ipv4))
	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		return nil
	}
	return cmd.Process
}
