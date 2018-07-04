variable "aws_region" {}

variable "instance_count" {}

# variable "custom_ami" {
#   default = ""
# }

variable "aws_public_key_file" {}

variable "aws_keypair_name" {}

variable "default_sg" {}

variable "aws_sg_id" {}

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
