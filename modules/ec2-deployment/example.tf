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
  ami           = "${data.aws_ami.ubuntu.id}"
  instance_type = "t2.micro"
  count         = "${var.region_count}"
  subnet_id     = "${element(data.aws_subnet_ids.all.ids, 0)}"

  tags {
    Name = "test"
  }

  depends_on = ["aws_security_group.allow_ssh"]
}

# module "ec2-instance" {
#   source  = "terraform-aws-modules/ec2-instance/aws"
#   version = "1.9.0"

#   # insert the 5 required variables here
#   ami                    = "${data.aws_ami.ubuntu.id}"
#   instance_type          = "t2.micro"
#   name                   = "test"
#   instance_count         = "${var.region_count}"
#   subnet_id              = "${element(data.aws_subnet_ids.all.ids, 0)}"
#   vpc_security_group_ids = ["${aws_security_group.allow_ssh.security_group_id}"]

#   depends_on = ["aws_security_group.allow_ssh"]
# }

# module "security_group" {
#   source      = "terraform-aws-modules/security-group/aws"
#   name        = "example"
#   description = "Security group for example usage with EC2 instance"
#   vpc_id      = "${data.aws_vpc.default.id}"

#   ingress_with_cidr_blocks = [
#     {
#       from_port   = 22
#       to_port     = 22
#       protocol    = "tcp"
#       description = "Default SSH Security Group"
#       cidr_blocks = "0.0.0.0/0"
#     },
#   ]
# }
resource "aws_security_group" "allow_ssh" {
  name        = "SSH_Inbound"
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

output "security_group_id" {
  value = "${aws_security_group.allow_ssh.*.id}"
}
