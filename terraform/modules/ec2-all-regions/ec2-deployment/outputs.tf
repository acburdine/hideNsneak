output "instance_id" {
  value = "${aws_instance.hideNsneak.*.id}"
}

output "ipInstanceId" {
  value = "${zipmap(aws_instance.hideNsneak.*.public_ip, aws_instance.hideNsneak.*.id)}"
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

output "region-map" {
  value = "${map(
    "region", var.aws_region,
    "region_count", var.region_count,
    "custom_ami", var.custom_ami,
    "public_key_file", var.aws_public_key_file,
    "private_key_file", var.aws_private_key_file,
    "default_sg_name", var.default_sg_name,
    "aws_sg_id", var.aws_sg_id,
    "aws_instance_type", var.aws_instance_type,
    "ec2_default_user", var.ec2_default_user,
    )}"
}
