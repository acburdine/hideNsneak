variable "region-count" {
  type = "map"

  default = {
    "us-east-1" = 0
    "us-west-1" = 0
  }
}

variable "aws_region" {}
variable "aws_access_key" {}
variable "aws_secret_key" {}

variable "region_count" {}

variable "custom_ami" {
  default = ""
}

variable "aws_public_key_file" {}

variable "aws_keypair_name" {}

variable "aws_new_keypair" {
  default = true
}

variable "default_sg_name" {}

variable "aws_sg_id" {
  default = ""
}

variable "aws_instance_type" {
  default = "t2.micro"
}

variable "ec2_default_user" {
  default = "ubuntu"
}

variable "aws_private_key_file" {}

variable "aws_tags" {
  default = ""
}

variable "ansible_groups" {
  default = []
}
