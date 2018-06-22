terraform {
	backend "s3" {
	  bucket         = "hidensneak-terraform"
	  key            = "filename.tfstate"
	  dynamodb_table = "terraform-state-lock-dynamo"
	  region         = "us-east-1"
	  encrypt        = true
	}
  }
module "aws-us-east-1" {
	source         = "modules/ec2-deployment"
	aws_region     = "us-east-1"
	aws_access_key = "${var.aws_access_key}"
	aws_secret_key = "${var.aws_secret_key}"
	default_sg_name = "tester1243"
	aws_keypair_file     = "/Users/mike.hodges/.ssh/do_rsa.pub"
	aws_keypair_name     = "do_rsa"
	aws_new_keypair      = "false"
	region_count   = 1
  }

module "aws-us-west-1" {
	source         = "modules/ec2-deployment"
	aws_region     = "us-west-1"
	aws_access_key = "${var.aws_access_key}"
	aws_secret_key = "${var.aws_secret_key}"
	default_sg_name = "tester1243"
	aws_keypair_file     = "/Users/mike.hodges/.ssh/do_rsa.pub"
	aws_keypair_name     = "do_rsa"
	aws_new_keypair      = "false"
	region_count   = 1
  }

module "aws-eu-west-1" {
	source         = "modules/ec2-deployment"
	aws_region     = "eu-west-1"
	aws_access_key = "${var.aws_access_key}"
	aws_secret_key = "${var.aws_secret_key}"
	default_sg_name = "tester1243"
	aws_keypair_file     = "/Users/mike.hodges/.ssh/do_rsa.pub"
	aws_keypair_name     = "do_rsa"
	aws_new_keypair      = "false"
	region_count   = 1
  }
