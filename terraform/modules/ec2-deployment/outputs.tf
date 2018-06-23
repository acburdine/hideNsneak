output "instance_id" {
  value = "${aws_instance.hideNsneak.*.id}"
}

output "availability_zone" {
  value = "${aws_instance.hideNsneak.*.availability_zone}"
}

output "key_name" {
  value = "${aws_instance.hideNsneak.*.key_name}"
}

output "public_ip" {
  value = "${aws_instance.hideNsneak.*.public_ip}"
}

output "private_ip" {
  value = "${aws_instance.hideNsneak.*.private_ip}"
}

output "security_group" {
  value = "${aws_instance.hideNsneak.*.security_groups}"
}

output "security_group_id" {
  value = "${aws_security_group.allow_ssh.*.id}"
}
