terraform {
  backend "s3" {
    bucket         = "hidensneak-terraform"
    key            = "filename.tfstate"
    dynamodb_table = "terraform-state-lock-dynamo"
    region         = "us-east-1"
    encrypt        = true
  }
}

# module "aws_tester" {
#   count  = 1
#   source = "modules/ec2-deployment"

#   aws_region = "${var.aws_regions[count.index % length(var.aws_regions)]}"

#   aws_access_key = "${var.aws_access_key}"
#   aws_secret_key = "${var.aws_secret_key}"
# }

# module "do_tester" {
#   count    = 1
#   source   = "modules/droplet-deployment"
#   do_token = "${var.do_token}"
# }

# module "gcp_tester" {
#   count  = 1
#   source = "modules/gcp-deployment"

#   gcp_region = "${var.gcp_regions[count.index % length(var.gcp_regions)]}"
# }

###AWS####
module "aws-us-east-1" {
  source         = "modules/ec2-deployment"
  aws_region     = "us-east-1"
  aws_access_key = "${var.aws_access_key}"
  aws_secret_key = "${var.aws_secret_key}"
  region_count   = 0
}

module "aws-us-east-2" {
  source         = "modules/ec2-deployment"
  aws_region     = "us-east-2"
  aws_access_key = "${var.aws_access_key}"
  aws_secret_key = "${var.aws_secret_key}"
  region_count   = 0
}

module "aws-us-west-1" {
  source         = "modules/ec2-deployment"
  aws_region     = "us-west-1"
  aws_access_key = "${var.aws_access_key}"
  aws_secret_key = "${var.aws_secret_key}"
  region_count   = 0
}

module "aws-us-west-2" {
  source         = "modules/ec2-deployment"
  aws_region     = "us-west-2"
  aws_access_key = "${var.aws_access_key}"
  aws_secret_key = "${var.aws_secret_key}"
  region_count   = 0
}

module "aws-ca-central-1" {
  source         = "modules/ec2-deployment"
  aws_region     = "ca-central-1"
  aws_access_key = "${var.aws_access_key}"
  aws_secret_key = "${var.aws_secret_key}"
  region_count   = 0
}

module "aws-eu-west-1" {
  source         = "modules/ec2-deployment"
  aws_region     = "eu-west-1"
  aws_access_key = "${var.aws_access_key}"
  aws_secret_key = "${var.aws_secret_key}"
  region_count   = 0
}

module "aws-eu-west-2" {
  source         = "modules/ec2-deployment"
  aws_region     = "eu-west-2"
  aws_access_key = "${var.aws_access_key}"
  aws_secret_key = "${var.aws_secret_key}"
  region_count   = 0
}

module "aws-eu-central-1" {
  source         = "modules/ec2-deployment"
  aws_region     = "eu-central-1"
  aws_access_key = "${var.aws_access_key}"
  aws_secret_key = "${var.aws_secret_key}"
  region_count   = 0
}

module "aws-ap-northeast-1" {
  source         = "modules/ec2-deployment"
  aws_region     = "ap-northeast-1"
  aws_access_key = "${var.aws_access_key}"
  aws_secret_key = "${var.aws_secret_key}"
  region_count   = 0
}

module "aws-ap-northeast-2" {
  source         = "modules/ec2-deployment"
  aws_region     = "ap-northeast-2"
  aws_access_key = "${var.aws_access_key}"
  aws_secret_key = "${var.aws_secret_key}"
  region_count   = 0
}

module "aws-ap-southeast-1" {
  source         = "modules/ec2-deployment"
  aws_region     = "ap-southeast-1"
  aws_access_key = "${var.aws_access_key}"
  aws_secret_key = "${var.aws_secret_key}"
  region_count   = 0
}

module "aws-ap-southeast-2" {
  source         = "modules/ec2-deployment"
  aws_region     = "ap-southeast-2"
  aws_access_key = "${var.aws_access_key}"
  aws_secret_key = "${var.aws_secret_key}"
  region_count   = 0
}

module "aws-ap-south-1" {
  source         = "modules/ec2-deployment"
  aws_region     = "ap-south-1"
  aws_access_key = "${var.aws_access_key}"
  aws_secret_key = "${var.aws_secret_key}"
  region_count   = 1
}

module "aws-sa-east-1" {
  source         = "modules/ec2-deployment"
  aws_region     = "sa-east-1"
  aws_access_key = "${var.aws_access_key}"
  aws_secret_key = "${var.aws_secret_key}"
  region_count   = 0
}

##########################################
################DO########################
module "do-example-1" {
  source   = "modules/droplet-deployment"
  do_token = "${var.do_token}"
}

##########################################
################AZURE#####################

module "azure-example-1" {
  source         = "modules/azure-deployment"
  azure_username = "${var.azure_username}"
  azure_password = "${var.azure_password}"
}

##########################################
####################GCP###################
##########################################

module "gcp-northamerica-northeast1-a" {
  source     = "modules/gcp-deployment"
  gcp_region = "northamerica-northeast1"
}

module "gcp-us-central1-f" {
  source     = "modules/gcp-deployment"
  gcp_region = "us-central1"
}

module "gcp-us-west1-b" {
  source     = "modules/gcp-deployment"
  gcp_region = "us-west1"
}

module "gcp-us-east4-c" {
  source     = "modules/gcp-deployment"
  gcp_region = "us-east4"
}

module "gcp-us-east1-c" {
  source     = "modules/gcp-deployment"
  gcp_region = "us-east1"
}

module "gcp-southamerica-east1-a" {
  source     = "modules/gcp-deployment"
  gcp_region = "southamerica-east1"
}

module "gcp-urope-north1-a" {
  source     = "modules/gcp-deployment"
  gcp_region = "europe-north1"
}

module "gcp-europe-west1-c" {
  source     = "modules/gcp-deployment"
  gcp_region = "europe-west1"
}

module "gcp-europe-west2-a" {
  source     = "modules/gcp-deployment"
  gcp_region = "europe-west2"
}

module "gcp-europe-west3-a" {
  source     = "modules/gcp-deployment"
  gcp_region = "europe-west3"
}

module "gcp-asia-south1-a" {
  source     = "modules/gcp-deployment"
  gcp_region = "asia-south1"
}

module "asia-southeast1-a" {
  source     = "modules/gcp-deployment"
  gcp_region = "asia-southeast1"
}

module "gcp-asia-east1-b" {
  source     = "modules/gcp-deployment"
  gcp_region = "asia-east1"
}

module "asia-northeast1-a" {
  source     = "modules/gcp-deployment"
  gcp_region = "asia-northeast1"
}

module "gcp-australia-southeast1-a" {
  source     = "modules/gcp-deployment"
  gcp_region = "australia-southeast1"
}
