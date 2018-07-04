variable "region_count_template" {
  type = "map"

  default = {
    "us-east-1"      = 0
    "us-east-2"      = 0
    "us-west-1"      = 0
    "us-west-2"      = 0
    "ca-central-1"   = 0
    "eu-central-1"   = 0
    "eu-west-1"      = 0
    "eu-west-2"      = 0
    "eu-west-3"      = 0
    "ap-northeast-1" = 0
    "ap-northeast-2" = 0
    "ap-southeast-1" = 0
    "ap-southeast-2" = 0
    "ap-south-1"     = 0
    "sa-east-1"      = 0
  }
}

variable "region_count" {
  type = "map"
}

variable "aws_access_key" {}
variable "aws_secret_key" {}

variable "custom_ami" {
  default = ""
}

variable "aws_public_key_file" {}

variable "aws_keypair_name" {
  default = "hidensneak"
}

# variable "default_sg_name" {
#   default = "ssh_inbound2"
# }

# variable "aws_sg_id" {
#   default = ""
# }

variable "aws_instance_type" {
  default = "t2.micro"
}

variable "ec2_default_user" {
  default = "ubuntu"
}

variable "aws_private_key_file" {}

# variable "aws_tags" {
#   default = ""
# }

