provider "digitalocean" {
  token = "${var.do_token}"
}

resource "random_string" "droplet_name" {
  length  = 8
  special = false
}

resource "ansible_host" "hideNsneak" {
  count = "${var.do_count}"

  //Element
  inventory_hostname = "${digitalocean_droplet.hideNsneak.*.ipv4_address[count.index]}"
  groups             = "${var.ansible_groups}"

  vars {
    ansible_user                 = "${var.do_default_user}"
    ansible_connection           = "ssh"
    ansible_ssh_private_key_file = "${var.pvt_key}"
    ansible_ssh_common_args      = "-o StrictHostKeyChecking=no"
  }

  depends_on = ["digitalocean_droplet.hideNsneak"]
}

resource "digitalocean_droplet" "hideNsneak" {
  image  = "${var.do_image}"
  name   = "${var.do_name}${random_string.droplet_name.result}"
  region = "${var.do_region}"
  size   = "${var.do_size}"
  count  = "${var.do_count}"

  ssh_keys = [
    "${var.ssh_fingerprint}",
  ]

  provisioner "local-exec" {
    command = "sleep 120; ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook -u ${var.do_default_user} --private-key ${var.pvt_key} -i '${self.ipv4_address},' ../ansible/setup.yml"
  }
}

resource "digitalocean_firewall" "hideNsneak" {
  name = "${var.do_firewall_name}${random_string.droplet_name.result}"

  droplet_ids = ["${digitalocean_droplet.hideNsneak.*.id}"]
  count       = "${digitalocean_droplet.hideNsneak.count > 0 ? 1 : 0}"

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
