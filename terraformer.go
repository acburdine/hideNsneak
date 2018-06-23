package main

import (
	"bytes"
	"html/template"
	"os"
)

type Aws_Deployer struct {
	Count          int
	Region         string
	Security_Group string
	Custom_Ami     string
	Keypair_File   string
	Keypair_Name   string
	New_Keypair    bool
}

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
	mainFile.Write([]byte(constants.Variables))

	varFile.Write([]byte(variables))

	tfvarsFile.Write([]byte(tfvars))

	//Creating a test array
	tester1 := Aws_Deployer{
		Count:          1,
		Region:         "us-east-1",
		Security_Group: "tester1243",
		Keypair_File:   "/Users/mike.hodges/.ssh/do_rsa.pub",
		Keypair_Name:   "do_rsa",
		New_Keypair:    false,
	}
	tester2 := Aws_Deployer{
		Count:          1,
		Region:         "us-west-1",
		Security_Group: "tester1243",
		Keypair_File:   "/Users/mike.hodges/.ssh/do_rsa.pub",
		Keypair_Name:   "do_rsa",
		New_Keypair:    false,
	}
	tester3 := Aws_Deployer{
		Count:          1,
		Region:         "eu-west-1",
		Security_Group: "tester1243",
		Keypair_File:   "/Users/mike.hodges/.ssh/do_rsa.pub",
		Keypair_Name:   "do_rsa",
		New_Keypair:    false,
	}
	testers := [...]Aws_Deployer{tester1, tester2, tester3}

	var totalFile string
	for _, test := range testers {
		tmpl, err := template.New("test").Parse(aws_module)

		checkErr(err)

		var tpl bytes.Buffer
		err = tmpl.Execute(&tpl, test)
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
