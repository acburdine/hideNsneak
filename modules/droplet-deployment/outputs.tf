output "instance_id" {
  value = "${digitalocean_droplet.default.*.id}"
}

output "region" {
  value = "${digitalocean_droplet.default.*.region}"
}

output "ipv4_address" {
  value = "${digitalocean_droplet.default.*.ipv4_address}"
}

output "status" {
  value = "${digitalocean_droplet.default.*.status}"
}
