terraform {
  backend "s3" {
    bucket         = "hidensneak-terraform"
    key            = "filename.tfstate"
    dynamodb_table = "terraform-state-lock-dynamo"
    region         = "us-east-1"
    encrypt        = true
  }
}

###AWS####
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

##########################################
################DO########################
module "do-example-1" {
  source   = "modules/droplet-deployment"
  do_token = "${var.do_token}"
}

##########################################
####################GCP###################
##########################################

module "gcp-northamerica-northeast1-a" {
  source   = "modules/gcp-deployment"
  gcp_zone = "northamerica-northeast1-a"
  region   = "northamerica-northeast1"
}

module "gcp-northamerica-northeast1-b" {
  source   = "modules/gcp-deployment"
  gcp_zone = "northamerica-northeast1-b"
  region   = "northamerica-northeast1"
}

module "gcp-northamerica-northeast1-c" {
  source   = "modules/gcp-deployment"
  gcp_zone = "northamerica-northeast1-c"
  region   = "northamerica-northeast1"
}

module "gcp-us-central1-b" {
  source   = "modules/gcp-deployment"
  gcp_zone = "us-central1-b"
  region   = "us-central1"
}

module "gcp-us-central1-c" {
  source   = "modules/gcp-deployment"
  gcp_zone = "us-central1-c"
  region   = "us-central1"
}

module "gcp-us-central1-f" {
  source   = "modules/gcp-deployment"
  gcp_zone = "us-central1-f"
  region   = "us-central1"
}

module "gcp-us-west1-b" {
  source   = "modules/gcp-deployment"
  gcp_zone = "us-west1-b"
  region   = "us-west1"
}

module "gcp-us-west1-c" {
  source   = "modules/gcp-deployment"
  gcp_zone = "us-west1-c"
  region   = "us-west1"
}

module "gcp-us-east4-a" {
  source   = "modules/gcp-deployment"
  gcp_zone = "us-east4-a"
  region   = "us-east4"
}

module "gcp-us-east4-b" {
  source   = "modules/gcp-deployment"
  gcp_zone = "us-east4-b"
  region   = "us-east4"
}

module "gcp-us-east4-c" {
  source   = "modules/gcp-deployment"
  gcp_zone = "us-east4-c"
  region   = "us-east4"
}

module "gcp-us-east1-c" {
  source   = "modules/gcp-deployment"
  gcp_zone = "us-east1-c"
  region   = "us-east1"
}

module "gcp-us-east1-d" {
  source   = "modules/gcp-deployment"
  gcp_zone = "us-east1-d"
  region   = "us-east1"
}

module "gcp-southamerica-east1-a" {
  source   = "modules/gcp-deployment"
  gcp_zone = "southamerica-east1-a"
  region   = "southamerica-east1"
}

module "gcp-southamerica-east1-b" {
  source   = "modules/gcp-deployment"
  gcp_zone = "southamerica-east1-b"
  region   = "southamerica-east1"
}

module "gcp-southamerica-east1-c" {
  source   = "modules/gcp-deployment"
  gcp_zone = "southamerica-east1-c"
  region   = "southamerica-east1"
}

module "gcp-urope-north1-a" {
  source   = "modules/gcp-deployment"
  gcp_zone = "europe-north1-a"
  region   = "europe-north1"
}

module "gcp-europe-north1-b" {
  source   = "modules/gcp-deployment"
  gcp_zone = "europe-north1-b"
  region   = "europe-north1"
}

module "gcp-europe-north1-c" {
  source   = "modules/gcp-deployment"
  gcp_zone = "europe-north1-c"
  region   = "europe-north1"
}

module "gcp-europe-west1-c" {
  source   = "modules/gcp-deployment"
  gcp_zone = "europe-west1-c"
  region   = "europe-west1"
}

module "gcp-europe-west1-d" {
  source   = "modules/gcp-deployment"
  gcp_zone = "europe-west1-d"
  region   = "europe-west1"
}

