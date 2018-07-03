output "security_group" {
  value = "${aws_instance.hideNsneak.*.security_groups}"
}

output "security_group_id" {
  value = "${aws_security_group.allow_ssh.*.id}"
}

output "region_info" {
  value = "${map(
    "config", map(
    "region_count", aws_instance.hideNsneak.count,
    "custom_ami", var.custom_ami,
    "public_key_file", var.aws_public_key_file,
    "private_key_file", var.aws_private_key_file,
    "default_sg_name", var.default_sg_name,
    "aws_sg_id", var.aws_sg_id,
    "aws_instance_type", var.aws_instance_type,
    "ec2_default_user", var.ec2_default_user,
    "region", var.aws_region
    ),
    "ip_id", map("ip",aws_instance.hideNsneak.*.public_ip, "id",aws_instance.hideNsneak.*.id),
  )}"
}
