output "instance_id" {
  value = "${aws_instance.terraform-play.*.id}"
}

output "availability_zone" {
  value = "${aws_instance.terraform-play.*.availability_zone}"
}

output "key_name" {
  value = "${aws_instance.terraform-play.*.key_name}"
}

output "public_ip" {
  value = "${aws_instance.terraform-play.*.public_ip}"
}

output "private_ip" {
  value = "${aws_instance.terraform-play.*.private_ip}"
}

output "security_group" {
  value = "${aws_instance.terraform-play.*.security_groups}"
}

output "security_group_id" {
  value = "${aws_security_group.allow_ssh.*.id}"
}
