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

resource "aws_instance" "terraform-play" {
  ami             = "${var.custom_ami == "" ? data.aws_ami.ubuntu.id : var.custom_ami}"
  instance_type   = "${var.aws_instance_type}"
  count           = "${var.region_count}"
  subnet_id       = "${element(data.aws_subnet_ids.all.ids, 0)}"
  security_groups = ["${aws_security_group.allow_ssh.name}"]

  tags {
    Name = "${var.aws_tags}"
  }

  depends_on = ["aws_security_group.allow_ssh"]
}

resource "aws_security_group" "allow_ssh" {
  name        = "${var.default_sg_name}"
  description = "Allow SSH Traffic"
  vpc_id      = "${data.aws_vpc.default.id}"
  count       = "${var.region_count > 0 ? 1 : 0}"

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}
