output "ec2_instance_id" {
  value = "${aws_instance.terraform-play.*.id}"
}

output "ec2_availability_zone" {
  value = "${aws_instance.terraform-play.*.availability_zone}"
}

output "ec2_key_name" {
  value = "${aws_instance.terraform-play.*.key_name}"
}

output "ec2_public_ip" {
  value = "${aws_instance.terraform-play.*.public_ip}"
}

output "ec2_private_ip" {
  value = "${aws_instance.terraform-play.*.private_ip}"
}

output "ec2_security_group" {
  value = "${aws_instance.terraform-play.*.security_groups}"
}

output "aws_security_group_id" {
  value = "${aws_security_group.allow_ssh.*.id}"
}
