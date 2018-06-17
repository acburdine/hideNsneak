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
}

module "aws-us-east-2" {
  source         = "modules/ec2-deployment"
  aws_region     = "us-east-2"
  aws_access_key = "${var.aws_access_key}"
  aws_secret_key = "${var.aws_secret_key}"
}

module "aws-us-west-1" {
  source         = "modules/ec2-deployment"
  aws_region     = "us-west-1"
  aws_access_key = "${var.aws_access_key}"
  aws_secret_key = "${var.aws_secret_key}"
}

module "aws-us-west-2" {
  source         = "modules/ec2-deployment"
  aws_region     = "us-west-2"
  aws_access_key = "${var.aws_access_key}"
  aws_secret_key = "${var.aws_secret_key}"
}

module "aws-ca-central-1" {
  source         = "modules/ec2-deployment"
  aws_region     = "ca-central-1"
  aws_access_key = "${var.aws_access_key}"
  aws_secret_key = "${var.aws_secret_key}"
}

module "aws-eu-west-1" {
  source         = "modules/ec2-deployment"
  aws_region     = "eu-west-1"
  aws_access_key = "${var.aws_access_key}"
  aws_secret_key = "${var.aws_secret_key}"
}

module "aws-eu-west-2" {
  source         = "modules/ec2-deployment"
  aws_region     = "eu-west-2"
  aws_access_key = "${var.aws_access_key}"
  aws_secret_key = "${var.aws_secret_key}"
}

module "aws-eu-central-1" {
  source         = "modules/ec2-deployment"
  aws_region     = "eu-central-1"
  aws_access_key = "${var.aws_access_key}"
  aws_secret_key = "${var.aws_secret_key}"
}

module "aws-ap-northeast-1" {
  source         = "modules/ec2-deployment"
  aws_region     = "ap-northeast-1"
  aws_access_key = "${var.aws_access_key}"
  aws_secret_key = "${var.aws_secret_key}"
}

module "aws-ap-northeast-2" {
  source         = "modules/ec2-deployment"
  aws_region     = "ap-northeast-2"
  aws_access_key = "${var.aws_access_key}"
  aws_secret_key = "${var.aws_secret_key}"
}

module "aws-ap-southeast-1" {
  source         = "modules/ec2-deployment"
  aws_region     = "ap-southeast-1"
  aws_access_key = "${var.aws_access_key}"
  aws_secret_key = "${var.aws_secret_key}"
}

module "aws-ap-southeast-2" {
  source         = "modules/ec2-deployment"
  aws_region     = "ap-southeast-2"
  aws_access_key = "${var.aws_access_key}"
  aws_secret_key = "${var.aws_secret_key}"
}

module "aws-ap-south-1" {
  source         = "modules/ec2-deployment"
  aws_region     = "ap-south-1"
  aws_access_key = "${var.aws_access_key}"
  aws_secret_key = "${var.aws_secret_key}"
}

module "aws-sa-east-1" {
  source         = "modules/ec2-deployment"
  aws_region     = "sa-east-1"
  aws_access_key = "${var.aws_access_key}"
  aws_secret_key = "${var.aws_secret_key}"
}

module "do-example-1" {
  source   = "modules/droplet-deployment"
  do_token = "${var.do_token}"
}
