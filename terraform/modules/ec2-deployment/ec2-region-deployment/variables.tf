variable "aws_region" {}

variable "instance_count" {}

variable "aws_public_key_file" {}

variable "aws_keypair_name" {}

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

variable "security_group" {
  default = "hidensneak"
}

variable "ansible_groups" {
  default = []
}
