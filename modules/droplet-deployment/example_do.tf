provider "digitalocean" {
  token = "${var.do_token}"
}

resource "digitalocean_droplet" "default" {
  image  = "${var.do_image}"
  name   = "example-droplet2"
  region = "${var.do_region}"
  size   = "${var.do_size}"
  count  = 1

  ssh_keys = [
    "${var.ssh_fingerprint}",
  ]

  # connection {
  #   user        = "root"
  #   type        = "ssh"
  #   private_key = "${file(var.pvt_key)}"
  #   timeout     = "2m"
  # }
}

resource "digitalocean_firewall" "default" {
  name = "only-22-2"

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
