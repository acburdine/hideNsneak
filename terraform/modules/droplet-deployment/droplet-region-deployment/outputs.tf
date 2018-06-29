output "instance_id" {
  value = "${digitalocean_droplet.hideNsneak.*.id}"
}

output "region" {
  value = "${var.do_region}"
}

output "ipv4_address" {
  value = "${digitalocean_droplet.hideNsneak.*.ipv4_address}"
}

output "status" {
  value = "${digitalocean_droplet.hideNsneak.*.status}"
}
