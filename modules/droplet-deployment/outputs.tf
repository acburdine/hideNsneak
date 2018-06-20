output "do_instance_id" {
  value = "${digitalocean_droplet.default.*.id}"
}

output "do_region" {
  value = "${digitalocean_droplet.default.*.region}"
}

output "do_ipv4_address" {
  value = "${digitalocean_droplet.default.*.ipv4_address}"
}

output "do_status" {
  value = "${digitalocean_droplet.default.*.status}"
}

output "do_tags" {
  value = "${digitalocean_droplet.default.*.tags}"
}
