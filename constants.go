package constants

const Ec2Module = `
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
  }cosnts
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

const Variables = `
	variable "do_token" {}
	variable "aws_access_key" {}
	variable "aws_secret_key" {}
	variable "azure_tenant_id" {}
	variable "azure_client_id" {}
	variable "azure_cosntsclient_secret" {}
	variable "azure_subscription_id" {}
`

const State = `
	terraform {
		backend "s3" {
		  bucket         = "hidensneak-terraform"
		  key            = "filename.tfstate"
		  dynamodb_table = "terraform-state-lock-dynamo"
		  region         = "us-east-1"
		  encrypt        = true
		}
	  }
`

const Tfvars = `
	aws_access_key = "AKIAIPNLFMEFDYNGBSLA"
	aws_secret_key = "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"
	do_token = "0f7e05467852e4d668b20df1cd6e5574747af7eda4dda0f72021a0e0fa4b4ffd"
	azure_tenant_id = "a8b80a08-1034-4a3b-b61f-72328ffbf63f"
	azure_client_id = "4c76ff10-7f72-4a6b-a226-b4bf0bd7b789"
	azure_client_secret = "x/V72EjrHtl0jFq3z+2euyXzu5lnWgw7KcrVDQy2wic="
	azure_subscription_id = "7704ddcf-943b-4039-a051-9e3bd167afae"
`