module "gcp-europe-west2-a" {
  source   = "modules/gcp-deployment"
  gcp_zone = "europe-west2-a"
  region   = "europe-west2"
}

module "gcp-europe-west2-b" {
  source   = "modules/gcp-deployment"
  gcp_zone = "europe-west2-b"
  region   = "europe-west2"
}

module "gcp-europe-west2-c" {
  source   = "modules/gcp-deployment"
  gcp_zone = "europe-west2-c"
  region   = "europe-west2"
}

module "gcp-europe-west3-a" {
  source   = "modules/gcp-deployment"
  gcp_zone = "europe-west3-a"
  region   = "europe-west3"
}

module "gcp-example-1" {
  source   = "modules/gcp-deployment"
  gcp_zone = "europe-west3-b"
  region   = "europe-west3"
}

module "gcp-europe-west3-c" {
  source   = "modules/gcp-deployment"
  gcp_zone = "europe-west3-c"
  region   = "europe-west3"
}

module "gcp-europe-west4-a" {
  source   = "modules/gcp-deployment"
  gcp_zone = "europe-west4-a"
  region   = "europe-west4"
}

module "gcp-europe-west4-b" {
  source   = "modules/gcp-deployment"
  gcp_zone = "europe-west4-b"
  region   = "europe-west4"
}

module "gcp-europe-west4-c" {
  source   = "modules/gcp-deployment"
  gcp_zone = "europe-west4-c"
  region   = "europe-west4"
}

module "gcp-asia-south1-a" {
  source   = "modules/gcp-deployment"
  gcp_zone = "asia-south1-a"
  region   = "asia-south1"
}

module "gcp-asia-south1-b" {
  source   = "modules/gcp-deployment"
  gcp_zone = "asia-south1-b"
  region   = "asia-south1"
}

module "gcp-asia-south1-c" {
  source   = "modules/gcp-deployment"
  gcp_zone = "asia-south1-c"
  region   = "asia-south1"
}

module "asia-southeast1-a" {
  source   = "modules/gcp-deployment"
  gcp_zone = "asia-southeast1-a"
  region   = "asia-southeast1"
}

module "asia-southeast1-b" {
  source   = "modules/gcp-deployment"
  gcp_zone = "asia-southeast1-b"
  region   = "asia-southeast1"
}

module "asia-southeast1-c" {
  source   = "modules/gcp-deployment"
  gcp_zone = "asia-southeast1-c"
  region   = "asia-southeast1"
}

module "gcp-asia-east1-b" {
  source   = "modules/gcp-deployment"
  gcp_zone = "asia-east1-b"
  region   = "asia-east1"
}

module "gcp-asia-east1-c" {
  source   = "modules/gcp-deployment"
  gcp_zone = "asia-east1-c"
  region   = "asia-east1"
}

module "asia-northeast1-a" {
  source   = "modules/gcp-deployment"
  gcp_zone = "asia-northeast1-a"
  region   = "asia-northeast1"
}

module "gcp-asia-northeast1-b" {
  source   = "modules/gcp-deployment"
  gcp_zone = "asia-northeast1-b"
  region   = "asia-northeast1"
}

module "gcp-asia-northeast1-c" {
  source   = "modules/gcp-deployment"
  gcp_zone = "asia-northeast1-c"
  region   = "asia-northeast1"
}

module "gcp-australia-southeast1-a" {
  source   = "modules/gcp-deployment"
  gcp_zone = "australia-southeast1-a"
  region   = "australia-southeast1"
}

module "gcp-australia-southeast1-b" {
  source   = "modules/gcp-deployment"
  gcp_zone = "australia-southeast1-b"
  region   = "australia-southeast1"
}

module "gcp-australia-southeast1-c" {
  source   = "modules/gcp-deployment"
  gcp_zone = "australia-southeast1-c"
  region   = "australia-southeast1"
}
