output "region_info" {
  value = "${map(
    "config", map(
    "region_count", digitalocean_droplet.hideNsneak.count,
    "fingerprint", var.do_ssh_fingerprint,
    "private_key_file", var.do_private_key,
    "size", var.do_size,
    "image", var.do_image,
    "default_user", var.do_default_user,
    "region", var.do_region,
    "self", terraform.
    ),
    "ip_id", zipmap(digitalocean_droplet.hideNsneak.*.ipv4_address, digitalocean_droplet.hideNsneak.*.id),
  )}"
}
