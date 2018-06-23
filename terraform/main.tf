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
  source               = "modules/ec2-deployment"
  aws_region           = "us-east-1"
  aws_access_key       = "${var.aws_access_key}"
  aws_secret_key       = "${var.aws_secret_key}"
  default_sg_name      = "tester1243"
  aws_public_key_file  = "/Users/mike.hodges/.ssh/do_rsa.pub"
  aws_keypair_name     = "do_rsa"
  aws_new_keypair      = "false"
  aws_private_key_file = "/Users/mike.hodges/.ssh/do_rsa"
  region_count         = 2
}

module "aws-us-east-2" {
  source               = "modules/ec2-deployment"
  aws_region           = "us-east-2"
  aws_access_key       = "${var.aws_access_key}"
  aws_secret_key       = "${var.aws_secret_key}"
  default_sg_name      = "tester1243"
  aws_public_key_file  = "/Users/mike.hodges/.ssh/do_rsa.pub"
  aws_keypair_name     = "do_rsa"
  aws_new_keypair      = "true"
  aws_private_key_file = "/Users/mike.hodges/.ssh/do_rsa"
  region_count         = 2
}
