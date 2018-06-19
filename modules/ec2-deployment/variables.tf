variable "aws_access_key" {}
variable "aws_secret_key" {}
variable "aws_region" {}

variable "provider_list" {
  default = ["aws.us-east-1", "aws.us-east-2", "aws.us-west-1", "aws.us-west-2", "aws.ca-central-1", "aws.eu-west-1", "aws.eu-west-2", "aws.eu-central-1", "aws.ap-northeast-1", "aws.ap-northeast-2", "aws.ap-southeast-1", "aws.ap-southeast-2", "aws.ap-south-1", "aws.sa-east-1"]
}

variable "region_count" {
  default = 0
}
