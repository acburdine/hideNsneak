# output "instance_id" {
#   value = "${digitalocean_droplet.default.*.id}"
# }

output "domainfront_url" {
  value = "${aws_cloudfront_distribution.domain_front.*.domain_name}"
}
