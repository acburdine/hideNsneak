locals {
  region_count_final = "${merge(var.region_count_template, var.region_count)}"
}

module "aws-us-east-1" {
  source = "ec2-region-deployment"

  providers = {
    "aws" = "aws.us-east-1"
  }

  #TODO: Fix Security Groups
  # default_sg_name      = "test"
  # aws_sg_id            = ""
  instance_count = "${local.region_count_final["us-east-1"]}"

  custom_ami        = "${var.custom_ami}"
  aws_instance_type = "${var.aws_instance_type}"
  ec2_default_user  = "${var.ec2_default_user}"

  aws_region = "us-east-1"

  #TODO: Add a map for regions to keypairs
  aws_keypair_name     = "${var.aws_keypair_name}"
  aws_private_key_file = "${var.aws_private_key_file}"
  aws_public_key_file  = "${var.aws_public_key_file}"
}

module "aws-us-east-2" {
  source = "ec2-region-deployment"

  providers = {
    "aws" = "aws.us-east-2"
  }

  #TODO: Fix Security Groups
  # default_sg_name      = "var"
  # aws_sg_id            = ""
  instance_count = "${local.region_count_final["us-east-2"]}"

  custom_ami        = "${var.custom_ami}"
  aws_instance_type = "${var.aws_instance_type}"
  ec2_default_user  = "${var.ec2_default_user}"

  aws_region           = "us-east-2"
  aws_keypair_name     = "${var.aws_keypair_name}"
  aws_private_key_file = "${var.aws_private_key_file}"
  aws_public_key_file  = "${var.aws_public_key_file}"
}

module "aws-us-west-1" {
  source = "ec2-region-deployment"

  providers = {
    "aws" = "aws.us-west-1"
  }

  #TODO: Fix Security Groups
  # default_sg_name      = "test"
  # aws_sg_id            = ""
  instance_count = "${local.region_count_final["us-west-1"]}"

  custom_ami        = "${var.custom_ami}"
  aws_instance_type = "${var.aws_instance_type}"
  ec2_default_user  = "${var.ec2_default_user}"

  aws_region           = "us-west-1"
  aws_keypair_name     = "${var.aws_keypair_name}"
  aws_private_key_file = "${var.aws_private_key_file}"
  aws_public_key_file  = "${var.aws_public_key_file}"
}

module "aws-us-west-2" {
  source = "ec2-region-deployment"

  providers = {
    "aws" = "aws.us-west-2"
  }

  #TODO: Fix Security Groups
  # default_sg_name      = "test"
  # aws_sg_id            = ""
  instance_count = "${local.region_count_final["us-west-2"]}"

  custom_ami        = "${var.custom_ami}"
  aws_instance_type = "${var.aws_instance_type}"
  ec2_default_user  = "${var.ec2_default_user}"

  aws_region           = "us-west-2"
  aws_keypair_name     = "${var.aws_keypair_name}"
  aws_private_key_file = "${var.aws_private_key_file}"
  aws_public_key_file  = "${var.aws_public_key_file}"
}

module "aws-ca-central-1" {
  source = "ec2-region-deployment"

  providers = {
    "aws" = "aws.ca-central-1"
  }

  #TODO: Fix Security Groups
  # default_sg_name      = "test"
  # aws_sg_id            = ""
  instance_count = "${local.region_count_final["ca-central-1"]}"

  custom_ami        = "${var.custom_ami}"
  aws_instance_type = "${var.aws_instance_type}"
  ec2_default_user  = "${var.ec2_default_user}"

  aws_region           = "ca-central-1"
  aws_keypair_name     = "${var.aws_keypair_name}"
  aws_private_key_file = "${var.aws_private_key_file}"
  aws_public_key_file  = "${var.aws_public_key_file}"
}

module "aws-eu-central-1" {
  source = "ec2-region-deployment"

  providers = {
    "aws" = "aws.eu-central-1"
  }

  #TODO: Fix Security Groups
  # default_sg_name      = "test"
  # aws_sg_id            = ""
  instance_count = "${local.region_count_final["eu-central-1"]}"

  custom_ami        = "${var.custom_ami}"
  aws_instance_type = "${var.aws_instance_type}"
  ec2_default_user  = "${var.ec2_default_user}"

  aws_region           = "eu-central-1"
  aws_keypair_name     = "${var.aws_keypair_name}"
  aws_private_key_file = "${var.aws_private_key_file}"
  aws_public_key_file  = "${var.aws_public_key_file}"
}

module "aws-eu-west-1" {
  source = "ec2-region-deployment"

  providers = {
    "aws" = "aws.eu-west-1"
  }

  #TODO: Fix Security Groups
  # default_sg_name      = "test"
  # aws_sg_id            = ""
  instance_count = "${local.region_count_final["eu-west-1"]}"

  custom_ami        = "${var.custom_ami}"
  aws_instance_type = "${var.aws_instance_type}"
  ec2_default_user  = "${var.ec2_default_user}"

  aws_region           = "eu-west-1"
  aws_keypair_name     = "${var.aws_keypair_name}"
  aws_private_key_file = "${var.aws_private_key_file}"
  aws_public_key_file  = "${var.aws_public_key_file}"
}

module "aws-eu-west-2" {
  source = "ec2-region-deployment"

  providers = {
    "aws" = "aws.eu-west-2"
  }

  #TODO: Fix Security Groups
  # default_sg_name      = "test"
  # aws_sg_id            = ""
  instance_count = "${local.region_count_final["eu-west-2"]}"

