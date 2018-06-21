########AZURE###########
output "azure_instance_id" {
  value = "${module.azure-example-1.instance_id}"
}

output "azure_ip_address" {
  value = "${module.azure-example-1.ip_address}"
}

output "azure_resource_group_id" {
  value = "${module.azure-example-1.resource_group_id}"
}

########DO###########
output "do_instance_id" {
  value = "${module.do-example-1.instance_id}"
}

output "do_region" {
  value = "${module.do-example-1.region}"
}

output "do_ipv4_address" {
  value = "${module.do-example-1.ipv4_address}"
}

output "do_status" {
  value = "${module.do-example-1.status}"
}

########AWS###########

output "domainfront_url" {
  value = "${module.cloudfront.domainfront_url}"
}

output "ec2_instance_id" {
  value = "${module.aws-us-east-1.instance_id}"
}

output "ec2_availability_zone" {
  value = "${module.aws-us-east-1.availability_zone}"
}

output "ec2_key_name" {
  value = "${module.aws-us-east-1.key_name}"
}

output "ec2_public_ip" {
  value = "${module.aws-us-east-1.public_ip}"
}

output "ec2_private_ip" {
  value = "${module.aws-us-east-1.private_ip}"
}

output "ec2_security_group" {
  value = "${module.aws-us-east-1.security_group}"
}

output "aws_security_group_id" {
  value = "${module.aws-us-east-1.security_group_id}"
}

#########GCP###########
output "gcp_instance_id" {
  value = "${module.gcp-northamerica-northeast1-a.instance_id}"
}

output "gcp_public_ip" {
  value = "${module.gcp-northamerica-northeast1-a.public_ip}"
}

output "gcp_private_ip" {
  value = "${module.gcp-northamerica-northeast1-a.private_ip}"
}

output "gcp_tags_fingerprint" {
  value = "${module.gcp-northamerica-northeast1-a.tags_fingerprint}"
}

output "gcp_metadata" {
  value = "${module.gcp-northamerica-northeast1-a.metadata}"
}
