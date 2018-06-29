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

module "do-nyc2" {
  source             = "droplet-region-deployment"
  do_image           = "${var.do_image}"
  do_size            = "${var.do_size}"
  do_count           = "${local.do_region_count_final["nyc2"]}"
  do_ssh_fingerprint = "${var.do_ssh_fingerprint}"
  do_private_key     = "${var.do_private_key}"
  do_default_user    = "${var.do_default_user}"
  do_region          = "nyc2"

  //Think about seperating firewalls
  do_firewall_name = "${var.do_firewall_name}"
}

module "do-nyc3" {
  source             = "droplet-region-deployment"
  do_image           = "${var.do_image}"
  do_size            = "${var.do_size}"
  do_count           = "${local.do_region_count_final["nyc3"]}"
  do_ssh_fingerprint = "${var.do_ssh_fingerprint}"
  do_private_key     = "${var.do_private_key}"
  do_default_user    = "${var.do_default_user}"
  do_region          = "nyc3"

  //Think about seperating firewalls
  do_firewall_name = "${var.do_firewall_name}"
}

module "do-ams2" {
  source             = "droplet-region-deployment"
  do_image           = "${var.do_image}"
  do_size            = "${var.do_size}"
  do_count           = "${local.do_region_count_final["ams2"]}"
  do_ssh_fingerprint = "${var.do_ssh_fingerprint}"
  do_private_key     = "${var.do_private_key}"
  do_default_user    = "${var.do_default_user}"
  do_region          = "ams2"

  //Think about seperating firewalls
  do_firewall_name = "${var.do_firewall_name}"
}

module "do-ams3" {
  source             = "droplet-region-deployment"
  do_image           = "${var.do_image}"
  do_size            = "${var.do_size}"
  do_count           = "${local.do_region_count_final["ams3"]}"
  do_ssh_fingerprint = "${var.do_ssh_fingerprint}"
  do_private_key     = "${var.do_private_key}"
  do_default_user    = "${var.do_default_user}"
  do_region          = "ams3"

  //Think about seperating firewalls
  do_firewall_name = "${var.do_firewall_name}"
}

module "do-blr1" {
  source             = "droplet-region-deployment"
  do_image           = "${var.do_image}"
  do_size            = "${var.do_size}"
  do_count           = "${local.do_region_count_final["blr1"]}"
  do_ssh_fingerprint = "${var.do_ssh_fingerprint}"
  do_private_key     = "${var.do_private_key}"
  do_default_user    = "${var.do_default_user}"
  do_region          = "blr1"

  //Think about seperating firewalls
  do_firewall_name = "${var.do_firewall_name}"
}

module "do-fra1" {
  source             = "droplet-region-deployment"
  do_image           = "${var.do_image}"
  do_size            = "${var.do_size}"
  do_count           = "${local.do_region_count_final["fra1"]}"
  do_ssh_fingerprint = "${var.do_ssh_fingerprint}"
  do_private_key     = "${var.do_private_key}"
  do_default_user    = "${var.do_default_user}"
  do_region          = "fra1"

  //Think about seperating firewalls
  do_firewall_name = "${var.do_firewall_name}"
}

module "do-lon1" {
  source             = "droplet-region-deployment"
  do_image           = "${var.do_image}"
  do_size            = "${var.do_size}"
  do_count           = "${local.do_region_count_final["lon1"]}"
  do_ssh_fingerprint = "${var.do_ssh_fingerprint}"
  do_private_key     = "${var.do_private_key}"
  do_default_user    = "${var.do_default_user}"
  do_region          = "lon1"

  //Think about seperating firewalls
  do_firewall_name = "${var.do_firewall_name}"
}

module "do-sfo1" {
  source             = "droplet-region-deployment"
  do_image           = "${var.do_image}"
  do_size            = "${var.do_size}"
  do_count           = "${local.do_region_count_final["sfo1"]}"
  do_ssh_fingerprint = "${var.do_ssh_fingerprint}"
  do_private_key     = "${var.do_private_key}"
  do_default_user    = "${var.do_default_user}"
  do_region          = "sfo1"

  //Think about seperating firewalls
  do_firewall_name = "${var.do_firewall_name}"
}

module "do-sfo2" {
  source             = "droplet-region-deployment"
  do_image           = "${var.do_image}"
  do_size            = "${var.do_size}"
  do_count           = "${local.do_region_count_final["sfo2"]}"
  do_ssh_fingerprint = "${var.do_ssh_fingerprint}"
  do_private_key     = "${var.do_private_key}"
  do_default_user    = "${var.do_default_user}"
  do_region          = "sfo2"

  //Think about seperating firewalls
  do_firewall_name = "${var.do_firewall_name}"
}

module "do-tor1" {
  source             = "droplet-region-deployment"
  do_image           = "${var.do_image}"
  do_size            = "${var.do_size}"
  do_count           = "${local.do_region_count_final["tor1"]}"
  do_ssh_fingerprint = "${var.do_ssh_fingerprint}"
  do_private_key     = "${var.do_private_key}"
  do_default_user    = "${var.do_default_user}"
  do_region          = "tor1"

  //Think about seperating firewalls
  do_firewall_name = "${var.do_firewall_name}"
}

module "do-sgp1" {
  source             = "droplet-region-deployment"
  do_image           = "${var.do_image}"
  do_size            = "${var.do_size}"
  do_count           = "${local.do_region_count_final["sgp1"]}"
  do_ssh_fingerprint = "${var.do_ssh_fingerprint}"
  do_private_key     = "${var.do_private_key}"
  do_default_user    = "${var.do_default_user}"
  do_region          = "sgp1"

  //Think about seperating firewalls
  do_firewall_name = "${var.do_firewall_name}"
}
