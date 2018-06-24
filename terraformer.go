package main

import (
	"bytes"
	"html/template"
	"os"
)

func main() {
	//Creating Files
	mainFile, err := os.Create("main.tf")
	checkErr(err)
	defer mainFile.Close()

	varFile, err := os.Create("variables.tf")
	checkErr(err)
	defer varFile.Close()

	tfvarsFile, err := os.Create("terraform.tfvars")
	checkErr(err)
	defer tfvarsFile.Close()

	//Writing Constants
	mainFile.Write([]byte(state))
	varFile.Write([]byte(variables))
	tfvarsFile.Write([]byte(tfvars))

	type BIGFUCKING struct {
		EC2JANKS []ec2Deployer,
		AZUREJANKS []azure,

	}

	//Creating a test array
	tester1 := ec2Deployer{
		Count:         1,
		Region:        "us-east-1",
		SecurityGroup: "tester1243",
		KeypairFile:   "/Users/mike.hodges/.ssh/do_rsa.pub",
		KeypairName:   "do_rsa",
		NewKeypair:    false,
	}
	tester2 := ec2Deployer{
		Count:         1,
		Region:        "us-west-1",
		SecurityGroup: "tester1243",
		KeypairFile:   "/Users/mike.hodges/.ssh/do_rsa.pub",
		KeypairName:   "do_rsa",
		NewKeypair:    false,
	}
	tester3 := ec2Deployer{
		Count:         1,
		Region:        "eu-west-1",
		SecurityGroup: "tester1243",
		KeypairFile:   "/Users/mike.hodges/.ssh/do_rsa.pub",
		KeypairName:   "do_rsa",
		NewKeypair:    false,
	}
	testers := [...]ec2Deployer{tester1, tester2, tester3}

	//Mapping all structs to a template and adding the result
	//to the totalFile string
	var totalFile string
	for _, ourStruct := range testers {
		tmpl, err := template.New("test").Parse(ec2Module)

		checkErr(err)

		var tpl bytes.Buffer
		err = tmpl.Execute(&tpl, ourStruct)
		totalFile = totalFile + tpl.String()
	}

	checkErr(err)

	//Opening Main.tf to append parsed template
	f, err := os.OpenFile("main.tf", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	checkErr(err)

	//Writing the result of the loop to file
	_, err = f.WriteString(totalFile)
	checkErr(err)

	err = f.Close()
	checkErr(err)

	//Perform all the terraform deployment
	terraformApply()

}
