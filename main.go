package main

import (
	"os"
)

func main() {
	//Create files for terraform
	mainFile, err := os.Create("terraform/main.tf")
	checkErr(err)
	defer mainFile.Close()

	varFile, err := os.Create("terraform/variables.tf")
	checkErr(err)
	defer varFile.Close()

	tfvarsFile, err := os.Create("terraform/terraform.tfvars")
	checkErr(err)
	defer tfvarsFile.Close()

	//Writing constants for terraform
	mainFile.Write([]byte(state))
	varFile.Write([]byte(variables))
	tfvarsFile.Write([]byte(tfvars))

	//TODO make user input
	//create an emoty struct from readList
	//assign each user input to values from struct
	//ass in that struct as userInput variable in createMasterList

	// parseUserInputIntoReadList()

	tester1 := ec2Deployer{
		Count:         1,
		Region:        "us-east-1",
		SecurityGroup: "tester1243",
		PublicKeyFile: "/Users/mike.hodges/.ssh/do_rsa.pub",
		KeypairName:   "do_rsa",
		NewKeypair:    false,
	}

	tester2 := ec2Deployer{
		Count:         1,
		Region:        "us-west-1",
		SecurityGroup: "tester1243",
		PublicKeyFile: "/Users/mike.hodges/.ssh/do_rsa.pub",
		KeypairName:   "do_rsa",
		NewKeypair:    false,
	}
	tester3 := ec2Deployer{
		Count:         1,
		Region:        "eu-west-1",
		SecurityGroup: "tester1243",
		PublicKeyFile: "/Users/mike.hodges/.ssh/do_rsa.pub",
		KeypairName:   "do_rsa",
		NewKeypair:    false,
	}

	ec2Stuff := [...]ec2Deployer{tester1, tester2, tester3}
	azureCdnStuff := []azureCdnDeployer{}
	azureStuff := []azureDeployer{}
	cloudFrontStuff := []cloudFrontDeployer{}
	doStuff := []digitalOceanDeployer{}
	gcpStuff := []googleCloudDeployer{}
	apiStuff := []apiGatewayDeployer{}

	//Creating a test readist
	userInput := readList{
		ec2Stuff[:],
		azureCdnStuff,
		azureStuff,
		cloudFrontStuff,
		doStuff,
		gcpStuff,
		apiStuff,
	}

	createMasterList(userInput) //TODO: userInput is whatever masterStruct they want to pass in

	//Opening Main.tf to append parsed template
	file, err := os.OpenFile("terraform/main.tf", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	checkErr(err)

	//Writing the masterString to file. masterString was instantiated in master_list.go
	_, err = file.WriteString(masterString)
	checkErr(err)

	err = mainFile.Close()
	checkErr(err)

	//Perform all the terraform deployment
	terraformApply()
}
