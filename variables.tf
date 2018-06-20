variable "do_token" {}
variable "aws_access_key" {}
variable "aws_secret_key" {}

variable "aws_regions" {
  default = ["us-east-1", "us-east-2", "us-west-1", "us-west-2", "ca-central-1", "eu-west-1", "eu-west-2", "eu-central-1", "ap-northeast-1", "ap-northeast-2", "ap-southeast-1", "ap-southeast-2", "ap-south-1", "sa-east-1"]
}

variable "do_regions" {
  default = ["NYC1"]
}

variable "gcp_regions" {
  default = ["northamerica-northeast1", "us-central1", "us-west1", "us-east4", "us-east1", "southamerica-east1", "europe-north1", "europe-west1", "europe-west2", "europe-west3", "asia-south1", "asia-southeast1", "asia-east1", "asia-northeast1", "australia-southeast1"]
}

variable "provider_map" {
  type = "map"

  default = {
    aws = 0
    do  = 0
    gcp = 0
  }
}

variable "aws_count" {
  default = 1
}

variable "do_count" {
  default = 0
}

variable "gcp_count" {
  default = 1
}

variable "azure_tenant_id" {}

variable "azure_client_id" {}

variable "azure_client_secret" {}

variable "azure_subscription_id" {}
