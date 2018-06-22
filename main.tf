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

# module "{.GOTEMPLATE}" {
#   source         = "modules/ec2-deployment"
#   aws_region     = "us-east-1"
#   aws_access_key = "${var.aws_access_key}"
#   aws_secret_key = "${var.aws_secret_key}"
#   region_count   = 0
# }

module "cloudfront" {
  source            = "modules/cloudfront-deployment"
  cloudfront_origin = "google.com"
  aws_access_key    = "${var.aws_access_key}"
  aws_secret_key    = "${var.aws_secret_key}"
}

module "aws-us-east-1" {
  source          = "modules/ec2-deployment"
  aws_region      = "us-east-1"
  aws_access_key  = "${var.aws_access_key}"
  aws_secret_key  = "${var.aws_secret_key}"
  default_sg_name = "tester-us-east-1"
  region_count    = 0

  #use_custom_ami = false
  #custom_ami = "<custom ami>"
}

# module "aws-us-east-2" {
#   source         = "modules/ec2-deployment"
#   aws_region     = "us-east-2"
#   aws_access_key = "${var.aws_access_key}"
#   aws_secret_key = "${var.aws_secret_key}"
#   region_count   = 0
# }

# module "aws-us-west-1" {
#   source         = "modules/ec2-deployment"
#   aws_region     = "us-west-1"
#   aws_access_key = "${var.aws_access_key}"
#   aws_secret_key = "${var.aws_secret_key}"
#   region_count   = 0
# }

# module "aws-us-west-2" {
#   source         = "modules/ec2-deployment"
#   aws_region     = "us-west-2"
#   aws_access_key = "${var.aws_access_key}"
#   aws_secret_key = "${var.aws_secret_key}"
#   region_count   = 0
# }

# module "aws-ca-central-1" {
#   source         = "modules/ec2-deployment"
#   aws_region     = "ca-central-1"
#   aws_access_key = "${var.aws_access_key}"
#   aws_secret_key = "${var.aws_secret_key}"
#   region_count   = 0
# }

# module "aws-eu-west-1" {
#   source         = "modules/ec2-deployment"
#   aws_region     = "eu-west-1"
#   aws_access_key = "${var.aws_access_key}"
#   aws_secret_key = "${var.aws_secret_key}"
#   region_count   = 0
# }

# module "aws-eu-west-2" {
#   source         = "modules/ec2-deployment"
#   aws_region     = "eu-west-2"
#   aws_access_key = "${var.aws_access_key}"
#   aws_secret_key = "${var.aws_secret_key}"
#   region_count   = 0
# }

# module "aws-eu-central-1" {
#   source         = "modules/ec2-deployment"
#   aws_region     = "eu-central-1"
#   aws_access_key = "${var.aws_access_key}"
#   aws_secret_key = "${var.aws_secret_key}"
#   region_count   = 0
# }

# module "aws-ap-northeast-1" {
#   source         = "modules/ec2-deployment"
#   aws_region     = "ap-northeast-1"
#   aws_access_key = "${var.aws_access_key}"
#   aws_secret_key = "${var.aws_secret_key}"
#   region_count   = 0
# }

# module "aws-ap-northeast-2" {
#   source         = "modules/ec2-deployment"
#   aws_region     = "ap-northeast-2"
#   aws_access_key = "${var.aws_access_key}"
#   aws_secret_key = "${var.aws_secret_key}"
#   region_count   = 0
# }

# module "aws-ap-southeast-1" {
#   source         = "modules/ec2-deployment"
#   aws_region     = "ap-southeast-1"
#   aws_access_key = "${var.aws_access_key}"
#   aws_secret_key = "${var.aws_secret_key}"
#   region_count   = 0
# }

# module "aws-ap-southeast-2" {
#   source         = "modules/ec2-deployment"
#   aws_region     = "ap-southeast-2"
#   aws_access_key = "${var.aws_access_key}"
#   aws_secret_key = "${var.aws_secret_key}"
#   region_count   = 0
# }

# module "aws-ap-south-1" {
#   source         = "modules/ec2-deployment"
#   aws_region     = "ap-south-1"
#   aws_access_key = "${var.aws_access_key}"
#   aws_secret_key = "${var.aws_secret_key}"
#   region_count   = 0
# }

# module "aws-sa-east-1" {
#   source         = "modules/ec2-deployment"
#   aws_region     = "sa-east-1"
#   aws_access_key = "${var.aws_access_key}"
#   aws_secret_key = "${var.aws_secret_key}"
#   region_count   = 0
# }

