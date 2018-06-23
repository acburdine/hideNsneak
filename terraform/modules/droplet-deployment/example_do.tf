provider "digitalocean" {
  token = "${var.do_token}"
}

resource "random_string" "droplet_name" {
  length  = 8
  special = false
}

resource "digitalocean_droplet" "default" {
  image  = "${var.do_image}"
  name   = "${var.do_name}${random_string.droplet_name.result}"
  region = "${var.do_region}"
  size   = "${var.do_size}"
  count  = "${var.do_count}"

  ssh_keys = [
    "${var.ssh_fingerprint}",
  ]

  provisioner "local-exec" {
    command = "sleep 120; ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook -u ${var.do_default_user} --private-key ${var.pvt_key} -i '${self.ipv4_address},' master.yml"
  }
}

resource "digitalocean_firewall" "default" {
  name = "${var.do_firewall_name}${random_string.droplet_name.result}"

  droplet_ids = ["${digitalocean_droplet.default.*.id}"]
  count       = "${digitalocean_droplet.default.count > 0 ? 1 : 0}"

  inbound_rule = [
    {
      protocol         = "tcp"
      port_range       = "22"
      source_addresses = ["${var.do_ssh_source_ip}"]
    },
  ]

  outbound_rule = [
    {
      protocol              = "tcp"
      port_range            = "1-65535"
      destination_addresses = ["0.0.0.0/0", "::/0"]
    },
    {
      protocol              = "udp"
      port_range            = "1-65535"
      destination_addresses = ["0.0.0.0/0", "::/0"]
    },
  ]
}