  custom_ami        = "${var.custom_ami}"
  aws_instance_type = "${var.aws_instance_type}"
  ec2_default_user  = "${var.ec2_default_user}"

  aws_region           = "eu-west-2"
  aws_keypair_name     = "${var.aws_keypair_name}"
  aws_private_key_file = "${var.aws_private_key_file}"
  aws_public_key_file  = "${var.aws_public_key_file}"
}

module "aws-eu-west-3" {
  source = "ec2-region-deployment"

  providers = {
    "aws" = "aws.eu-west-3"
  }

  #TODO: Fix Security Groups
  # default_sg_name      = "test"
  # aws_sg_id            = ""
  instance_count = "${local.region_count_final["eu-west-3"]}"

  custom_ami        = "${var.custom_ami}"
  aws_instance_type = "${var.aws_instance_type}"
  ec2_default_user  = "${var.ec2_default_user}"

  aws_region           = "eu-west-3"
  aws_keypair_name     = "${var.aws_keypair_name}"
  aws_private_key_file = "${var.aws_private_key_file}"
  aws_public_key_file  = "${var.aws_public_key_file}"
}

module "aws-ap-northeast-1" {
  source = "ec2-region-deployment"

  providers = {
    "aws" = "aws.ap-northeast-1"
  }

  #TODO: Fix Security Groups
  # default_sg_name      = "test"
  # aws_sg_id            = ""
  instance_count = "${local.region_count_final["ap-northeast-1"]}"

  custom_ami        = "${var.custom_ami}"
  aws_instance_type = "${var.aws_instance_type}"
  ec2_default_user  = "${var.ec2_default_user}"

  aws_region           = "ap-northeast-1"
  aws_keypair_name     = "${var.aws_keypair_name}"
  aws_private_key_file = "${var.aws_private_key_file}"
  aws_public_key_file  = "${var.aws_public_key_file}"
}

module "aws-ap-northeast-2" {
  source = "ec2-region-deployment"

  providers = {
    "aws" = "aws.ap-northeast-2"
  }

  #TODO: Fix Security Groups
  # default_sg_name      = "test"
  # aws_sg_id            = ""
  instance_count = "${local.region_count_final["ap-northeast-2"]}"

  custom_ami        = "${var.custom_ami}"
  aws_instance_type = "${var.aws_instance_type}"
  ec2_default_user  = "${var.ec2_default_user}"

  aws_region           = "ap-northeast-2"
  aws_keypair_name     = "${var.aws_keypair_name}"
  aws_private_key_file = "${var.aws_private_key_file}"
  aws_public_key_file  = "${var.aws_public_key_file}"
}

module "aws-ap-southeast-1" {
  source = "ec2-region-deployment"

  providers = {
    "aws" = "aws.ap-southeast-1"
  }

  #TODO: Fix Security Groups
  # default_sg_name      = "test"
  # aws_sg_id            = ""
  instance_count = "${local.region_count_final["ap-southeast-1"]}"

  custom_ami        = "${var.custom_ami}"
  aws_instance_type = "${var.aws_instance_type}"
  ec2_default_user  = "${var.ec2_default_user}"

  aws_region           = "ap-southeast-1"
  aws_keypair_name     = "${var.aws_keypair_name}"
  aws_private_key_file = "${var.aws_private_key_file}"
  aws_public_key_file  = "${var.aws_public_key_file}"
}

module "aws-ap-southeast-2" {
  source = "ec2-region-deployment"

  providers = {
    "aws" = "aws.ap-southeast-2"
  }

  #TODO: Fix Security Groups
  # default_sg_name      = "test"
  # aws_sg_id            = ""
  instance_count = "${local.region_count_final["ap-southeast-2"]}"

  custom_ami        = "${var.custom_ami}"
  aws_instance_type = "${var.aws_instance_type}"
  ec2_default_user  = "${var.ec2_default_user}"

  aws_region           = "ap-southeast-2"
  aws_keypair_name     = "${var.aws_keypair_name}"
  aws_private_key_file = "${var.aws_private_key_file}"
  aws_public_key_file  = "${var.aws_public_key_file}"
}

module "aws-ap-south-1" {
  source = "ec2-region-deployment"

  providers = {
    "aws" = "aws.ap-south-1"
  }

  #TODO: Fix Security Groups
  # default_sg_name      = "test"
  # aws_sg_id            = ""
  instance_count = "${local.region_count_final["ap-south-1"]}"

  custom_ami        = "${var.custom_ami}"
  aws_instance_type = "${var.aws_instance_type}"
  ec2_default_user  = "${var.ec2_default_user}"

  aws_region           = "ap-south-1"
  aws_keypair_name     = "${var.aws_keypair_name}"
  aws_private_key_file = "${var.aws_private_key_file}"
  aws_public_key_file  = "${var.aws_public_key_file}"
}

module "aws-sa-east-1" {
  source = "ec2-region-deployment"

  providers = {
    "aws" = "aws.sa-east-1"
  }

  #TODO: Fix Security Groups
  # default_sg_name      = "test"
  # aws_sg_id            = ""
  instance_count = "${local.region_count_final["sa-east-1"]}"

  custom_ami        = "${var.custom_ami}"
  aws_instance_type = "${var.aws_instance_type}"
  ec2_default_user  = "${var.ec2_default_user}"

  aws_region           = "sa-east-1"
  aws_keypair_name     = "${var.aws_keypair_name}"
  aws_private_key_file = "${var.aws_private_key_file}"
  aws_public_key_file  = "${var.aws_public_key_file}"
}