##########################################
################DO########################
module "do-example-1" {
  source          = "modules/droplet-deployment"
  do_token        = "${var.do_token}"
  do_image        = "ubuntu-14-04-x64"
  pvt_key         = "/Users/mike.hodges/.ssh/do_rsa"
  ssh_fingerprint = "b3:b2:c7:b1:73:9e:28:c6:61:8d:15:e1:0e:61:7e:35"
  do_region       = "NYC1"
  do_size         = "512mb"
  do_count        = 0
}

##########################################
################AZURE#####################

##FROM TERRAFORM####
# We recommend using a Service Principal 
# when running in a shared environment 
# (such as within a CI server/automation) - 
# and authenticating via the Azure CLI when 
# you're running Terraform locally.

provider "azurerm" {
  subscription_id = "${var.azure_subscription_id}"
  client_id       = "${var.azure_client_id}"
  client_secret   = "${var.azure_client_secret}"
  tenant_id       = "${var.azure_tenant_id}"
}

resource "azurerm_resource_group" "test" {
  name     = "acctestrg"
  count    = 1
  location = "West US"
}

module "azure-cdn-example" {
  source                  = "modules/azure-cdn-deployment"
  azure_subscription_id   = "${var.azure_subscription_id}"
  azure_tenant_id         = "${var.azure_tenant_id}"
  azure_client_id         = "${var.azure_client_id}"
  azure_client_secret     = "${var.azure_client_secret}"
  azure_cdn_hostname      = "google.com"
  azure_cdn_profile_name  = "tester123"
  azure_cdn_endpoint_name = "whoknows123"
  azure_location          = "West US"
}

module "azure-example-1" {
  source                = "modules/azure-deployment"
  azure_subscription_id = "${var.azure_subscription_id}"
  azure_tenant_id       = "${var.azure_tenant_id}"
  azure_client_id       = "${var.azure_client_id}"
  azure_client_secret   = "${var.azure_client_secret}"
  azure_location        = "West US"
  azure_instance_count  = 0
}

##########################################
####################GCP###################
##########################################

module "gcp-northamerica-northeast1-a" {
  source               = "modules/gcp-deployment"
  gcp_region           = "northamerica-northeast1"
  gcp_project          = "inboxa90"
  gcp_instance_count   = 0
  gcp_ssh_user         = "mike.hodges"
  gcp_ssh_pub_key_file = "/Users/mike.hodges/.ssh/do_rsa.pub"

  #gcp_machine_type
}

# module "gcp-us-central1-f" {
#   source     = "modules/gcp-deployment"
#   gcp_region = "us-central1"
# }


# module "gcp-us-west1-b" {
#   source     = "modules/gcp-deployment"
#   gcp_region = "us-west1"
# }


# module "gcp-us-east4-c" {
#   source     = "modules/gcp-deployment"
#   gcp_region = "us-east4"
# }


# module "gcp-us-east1-c" {
#   source     = "modules/gcp-deployment"
#   gcp_region = "us-east1"
# }


# module "gcp-southamerica-east1-a" {
#   source     = "modules/gcp-deployment"
#   gcp_region = "southamerica-east1"
# }


# module "gcp-urope-north1-a" {
#   source     = "modules/gcp-deployment"
#   gcp_region = "europe-north1"
# }


# module "gcp-europe-west1-c" {
#   source     = "modules/gcp-deployment"
#   gcp_region = "europe-west1"
# }


# module "gcp-europe-west2-a" {
#   source     = "modules/gcp-deployment"
#   gcp_region = "europe-west2"
# }


# module "gcp-europe-west3-a" {
#   source     = "modules/gcp-deployment"
#   gcp_region = "europe-west3"
# }


# module "gcp-asia-south1-a" {
#   source     = "modules/gcp-deployment"
#   gcp_region = "asia-south1"
# }


# module "asia-southeast1-a" {
#   source     = "modules/gcp-deployment"
#   gcp_region = "asia-southeast1"
# }


# module "gcp-asia-east1-b" {
#   source     = "modules/gcp-deployment"
#   gcp_region = "asia-east1"
# }


# module "asia-northeast1-a" {
#   source     = "modules/gcp-deployment"
#   gcp_region = "asia-northeast1"
# }


# module "gcp-australia-southeast1-a" {
#   source     = "modules/gcp-deployment"
#   gcp_region = "australia-southeast1"
# }

