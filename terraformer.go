package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"os/exec"
)

const variables = `
variable "do_token" {}
variable "aws_access_key" {}
variable "aws_secret_key" {}

variable "azure_tenant_id" {}

variable "azure_client_id" {}

variable "azure_client_secret" {}

variable "azure_subscription_id" {}
`

const state = `terraform {
	backend "s3" {
	  bucket         = "hidensneak-terraform"
	  key            = "filename.tfstate"
	  dynamodb_table = "terraform-state-lock-dynamo"
	  region         = "us-east-1"
	  encrypt        = true
	}
  }`

const tfvars = `aws_access_key = "AKIAIPNLFMEFDYNGBSLA"

aws_secret_key = "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"

do_token = "0f7e05467852e4d668b20df1cd6e5574747af7eda4dda0f72021a0e0fa4b4ffd"

azure_tenant_id = "a8b80a08-1034-4a3b-b61f-72328ffbf63f"

azure_client_id = "4c76ff10-7f72-4a6b-a226-b4bf0bd7b789"

azure_client_secret = "x/V72EjrHtl0jFq3z+2euyXzu5lnWgw7KcrVDQy2wic="

azure_subscription_id = "7704ddcf-943b-4039-a051-9e3bd167afae"
`

const aws_module = `
module "aws-{{.Region}}" {
	source         = "modules/ec2-deployment"
	aws_region     = "{{.Region}}"
	aws_access_key = "${var.aws_access_key}"
	aws_secret_key = "${var.aws_secret_key}"
	default_sg_name = "{{.Security_Group}}"
	aws_keypair_file     = "{{.Keypair_File}}"
	aws_keypair_name     = "{{.Keypair_Name}}"
	aws_new_keypair      = "{{.New_Keypair}}"
	region_count   = {{.Count}}
  }
`

// module "aws-us-east-1" {
// 	source           = "modules/ec2-deployment"
// 	aws_region       = "us-east-1"
// 	aws_access_key   = "${var.aws_access_key}"
// 	aws_secret_key   = "${var.aws_secret_key}"
// 	default_sg_name  = "tester-us-east-1"
// 	aws_keypair_file = "/Users/mike.hodges/.ssh/do_rsa.pub"
// 	aws_keypair_name = "do_rsa"
// 	aws_new_keypair  = false
// 	region_count     = 0
//   }

type Aws_Deployer struct {
	Count          int
	Region         string
	Security_Group string
	Custom_Ami     string
	Keypair_File   string
	Keypair_Name   string
	New_Keypair    bool
}

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
	args := []string{"init", "-input=false"}
	execCmd(binary, args)

	//Planning Terraform changes and saving plan to file tfplan
	args = []string{"plan", "-out=tfplan", "-input=false"}
	execCmd(binary, args)

	//Applying Changes Identified in tfplan
	args = []string{"apply", "-input=false", "tfplan"}
	execCmd(binary, args)

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
	mainFile.Write([]byte(state))

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
