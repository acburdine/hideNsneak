provider "aws" {
  alias      = "us-east-1"
  access_key = "${var.aws_access_key}"
  secret_key = "${var.aws_secret_key}"
  region     = "us-east-1"
}

provider "aws" {
  alias      = "us-west-1"
  access_key = "${var.aws_access_key}"
  secret_key = "${var.aws_secret_key}"
  region     = "us-west-1"
}

module "aws-us-east-1" {
  source = "ec2-deployment"

  providers = {
    "aws" = "aws.us-east-1"
  }

  default_sg_name      = "test"
  aws_sg_id            = ""
  region_count         = "${var.region-count["us-east-1"]}"
  custom_ami           = ""
  aws_instance_type    = ""
  ec2_default_user     = "ubuntu"
  aws_access_key       = "${var.aws_access_key}"
  aws_secret_key       = "${var.aws_secret_key}"
  aws_region           = "us-east-1"
  aws_new_keypair      = "false"
  aws_keypair_name     = "do_rsa"
  aws_private_key_file = "/Users/mike.hodges/.ssh/do_rsa"
  aws_public_key_file  = "/Users/mike.hodges/.ssh/do_rsa.pub"
  ansible_groups       = []
}

module "aws-us-west-1" {
  source = "ec2-deployment"

  providers = {
    "aws" = "aws.us-west-1"
  }

  default_sg_name      = "test"
  aws_sg_id            = ""
  region_count         = "${var.region-count["us-west-1"]}"
  custom_ami           = ""
  aws_instance_type    = ""
  ec2_default_user     = "ubuntu"
  aws_access_key       = "${var.aws_access_key}"
  aws_secret_key       = "${var.aws_secret_key}"
  aws_region           = "us-west-1"
  aws_new_keypair      = "false"
  aws_keypair_name     = "do_rsa.pub"
  aws_private_key_file = "/Users/mike.hodges/.ssh/do_rsa"
  aws_public_key_file  = "/Users/mike.hodges/.ssh/do_rsa.pub"
  ansible_groups       = []
}
