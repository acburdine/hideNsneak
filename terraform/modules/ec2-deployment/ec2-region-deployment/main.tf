provider "aws" {}

resource "random_string" "sg-extension" {
  length  = 4
  special = false
}

resource "ansible_host" "hideNsneak" {
  count = "${var.instance_count}"

  inventory_hostname = "${aws_instance.hideNsneak.*.public_ip[count.index]}"
  groups             = "${var.ansible_groups}"

  vars {
    ansible_user                 = "${var.ec2_default_user}"
    ansible_connection           = "ssh"
    ansible_ssh_private_key_file = "${var.aws_private_key_file}"
    ansible_ssh_common_args      = "-o StrictHostKeyChecking=no"
  }

  depends_on = ["aws_instance.hideNsneak"]
}

locals {
  keyCount = "${var.instance_count > 0 ? 1 : 0}"
}

data "aws_subnet_ids" "all" {
  count  = "${var.instance_count > 0 ? 1 : 0}"
  vpc_id = "${data.aws_vpc.default.id}"
}

data "aws_vpc" "default" {
  count   = "${var.instance_count > 0 ? 1 : 0}"
  default = true
}

data "aws_ami" "ubuntu" {
  count       = "${var.instance_count > 0 ? 1 : 0}"
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

//Removed due to other key import functionality
# resource "aws_key_pair" "hideNsneak" {
#   key_name   = "${var.aws_keypair_name}"
#   public_key = "${file(var.aws_public_key_file)}"
#   count      = "${var.aws_keypair_name == "hidensneak" ? local.keyCount : 0}"
# }

resource "aws_instance" "hideNsneak" {
  ami = "${data.aws_ami.ubuntu.id}"

  # Old ami identification
  # ami                    = "${var.custom_ami == "" ? data.aws_ami.ubuntu.id : var.custom_ami}"
  instance_type = "${var.aws_instance_type == "" ? "t2.micro" :  var.aws_instance_type}"

  count     = "${var.instance_count}"
  subnet_id = "${element(data.aws_subnet_ids.all.ids, 0)}"

  # vpc_security_group_ids = ["${aws_security_group.allow_ssh.id}"]

  //Rething how we are going to apply firewalls
  # "${var.aws_sg_id == "" ? element(concat(aws_security_group.allow_ssh.*.id, list("")), 0) : var.aws_sg_id }"
  key_name = "${var.aws_keypair_name}"
  tags {
    Name = "hideNsneak"
  }

  # provisioner "local-exec" {
  #   command = "sleep 120; ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook -u ${var.ec2_default_user} --private-key ${var.aws_private_key_file} -i '${self.public_ip},' ../ansible/setup.yml"
  # }

  # depends_on = ["aws_security_group.allow_ssh"]
}

//TODO: Pop security groups out into their own module in order to
//keep configurations upon the creation of new instances
# resource "aws_security_group" "allow_ssh" {
#   name        = "${var.default_sg}${random_string.sg-extension.result}"
#   description = "Allow SSH Traffic"
#   vpc_id      = "${data.aws_vpc.default.id}"
#   count       = "${var.instance_count > 0 ? 1 : 0}"


#   ingress {
#     from_port   = 22
#     to_port     = 22
#     protocol    = "tcp"
#     cidr_blocks = ["0.0.0.0/0"]
#   }


#   egress {
#     from_port   = 0
#     to_port     = 0
#     protocol    = "-1"
#     cidr_blocks = ["0.0.0.0/0"]
#   }
# }

