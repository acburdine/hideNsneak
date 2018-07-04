package deployer

import (
	"bytes"
	"io"
	"os/exec"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

//checkEc2KeyExistence queries the Amazon EC2 API for the keypairs with the specified keyname
//Returns true if the resulting array is > 0, false otherwise
func checkEC2KeyExistance(secret string, accessID string, region string, privateKey string) (bool, string) {
	keyFingerprint := genEC2KeyFingerprint(privateKey)

	svc := ec2.New(session.New(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessID, secret, ""),
	}))
	keyPairOutput, _ := svc.DescribeKeyPairs(&ec2.DescribeKeyPairsInput{
		Filters: []*ec2.Filter{
			&ec2.Filter{
				Name:   aws.String("fingerprint"),
				Values: aws.StringSlice([]string{keyFingerprint}),
			},
		},
	})
	if len(keyPairOutput.KeyPairs) == 0 {
		return false, ""
	}
	return true, *keyPairOutput.KeyPairs[0].KeyName
}

func genEC2KeyFingerprint(privateKey string) (keyFingerprint string) {
	args1 := []string{"pkey", "-in", privateKey, "-pubout", "-outform", "DER"}
	args2 := []string{"md5", "-c"}

	pipeReader, pipeWriter := io.Pipe()

	cmd1 := exec.Command("openssl", args1...)
	cmd2 := exec.Command("openssl", args2...)

	cmd1.Stdout = pipeWriter
	cmd2.Stdin = pipeReader

	var cmd2Out bytes.Buffer

	cmd2.Stdout = &cmd2Out

	cmd1.Start()
	cmd2.Start()
	cmd1.Wait()
	pipeWriter.Close()
	cmd2.Wait()

	keyFingerprint = strings.Split(strings.TrimSpace(cmd2Out.String()), " ")[1]

	return
}

//checkEc2KeyExistence queries the Amazon EC2 API for the security groups
//with the specified name
//Returns true if the resulting array is > 0, false otherwise
func checkEC2SecurityGroupExistence(secret string, accessID string, region string, securityGroupName string) (bool, string) {
	svc := ec2.New(session.New(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessID, secret, ""),
	}))
	securityGroupOutput, _ := svc.DescribeSecurityGroups(&ec2.DescribeSecurityGroupsInput{
		GroupNames: aws.StringSlice([]string{securityGroupName}),
	})

	if len(securityGroupOutput.SecurityGroups) == 0 {
		return false, ""
	}

	return true, *securityGroupOutput.SecurityGroups[0].GroupId
}

func compareAWSInstance(instanceOne EC2ConfigWrapper, instanceTwo EC2ConfigWrapper) bool {
	//TODO: Reimplement ami checks
	if instanceOne.DefaultUser == instanceTwo.DefaultUser &&
		instanceOne.InstanceType == instanceTwo.InstanceType &&
		instanceOne.PrivateKeyFile == instanceTwo.PrivateKeyFile &&
		instanceOne.PublicKeyFile == instanceTwo.PublicKeyFile {
		return true
	}
	return false
}
