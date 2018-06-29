provider "digitalocean" {
  token = "${var.do_token}"
}

locals {
  do_region_count_final = "${merge(var.do_region_count_template, var.do_region_count)}"
}

module "do-nyc1" {
  source             = "droplet-region-deployment"
  do_image           = "${var.do_image}"
  do_size            = "${var.do_size}"
  do_count           = "${local.do_region_count_final["nyc1"]}"
  do_ssh_fingerprint = "${var.do_ssh_fingerprint}"
  do_private_key     = "${var.do_private_key}"
  do_default_user    = "${var.do_default_user}"
  do_region          = "nyc1"

  //Think about seperating firewalls
  do_firewall_name = "${var.do_firewall_name}"
}
