variable "aws_access_key" {}
variable "aws_secret_key" {}
variable "aws_region" {}

variable "region_count" {
  default = 0
}

variable "custom_ami" {
  default = ""
}

variable "default_sg_name" {}

variable "aws_instance_type" {
  default = "t2.micro"
}

variable "aws_tags" {
  default = ""
}
