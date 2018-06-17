provider "aws" {
  access_key = "${var.aws_access_key}"
  secret_key = "${var.aws_secret_key}"
  region     = "${var.aws_region}"
}

data "aws_subnet_ids" "all" {
  vpc_id = "${data.aws_vpc.default.id}"
}

data "aws_vpc" "default" {
  default = true
}

data "aws_ami" "ubuntu" {
  most_recent = true

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-xenial-16.04-amd64-server-*"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  owners = ["099720109477"] # Canonical
}

module "ec2-instance" {
  source  = "terraform-aws-modules/ec2-instance/aws"
  version = "1.9.0"

  # insert the 5 required variables here
  ami                    = "${data.aws_ami.ubuntu.id}"
  instance_type          = "t2.micro"
  name                   = "test"
  instance_count         = 0
  subnet_id              = "${element(data.aws_subnet_ids.all.ids, 0)}"
  vpc_security_group_ids = ["${module.security_group.this_security_group_id}"]
}

module "security_group" {
  source = "terraform-aws-modules/security-group/aws"

  name        = "example"
  description = "Security group for example usage with EC2 instance"
  vpc_id      = "${data.aws_vpc.default.id}"

  ingress_with_cidr_blocks = [
    {
      from_port   = 22
      to_port     = 22
      protocol    = "tcp"
      description = "Default SSH Security Group"
      cidr_blocks = "0.0.0.0/0"
    },
  ]
}
